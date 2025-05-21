package cli

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/kevlar_dolo"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathways"
	"github.com/boggydigital/redux"
	"net/http"
	"net/url"
	"strings"
)

func GetLitResContentsHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	force := u.Query().Has("force")

	return GetLitresContents(nil, force, ids...)
}

func GetLitresContents(hc *http.Client, force bool, ids ...string) error {

	dlca := nod.NewProgress("downloading litres contents...")
	defer dlca.Done()

	rdx, err := data.NewReduxReader(data.ContentsUrlProperty)
	if err != nil {
		return err
	}

	if len(ids) == 0 {
		ids, err = GetRecentArts(force) // = rdx.Keys(data.ContentsUrlProperty)
		if err != nil {
			return err
		}
	}

	dlca.TotalInt(len(ids))

	if hc == nil {
		hc, err = getHttpClient()
		if err != nil {
			return err
		}
	}

	dc := dolo.NewClient(hc, dolo.Defaults())

	if err := getSetContents(dc, force, rdx, ids...); err != nil {
		return err
	}

	return nil

}

func getSetContents(dc *dolo.Client, force bool, rdx redux.Readable, ids ...string) error {

	gsc := nod.NewProgress(" contents...")
	defer gsc.Done()

	if err := rdx.MustHave(data.ContentsUrlProperty); err != nil {
		return err
	}

	absContentsDir, err := pathways.GetAbsRelDir(data.Contents)
	if err != nil {
		return err
	}

	kv, err := kevlar.New(absContentsDir, kevlar.XmlExt)
	if err != nil {
		return err
	}

	newIds := make([]string, 0, len(ids))
	for _, id := range ids {
		if kv.Has(id) && !force {
			continue
		}
		newIds = append(newIds, id)
	}

	// filtering ids to only those that actually have contents-url
	filteredIds := make([]string, 0)
	urls := make([]*url.URL, 0, len(newIds))
	for _, id := range newIds {
		if path, ok := rdx.GetLastVal(data.ContentsUrlProperty, id); ok && path != "" {
			urls = append(urls, litres_integration.ContentsUrl(path))
			filteredIds = append(filteredIds, id)
		}
	}

	indexSetter := kevlar_dolo.NewIndexSetter(kv, filteredIds...)
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
