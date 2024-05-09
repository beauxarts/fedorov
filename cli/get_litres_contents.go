package cli

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/coost"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/kvas_dolo"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pasu"
	"net/url"
	"strings"
)

func GetLitResContentsHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	force := u.Query().Has("force")

	return GetLitresContents(force, ids...)
}

func GetLitresContents(force bool, ids ...string) error {

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

	absCookiesFilename, err := data.AbsCookiesFilename()
	if err != nil {
		return dlca.EndWithError(err)
	}

	hc, err := coost.NewHttpClientFromFile(absCookiesFilename)
	if err != nil {
		return dlca.EndWithError(err)
	}

	dc := dolo.NewClient(hc, dolo.Defaults())

	if err := getSetContents(dc, force, rdx, ids...); err != nil {
		return dlca.EndWithError(err)
	}

	dlca.EndWithResult("done")

	return nil

}

func getSetContents(dc *dolo.Client, force bool, rdx kvas.ReadableRedux, ids ...string) error {

	gsc := nod.NewProgress(" contents...")
	defer gsc.End()

	if err := rdx.MustHave(data.ContentsUrlProperty); err != nil {
		return gsc.EndWithError(err)
	}

	absContentsDir, err := pasu.GetAbsRelDir(data.Contents)
	if err != nil {
		return gsc.EndWithError(err)
	}

	kv, err := kvas.ConnectLocal(absContentsDir, kvas.XmlExt)
	if err != nil {
		return gsc.EndWithError(err)
	}

	newIds := make([]string, 0, len(ids))
	for _, id := range ids {
		if !force && kv.Has(id) {
			continue
		}
		newIds = append(newIds, id)
	}

	indexSetter := kvas_dolo.NewIndexSetter(kv, newIds...)
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
