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
	"golang.org/x/exp/maps"
	"net/url"
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

	return GetLitResSeries(seriesTypes, force, ids...)
}

func GetLitResSeries(seriesTypes []litres_integration.SeriesType, force bool, ids ...string) error {
	glsa := nod.NewProgress("getting litres series...")
	defer glsa.End()

	absReduxDir, err := pasu.GetAbsRelDir(data.Redux)
	if err != nil {
		return glsa.EndWithError(err)
	}

	if len(ids) == 0 {

		series := make(map[string]interface{})

		rdx, err := kvas.NewReduxReader(absReduxDir, data.ArtsHistoryOrderProperty, data.SeriesIdProperty)
		if err != nil {
			return glsa.EndWithError(err)
		}

		if artsIds, ok := rdx.GetAllValues(data.ArtsHistoryOrderProperty, data.ArtsHistoryOrderProperty); ok {
			for _, id := range artsIds {
				if seriesIds, sure := rdx.GetAllValues(data.SeriesIdProperty, id); sure {
					for _, sid := range seriesIds {
						series[sid] = nil
					}
				}
			}
		}

		ids = maps.Keys(series)
	}

	glsa.TotalInt(len(ids))

	absCookiesFilename, err := data.AbsCookiesFilename()
	if err != nil {
		return glsa.EndWithError(err)
	}

	cj, err := coost.NewJar(absCookiesFilename)
	if err != nil {
		return glsa.EndWithError(err)
	}

	hc := cj.NewHttpClient()

	dc := dolo.NewClient(hc, dolo.Defaults())

	for _, st := range seriesTypes {
		if err := getSetSeriesType(dc, st, force, ids...); err != nil {
			return glsa.EndWithError(err)
		}
	}

	glsa.EndWithResult("done")

	return nil
}

func getSetSeriesType(dc *dolo.Client, st litres_integration.SeriesType, force bool, ids ...string) error {
	gsst := nod.NewProgress(" %s...", st)
	defer gsst.End()

	absSeriesTypeDir, err := data.AbsSeriesTypeDir(st)
	if err != nil {
		return gsst.EndWithError(err)
	}

	kv, err := kvas.ConnectLocal(absSeriesTypeDir, kvas.JsonExt)
	if err != nil {
		return gsst.EndWithError(err)
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
