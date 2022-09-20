package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/stencil_app"
	"github.com/boggydigital/nod"
	"golang.org/x/exp/maps"
	"net/http"
	"sort"
	"strings"
)

func GetSearch(w http.ResponseWriter, r *http.Request) {

	// GET /search?(search_params)

	q := r.URL.Query()

	query := make(map[string][]string)

	shortQuery := false
	queryProperties := stencil_app.SearchProperties
	for _, p := range queryProperties {
		if v := q.Get(p); v != "" {
			query[p] = strings.Split(v, ",")
		} else {
			if q.Has(p) {
				q.Del(p)
				shortQuery = true
			}
		}
	}

	//if we removed some properties with no values - redirect to the shortest URL
	if shortQuery {
		r.URL.RawQuery = q.Encode()
		http.Redirect(w, r, r.URL.String(), http.StatusPermanentRedirect)
		return
	}

	var ids []string

	var err error
	if rxa, err = rxa.RefreshReduxAssets(); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	if len(query) > 0 {

		ids = maps.Keys(rxa.Match(query, true))

		if sort := r.URL.Query().Get(data.SortProperty); sort != "" {
			desc := r.URL.Query().Get(data.DescendingProperty) == "true"
			ids, err = rxa.Sort(ids, desc, sort, data.TitleProperty)
			if err != nil {
				http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	digests := getDigests(stencil_app.DigestProperties...)

	digests[data.SortProperty] = []string{
		data.TitleProperty,
		data.DateCreatedProperty,
		data.DateTranslatedProperty,
		data.DateReleasedProperty}

	digests[data.DescendingProperty] = []string{
		"true",
		"false"}

	DefaultHeaders(w)

	if err := rapp.RenderSearch("Поиск", query, ids, digests, w); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}

func getDigests(properties ...string) map[string][]string {
	digests := make(map[string][]string)
	for _, p := range properties {
		dv := make(map[string]interface{})
		for _, id := range rxa.Keys(p) {
			values, ok := rxa.GetAllValues(p, id)
			if !ok {
				continue
			}
			for _, v := range values {
				dv[v] = nil
			}
		}
		sortedDv := maps.Keys(dv)
		sort.Strings(sortedDv)
		digests[p] = sortedDv
	}
	return digests
}
