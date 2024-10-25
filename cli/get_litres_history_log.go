package cli

import (
	"bufio"
	"encoding/json"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/match_node"
	"github.com/boggydigital/nod"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const maxSupportedPages = 1000

func GetLitResHistoryLogHandler(u *url.URL) error {
	sessionId := u.Query().Get("session-id")

	return GetLitResHistoryLog(sessionId, nil)
}

func GetLitResHistoryLog(sessionId string, hc *http.Client) error {
	ghla := nod.NewProgress("fetching LitRes history log...")
	defer ghla.End()

	absLitResHistoryLogDir, err := data.AbsDataTypeDir(litres_integration.LitResHistoryLog)
	if err != nil {
		return ghla.EndWithError(err)
	}

	kv, err := kevlar.NewKeyValues(absLitResHistoryLogDir, kevlar.HtmlExt)
	if err != nil {
		return ghla.EndWithError(err)
	}

	if hc == nil {
		hc, err = getHttpClient()
		if err != nil {
			return ghla.EndWithError(err)
		}
	}

	// get the first page and extract total pages

	page := 1

	if err := getHistoryLogPage(page, hc, kv, ghla); err != nil {
		return ghla.EndWithError(err)
	}

	totalPages, err := getTotalHistoryLogPages(kv)
	if err != nil {
		return ghla.EndWithError(err)
	}

	ghla.TotalInt(totalPages)

	for page = 2; page <= totalPages; page++ {
		if err := getHistoryLogPage(page, hc, kv, ghla); err != nil {
			return ghla.EndWithError(err)
		}
	}

	return nil
}

func getHistoryLogPage(page int, hc *http.Client, kv kevlar.KeyValues, tpw nod.TotalProgressWriter) error {
	resp, err := hc.Get(litres_integration.HistoryLogPageUrl(page).String())
	if err != nil {
		return tpw.EndWithError(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		tpw.Increment()
		return nil
	}

	if err = kv.Set(strconv.Itoa(page), resp.Body); err != nil {
		return err
	}

	tpw.Increment()
	return nil
}

const (
	pchlProperty = "personal_cabinet_history_log:"
)

type PersonalCabinetHistoryLog struct {
	Pages int `json:"pages"`
}

func getTotalHistoryLogPages(kv kevlar.KeyValues) (int, error) {

	if fp, err := kv.Get("1"); err != nil {
		return 1, err
	} else {
		body, err := html.Parse(fp)
		if err != nil {
			return 1, err
		}
		etcScript := match_node.NewEtc(atom.Script, "", false)
		for _, script := range match_node.Matches(body, etcScript, -1) {
			if script.FirstChild != nil && strings.Contains(script.FirstChild.Data, pchlProperty) {
				scanner := bufio.NewScanner(strings.NewReader(script.FirstChild.Data))
				found := false
				sb := &strings.Builder{}
				for scanner.Scan() {
					t := scanner.Text()
					if found {
						if strings.Contains(t, "}") {
							sb.WriteString("}")
							break
						} else {
							sb.WriteString(strings.TrimSpace(t))
						}
					}

					if strings.Contains(t, pchlProperty) {
						found = true
						sb.WriteString("{")
					}
				}

				pchlStr := strings.Replace(sb.String(), "pages", "\"pages\"", 1)

				var pchl PersonalCabinetHistoryLog

				decoder := json.NewDecoder(strings.NewReader(pchlStr))
				if err := decoder.Decode(&pchl); err != nil {
					return 1, err
				}

				return pchl.Pages, nil
			}
		}
	}
	return 1, nil
}
