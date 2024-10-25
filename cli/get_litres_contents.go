package cli

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/kevlar_dolo"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathways"
	"net/http"
	"net/url"
	"strings"
)

func GetLitResContentsHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	sessionId := u.Query().Get("session-id")

	force := u.Query().Has("force")

	return GetLitresContents(sessionId, nil, force, ids...)
}

func GetLitresContents(sessionId string, hc *http.Client, force bool, ids ...string) error {

	dlca := nod.NewProgress("downloading litres contents...")
	defer dlca.End()

	rdx, err := data.NewReduxReader(data.ContentsUrlProperty)
	if err != nil {
		return dlca.EndWithError(err)
	}

	if len(ids) == 0 {
		ids = rdx.Keys(data.ContentsUrlProperty)
	}

	dlca.TotalInt(len(ids))

	if hc == nil {
		hc, err = getHttpClient()
		if err != nil {
			return dlca.EndWithError(err)
		}
	}

	dc := dolo.NewClient(hc, dolo.Defaults())

	if err := getSetContents(dc, force, rdx, ids...); err != nil {
		return dlca.EndWithError(err)
	}

	dlca.EndWithResult("done")

	return nil

}

func getSetContents(dc *dolo.Client, force bool, rdx kevlar.ReadableRedux, ids ...string) error {

	gsc := nod.NewProgress(" contents...")
	defer gsc.End()

	if err := rdx.MustHave(data.ContentsUrlProperty); err != nil {
		return gsc.EndWithError(err)
	}

	absContentsDir, err := pathways.GetAbsRelDir(data.Contents)
	if err != nil {
		return gsc.EndWithError(err)
	}

	kv, err := kevlar.NewKeyValues(absContentsDir, kevlar.XmlExt)
	if err != nil {
		return gsc.EndWithError(err)
	}

	newIds := make([]string, 0, len(ids))
	for _, id := range ids {
		ok, err := kv.Has(id)
		if err != nil {
			return gsc.EndWithError(err)
		}
		if ok && !force {
			continue
		}
		newIds = append(newIds, id)
	}

	indexSetter := kevlar_dolo.NewIndexSetter(kv, newIds...)
	urls := make([]*url.URL, 0, len(newIds))
	for _, id := range newIds {
		if path, ok := rdx.GetLastVal(data.ContentsUrlProperty, id); ok && path != "" {
			urls = append(urls, litres_integration.ContentsUrl(path))
		}
	}

	result := "done"

	if errs := dc.GetSet(urls, indexSetter, gsc, force); len(errs) > 0 {
		errIds := make([]string, 0, len(errs))
		for ii := range errs {
			errIds = append(errIds, newIds[ii])
		}
		result = fmt.Sprintf("GetSet error ids: %s", strings.Join(errIds, ","))
	}

	gsc.EndWithResult(result)

	return nil
}
