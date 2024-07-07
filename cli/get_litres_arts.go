package cli

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/coost"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/kevlar_dolo"
	"github.com/boggydigital/nod"
	"net/url"
	"strings"
)

func GetLitResArtsHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	var artsTypesStr []string
	if ats := u.Query().Get("arts-type"); ats != "" {
		artsTypesStr = strings.Split(ats, ",")
	}

	artsTypes := make([]litres_integration.ArtsType, 0, len(artsTypesStr))

	if allArtsTypes := u.Query().Has("all-arts-types"); allArtsTypes {
		artsTypes = litres_integration.AllArtsTypes()
	} else {
		for _, ats := range artsTypesStr {
			artsTypes = append(artsTypes, litres_integration.ParseArtsType(ats))
		}
	}

	force := u.Query().Has("force")

	return GetLitResArts(artsTypes, force, ids...)
}

func GetLitResArts(artsTypes []litres_integration.ArtsType, force bool, ids ...string) error {

	glaa := nod.NewProgress("getting litres arts...")
	defer glaa.End()

	if len(ids) == 0 {
		rdx, err := data.NewReduxReader(data.ArtsHistoryOrderProperty)
		if err != nil {
			return glaa.EndWithError(err)
		}

		if artsIds, ok := rdx.GetAllValues(data.ArtsHistoryOrderProperty, data.ArtsHistoryOrderProperty); ok {
			for _, mbid := range artsIds {
				ids = append(ids, mbid)
			}
		}
	}

	glaa.TotalInt(len(ids))

	absCookiesFilename, err := data.AbsCookiesFilename()
	if err != nil {
		return glaa.EndWithError(err)
	}

	hc, err := coost.NewHttpClientFromFile(absCookiesFilename)
	if err != nil {
		return glaa.EndWithError(err)
	}

	dc := dolo.NewClient(hc, dolo.Defaults())

	for _, at := range artsTypes {
		if err := getSetArtsType(dc, at, force, ids...); err != nil {
			return glaa.EndWithError(err)
		}
	}

	glaa.EndWithResult("done")

	return nil
}

func getSetArtsType(dc *dolo.Client, at litres_integration.ArtsType, force bool, ids ...string) error {
	gsat := nod.NewProgress(" %s...", at)
	defer gsat.End()

	absArtsTypeDir, err := data.AbsArtsTypeDir(at)
	if err != nil {
		return gsat.EndWithError(err)
	}

	kv, err := kevlar.NewKeyValues(absArtsTypeDir, kevlar.JsonExt)
	if err != nil {
		return gsat.EndWithError(err)
	}

	newIds := make([]string, 0, len(ids))
	for _, id := range ids {
		ok, err := kv.Has(id)
		if err != nil {
			return gsat.EndWithError(err)
		}
		if ok && !force {
			continue
		}
		newIds = append(newIds, id)
	}

	indexSetter := kevlar_dolo.NewIndexSetter(kv, newIds...)
	urls := make([]*url.URL, 0, len(newIds))
	for _, id := range newIds {
		urls = append(urls, litres_integration.ArtsTypeUrl(at, id))
	}

	result := "done"

	if errs := dc.GetSet(urls, indexSetter, gsat, force); len(errs) > 0 {
		errIds := make([]string, 0, len(errs))
		for ii := range errs {
			errIds = append(errIds, newIds[ii])
		}
		result = fmt.Sprintf("GetSet error ids: %s", strings.Join(errIds, ","))
	}

	gsat.EndWithResult(result)

	return nil
}
