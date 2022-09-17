package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/view_models"
	"github.com/boggydigital/nod"
	"golang.org/x/exp/maps"
	"net/http"
	"sort"
)

func GetSearch(w http.ResponseWriter, r *http.Request) {

	// GET /search?(search_params)

	q := r.URL.Query()

	query := make(map[string]string)

	shortQuery := false
	queryProperties := view_models.SearchProperties
	for _, p := range queryProperties {
		if v := q.Get(p); v != "" {
			query[p] = v
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

		mq := make(map[string][]string)
		for k, v := range query {
			if k == data.SortProperty ||
				k == data.DescendingProperty {
				continue
			}
			mq[k] = []string{v}
		}

		ids = maps.Keys(rxa.Match(mq, true))

		if sort := r.URL.Query().Get(data.SortProperty); sort != "" {
			desc := r.URL.Query().Get(data.DescendingProperty) == "true"
			ids, err = rxa.Sort(ids, sort, desc)
			if err != nil {
				http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	svm := view_models.NewSearchProducts(ids, rxa)
	svm.Query = query

	svm.Digests = getDigests(view_models.DigestProperties...)

	svm.Digests[data.SortProperty] = []string{
		data.TitleProperty,
		data.DateCreatedProperty,
		data.DateTranslatedProperty,
		data.DateReleasedProperty}

	svm.Digests[data.DescendingProperty] = []string{
		"true",
		"false"}

	svm.DigestsTitles = view_models.DigestTitles

	DefaultHeaders(w)

	if err := app.RenderSearch("Поиск", ids, w); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	//if err := tmpl.ExecuteTemplate(w, "search-page", svm); err != nil {
	//	http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
	//	return
	//}
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
