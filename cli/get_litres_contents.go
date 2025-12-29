package cli

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/redux"
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

	reduxDir := data.Pwd.AbsRelDirPath(data.Redux, data.Metadata)

	rdx, err := redux.NewReader(reduxDir, data.ContentsUrlProperty)
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

	if err = getSetContents(hc, force, rdx, ids...); err != nil {
		return err
	}

	return nil

}

func getSetContents(hc *http.Client, force bool, rdx redux.Readable, ids ...string) error {

	gsc := nod.NewProgress(" contents...")
	defer gsc.Done()

	if err := rdx.MustHave(data.ContentsUrlProperty); err != nil {
		return err
	}

	absContentsDir := data.Pwd.AbsRelDirPath(data.Contents, data.Metadata)

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

	gsc.TotalInt(len(newIds))

	errs := make(map[string]error)
	for _, id := range newIds {
		if path, ok := rdx.GetLastVal(data.ContentsUrlProperty, id); ok && path != "" {
			if err = getSetData(id, litres_integration.ContentsUrl(path), hc, kv); err != nil {
				errs[id] = err
			}
		}
		gsc.Increment()
	}

	if len(errs) > 0 {
		errStrs := make([]string, 0, len(errs))
		for id, err := range errs {
			errStrs = append(errStrs, fmt.Sprintf("%s: %s", id, err.Error()))
		}
		gsc.EndWithResult("errors: " + strings.Join(errStrs, ","))
	}

	return nil
}
