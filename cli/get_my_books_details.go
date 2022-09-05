package cli

import (
	"errors"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/litres_integration"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"net/http"
)

func GetMyBooksDetails(hc *http.Client) error {

	gmbda := nod.NewProgress("getting my books details...")
	defer gmbda.End()

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), nil,
		data.MyBooksIdsProperty,
		data.HrefProperty)
	if err != nil {
		return gmbda.EndWithError(err)
	}

	kv, err := kvas.ConnectLocal(data.AbsMyBooksDetailsDir(), kvas.HtmlExt)
	if err != nil {
		return gmbda.EndWithError(err)
	}

	ids, ok := rxa.GetAllValues(data.MyBooksIdsProperty, data.MyBooksIdsProperty)
	if !ok {
		err = errors.New("no my books found")
		return gmbda.EndWithError(err)
	}

	gmbda.TotalInt(len(ids))

	for _, id := range ids {
		href, ok := rxa.GetFirstVal(data.HrefProperty, id)
		if !ok {
			err = errors.New("no href for book " + id)
			return gmbda.EndWithError(err)
		}

		resp, err := hc.Get(litres_integration.HrefUrl(href).String())
		if err != nil {
			resp.Body.Close()
			return gmbda.EndWithError(err)
		}

		if err := kv.Set(id, resp.Body); err != nil {
			resp.Body.Close()
			return gmbda.EndWithError(err)
		}

		resp.Body.Close()
		gmbda.Increment()
	}

	gmbda.EndWithResult("done")

	return nil
}
