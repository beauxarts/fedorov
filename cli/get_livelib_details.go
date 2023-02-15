package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/livelib_integration"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"net/http"
	"net/url"
	"strings"
)

func GetLiveLibDetailsHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	newOnly := u.Query().Has("new-only")

	return GetLiveLibDetails(ids, http.DefaultClient, newOnly)
}

func GetLiveLibDetails(ids []string, hc *http.Client, newOnly bool) error {

	glbda := nod.NewProgress("getting LiveLib books details...")
	defer glbda.End()

	kv, err := kvas.ConnectLocal(data.AbsLiveLibDetailsDir(), kvas.HtmlExt)
	if err != nil {
		return glbda.EndWithError(err)
	}

	if len(ids) == 0 {
		glbda.EndWithResult("no id specified")
		return nil
	}

	glbda.TotalInt(len(ids))

	for _, id := range ids {

		// don't attempt downloading details for imported books
		if newOnly && kv.Has(id) {
			continue
		}

		resp, err := hc.Get(livelib_integration.BookUrl(id).String())
		if err != nil {
			nod.Log(err.Error())
			continue
		}

		if err := kv.Set(id, resp.Body); err != nil {
			resp.Body.Close()
			return glbda.EndWithError(err)
		}

		resp.Body.Close()

		glbda.Increment()
	}

	glbda.EndWithResult("done")

	return nil
}
