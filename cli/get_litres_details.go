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
	"github.com/boggydigital/pathology"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func GetLitResDetailsHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	newOnly := u.Query().Has("new-only")
	noThrottle := u.Query().Has("no-throttle")

	return GetLitResDetails(ids, newOnly, noThrottle)
}

func GetLitResDetails(ids []string, newOnly, noThrottle bool) error {

	gmbda := nod.NewProgress("getting LitRes my books details...")
	defer gmbda.End()

	absReduxDir, err := pathology.GetAbsRelDir(data.Redux)
	if err != nil {
		return gmbda.EndWithError(err)
	}

	rdx, err := kvas.NewReduxReader(absReduxDir,
		data.MyBooksIdsProperty,
		data.HrefProperty,
		data.ImportedProperty)
	if err != nil {
		return gmbda.EndWithError(err)
	}

	absCookiesFilename, err := data.AbsCookiesFilename()
	if err != nil {
		return gmbda.EndWithError(err)
	}

	absLitResMyBooksDetailsDir, err := data.AbsDataTypeDir(litres_integration.LitResMyBooksDetails)
	if err != nil {
		return gmbda.EndWithError(err)
	}

	kv, err := kvas.ConnectLocal(absLitResMyBooksDetailsDir, kvas.HtmlExt)
	if err != nil {
		return gmbda.EndWithError(err)
	}

	cj, err := coost.NewJar(absCookiesFilename)
	if err != nil {
		return gmbda.EndWithError(err)
	}

	hc := cj.NewHttpClient()

	if len(ids) == 0 {
		var ok bool
		ids, ok = rdx.GetAllValues(data.MyBooksIdsProperty, data.MyBooksIdsProperty)
		if !ok {
			err = errors.New("no my books found")
			return gmbda.EndWithError(err)
		}
	}

	gmbda.TotalInt(len(ids))

	for _, id := range ids {

		if err := getMyBookById(id, hc, rdx, kv, newOnly); err != nil {
			return err
		}

		if err := cj.Store(absCookiesFilename); err != nil {
			return err
		}

		gmbda.Increment()
	}

	gmbda.EndWithResult("done")

	return nil
}

func getMyBookById(id string, hc *http.Client, rdx kvas.ReadableRedux, kv kvas.KeyValues, newOnly bool) error {

	gia := nod.Begin(" getting %s...", id)
	defer gia.End()

	// don't attempt downloading details for imported books
	if IsImported(id, rdx) {
		gia.EndWithResult("imported")
		return nil
	}

	if newOnly && kv.Has(id) {
		gia.EndWithResult("already exists")
		return nil
	}

	href, ok := rdx.GetFirstVal(data.HrefProperty, id)
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

	if err := kv.Set(id, strings.NewReader(str)); err != nil {
		return gia.EndWithError(err)
	}

	return nil
}
