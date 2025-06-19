package litres_integration

import (
	"encoding/json"
	"errors"
	"github.com/boggydigital/match_node"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/http"
	"net/url"
	"strings"
)

type nextDataBuildId struct {
	BuildId string `json:"buildId"`
}

func GetSessionId(httpClient *http.Client) (string, error) {
	rootUrl := &url.URL{
		Scheme: httpsScheme,
		Host:   LitResHost,
	}

	resp, err := httpClient.Get(rootUrl.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	if buildId, err := matchBuildId(doc); err == nil && buildId != "" {
		return buildId, nil
	} else if err != nil {
		return "", err
	} else {
		return "", errors.New("buildId is empty")
	}
}

func matchBuildId(doc *html.Node) (string, error) {

	if ndsm := match_node.Match(doc, &nextDataScriptMatcher{}); ndsm != nil && ndsm.FirstChild != nil {
		nextData := ndsm.FirstChild.Data

		var ndBuildId nextDataBuildId
		if err := json.NewDecoder(strings.NewReader(nextData)).Decode(&ndBuildId); err != nil {
			return "", err
		}

		return ndBuildId.BuildId, nil
	}

	return "", errors.New("next data buildId not found")
}

type nextDataScriptMatcher struct {
}

func (ndsm *nextDataScriptMatcher) Match(node *html.Node) bool {
	if node.DataAtom == atom.Script &&
		len(node.Attr) > 0 {

		for _, attr := range node.Attr {
			if attr.Key == "id" &&
				attr.Val == "__NEXT_DATA__" {
				return true
			}
		}

	}
	return false
}
