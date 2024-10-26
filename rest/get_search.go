package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/stencil_app"
	"github.com/boggydigital/nod"
	"golang.org/x/exp/maps"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func GetSearch(w http.ResponseWriter, r *http.Request) {

	// GET /search?(search_params)&from

	q := r.URL.Query()

	from, to := 0, 0
	if q.Has("from") {
		from64, err := strconv.ParseInt(q.Get("from"), 10, 32)
		if err != nil {
			http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
			return
		}
		from = int(from64)
	}

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
	if rdx, err = rdx.RefreshReader(); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	if len(query) > 0 {

		ids = rdx.Match(query)

		if sort := r.URL.Query().Get(data.SortProperty); sort != "" {
			desc := r.URL.Query().Get(data.DescendingProperty) == "true"
			ids, err = rdx.Sort(ids, desc, sort, data.TitleProperty)
			if err != nil {
				http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
				return
			}
		}

		if from > len(ids)-1 {
			from = 0
		}

		to = from + SearchResultsLimit
		if to > len(ids) {
			to = len(ids)
		} else if to+SearchResultsLimit > len(ids) {
			to = len(ids)
		}
	}

	if err := app.RenderSearch("Поиск", query, ids[from:to], from, to, len(ids), r.URL, rdx, w); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}

func getDigests(properties ...string) map[string][]string {
	digests := make(map[string][]string)
	for _, p := range properties {
		dv := make(map[string]interface{})
		for _, id := range rdx.Keys(p) {
			values, ok := rdx.GetAllValues(p, id)
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
