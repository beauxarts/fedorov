package cli

import (
	"errors"
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/coost"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultThrottleSeconds = 20
	ddosThrottleSeconds    = 200
	ddosGuardStr           = "ddos-guard"
)

var ddosGuardError = errors.New("DDoS Guard detected")

func GetLitResDetailsHandler(u *url.URL) error {
	hc, err := coost.NewHttpClientFromFile(data.AbsCookiesFilename(), litres_integration.LitResHost)
	if err != nil {
		return err
	}

	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	newOnly := u.Query().Has("new-only")
	noThrottle := u.Query().Has("no-throttle")

	return GetLitResDetails(ids, hc, newOnly, noThrottle)
}

func GetLitResDetails(ids []string, hc *http.Client, newOnly, noThrottle bool) error {

	gmbda := nod.NewProgress("getting LitRes my books details...")
	defer gmbda.End()

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(),
		data.MyBooksIdsProperty,
		data.HrefProperty,
		data.ImportedProperty)
	if err != nil {
		return gmbda.EndWithError(err)
	}

	kv, err := kvas.ConnectLocal(data.AbsLitResMyBooksDetailsDir(), kvas.HtmlExt)
	if err != nil {
		return gmbda.EndWithError(err)
	}

	if len(ids) == 0 {
		var ok bool
		ids, ok = rxa.GetAllValues(data.MyBooksIdsProperty, data.MyBooksIdsProperty)
		if !ok {
			err = errors.New("no my books found")
			return gmbda.EndWithError(err)
		}
	}

	gmbda.TotalInt(len(ids))

	ddosIds := make([]string, 0)

	for _, id := range ids {

		if err := getMyBookById(id, hc, defaultThrottleSeconds, rxa, kv, newOnly, noThrottle); errors.Is(err, ddosGuardError) {
			ddosIds = append(ddosIds, id)
		} else if err != nil {
			nod.Log(err.Error())
		}

		gmbda.Increment()
	}

	gmbda.EndWithResult("done")

	if len(ddosIds) > 0 {
		gdia := nod.NewProgress("attempting to get LitRes my books details for DDoS guarded ids")
		defer gdia.End()

		gdia.TotalInt(len(ddosIds))

		for _, id := range ddosIds {

			if err := getMyBookById(id, hc, ddosThrottleSeconds, rxa, kv, newOnly, noThrottle); err != nil {
				nod.Log(err.Error())
			}

			gdia.Increment()
		}

		gdia.EndWithResult("done")

	}

	return nil
}

func getMyBookById(id string, hc *http.Client, ts int, rxa kvas.ReduxAssets, kv kvas.KeyValues, newOnly, noThrottle bool) error {

	// sleep to throttle server requests
	if !noThrottle {
		ta := nod.Begin(" throttling server requests by %ds...", ts)
		time.Sleep(time.Second * time.Duration(ts))
		ta.End()
	}

	gia := nod.Begin(" getting %s...", id)
	defer gia.End()

	// don't attempt downloading details for imported books
	if IsImported(id, rxa) {
		gia.EndWithResult("imported")
		return nil
	}

	if newOnly && kv.Has(id) {
		gia.EndWithResult("already exists")
		return nil
	}

	href, ok := rxa.GetFirstVal(data.HrefProperty, id)
	if !ok {
		err := fmt.Errorf("no href for book " + id)
		return gia.EndWithError(err)
	}

	req, err := http.NewRequest(http.MethodGet, litres_integration.HrefUrl(href).String(), nil)
	if err != nil {
		return gia.EndWithError(err)
	}

	req.Header.Add("User-Agent", dolo.DefaultUserAgent)
	cmt, err := kv.CurrentModTime(id)
	req.Header.Add("If-Modified-Since", time.Unix(cmt, 0).UTC().Format(http.TimeFormat))

	resp, err := hc.Do(req)
	if err != nil {
		return gia.EndWithError(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == 304 {
		gia.EndWithResult("not modified")
		return nil
	}

	if resp.StatusCode < 100 || resp.StatusCode > 299 {
		return gia.EndWithError(err)
	}

	sb := &strings.Builder{}
	if _, err := io.Copy(sb, resp.Body); err != nil {
		return gia.EndWithError(err)
	}

	str := sb.String()

	//do not write out the content if it's a DDoS Guard page
	if strings.Contains(str, ddosGuardStr) {
		return gia.EndWithError(ddosGuardError)
	}

	if err := kv.Set(id, strings.NewReader(str)); err != nil {
		return gia.EndWithError(err)
	}

	return nil
}
