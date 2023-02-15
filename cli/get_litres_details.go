package cli

import (
	"errors"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/litres_integration"
	"github.com/boggydigital/coost"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"net/http"
	"net/url"
	"strings"
)

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

	return GetLitResDetails(ids, hc, newOnly)
}

func GetLitResDetails(ids []string, hc *http.Client, newOnly bool) error {

	gmbda := nod.NewProgress("getting LitRes my books details...")
	defer gmbda.End()

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), nil,
		data.MyBooksIdsProperty,
		data.HrefProperty,
		data.ImportedProperty)
	if err != nil {
		return gmbda.EndWithError(err)
	}

	kv, err := kvas.ConnectLocal(data.AbsMyBooksDetailsDir(), kvas.HtmlExt)
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

	for _, id := range ids {

		// don't attempt downloading details for imported books
		if IsImported(id, rxa) {
			continue
		}

		if newOnly && kv.Has(id) {
			continue
		}

		href, ok := rxa.GetFirstVal(data.HrefProperty, id)
		if !ok {
			err = errors.New("no href for book " + id)
			return gmbda.EndWithError(err)
		}

		resp, err := hc.Get(litres_integration.HrefUrl(href).String())
		if err != nil {
			nod.Log(err.Error())
			continue
		}

		if err := kv.Set(id, resp.Body); err != nil {
			resp.Body.Close()
			return gmbda.EndWithError(err)
		}

		resp.Body.Close()

		// sleep for 5 seconds to throttle server requests
		//time.Sleep(time.Second * 5)

		gmbda.Increment()
	}

	gmbda.EndWithResult("done")

	return nil
}
