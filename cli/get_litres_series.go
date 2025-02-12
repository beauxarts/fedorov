package cli

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/kevlar_dolo"
	"github.com/boggydigital/nod"
	"maps"
	"net/http"
	"net/url"
	"slices"
	"strings"
)

func GetLitResSeriesHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	var seriesTypesStr []string
	if ats := u.Query().Get("series-type"); ats != "" {
		seriesTypesStr = strings.Split(ats, ",")
	}

	seriesTypes := make([]litres_integration.SeriesType, 0, len(seriesTypesStr))

	if allSeriesTypes := u.Query().Has("all-series-types"); allSeriesTypes {
		seriesTypes = litres_integration.AllSeriesTypes()
	} else {
		for _, sts := range seriesTypesStr {
			seriesTypes = append(seriesTypes, litres_integration.ParseSeriesType(sts))
		}
	}

	force := u.Query().Has("force")

	return GetLitResSeries(seriesTypes, nil, force, ids...)
}

func GetLitResSeries(seriesTypes []litres_integration.SeriesType, hc *http.Client, force bool, seriesIds ...string) error {
	glsa := nod.NewProgress("getting litres series...")
	defer glsa.End()

	if len(seriesIds) == 0 {
		var err error
		seriesIds, err = getSeriesIds(force)
		if err != nil {
			return err
		}
	}

	glsa.TotalInt(len(seriesIds))

	if hc == nil {
		var err error
		hc, err = getHttpClient()
		if err != nil {
			return err
		}
	}

	dc := dolo.NewClient(hc, dolo.Defaults())

	for _, st := range seriesTypes {
		if err := getSetSeriesType(dc, st, force, seriesIds...); err != nil {
			return err
		}
	}

	glsa.EndWithResult("done")

	return nil
}

func getSeriesIds(force bool, artsIds ...string) ([]string, error) {
	series := make(map[string]interface{})

	rdx, err := data.NewReduxReader(data.ArtsOperationsOrderProperty, data.SeriesIdProperty)
	if err != nil {
		return nil, err
	}

	if len(artsIds) == 0 && force {
		if allArtsIds, ok := rdx.GetAllValues(data.ArtsOperationsOrderProperty, data.ArtsOperationsOrderProperty); ok {
			artsIds = allArtsIds
		}
	}

	for _, id := range artsIds {
		if seriesIds, sure := rdx.GetAllValues(data.SeriesIdProperty, id); sure {
			for _, sid := range seriesIds {
				series[sid] = nil
			}
		}
	}

	return slices.Collect(maps.Keys(series)), nil
}

func getSetSeriesType(dc *dolo.Client, st litres_integration.SeriesType, force bool, ids ...string) error {
	gsst := nod.NewProgress(" %s...", st)
	defer gsst.End()

	absSeriesTypeDir, err := data.AbsSeriesTypeDir(st)
	if err != nil {
		return err
	}

	kv, err := kevlar.New(absSeriesTypeDir, kevlar.JsonExt)
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

	indexSetter := kevlar_dolo.NewIndexSetter(kv, newIds...)
	urls := make([]*url.URL, 0, len(newIds))
	for _, id := range newIds {
		urls = append(urls, litres_integration.SeriesUrl(st, id))
	}

	result := "done"

	if errs := dc.GetSet(urls, indexSetter, gsst, force); len(errs) > 0 {
		errIds := make([]string, 0, len(errs))
		for ii := range errs {
			errIds = append(errIds, newIds[ii])
		}
		result = fmt.Sprintf("GetSet error ids: %s", strings.Join(errIds, ","))
	}

	gsst.EndWithResult(result)

	return nil
}
