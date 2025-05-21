package cli

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathways"
	"github.com/boggydigital/redux"
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
	defer glsa.Done()

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

	for _, st := range seriesTypes {
		if err := getSetSeriesType(hc, st, force, seriesIds...); err != nil {
			return err
		}
	}

	return nil
}

func getSeriesIds(force bool, artsIds ...string) ([]string, error) {
	series := make(map[string]interface{})

	reduxDir, err := pathways.GetAbsRelDir(data.Redux)
	if err != nil {
		return nil, err
	}

	rdx, err := redux.NewReader(reduxDir, data.ArtsOperationsOrderProperty, data.SeriesIdProperty)
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

func getSetSeriesType(hc *http.Client, st litres_integration.SeriesType, force bool, ids ...string) error {
	gsst := nod.NewProgress(" %s...", st)
	defer gsst.Done()

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

	gsst.TotalInt(len(newIds))

	errs := make(map[string]error)
	for _, id := range newIds {
		if err = getSetData(id, litres_integration.SeriesUrl(st, id), hc, kv); err != nil {
			errs[id] = err
		}
		gsst.Increment()
	}

	if len(errs) > 0 {
		errStrs := make([]string, 0, len(errs))
		for id, err := range errs {
			errStrs = append(errStrs, fmt.Sprintf("%s: %s", id, err.Error()))
		}
		gsst.EndWithResult("errors: " + strings.Join(errStrs, ","))
	}

	return nil
}
