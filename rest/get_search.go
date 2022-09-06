package rest

import (
	"github.com/beauxarts/fedorov/view_models"
	"github.com/boggydigital/nod"
	"golang.org/x/exp/maps"
	"net/http"
	"strings"
)

func GetSearch(w http.ResponseWriter, r *http.Request) {

	// GET /search?(search_params)

	q := r.URL.Query()

	//svm := view_models.NewSearchProducts([]string{}, rxa)
	query := make(map[string][]string)

	shortQuery := false
	queryProperties := view_models.SearchProperties
	for _, p := range queryProperties {
		if v := q.Get(p); v != "" {
			query[p] = []string{v}
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

	if len(query) > 0 {
		ids = maps.Keys(rxa.Match(query, true))
	}

	svm := view_models.NewSearchProducts(ids, rxa)

	for k, v := range query {
		svm.Query[k] = strings.Join(v, " ")
	}

	//digests, cached, err := getDigests(dc, view_models.DigestProperties...)

	//digests[vangogh_local_data.SortProperty] = []string{
	//	vangogh_local_data.GlobalReleaseDateProperty,
	//	vangogh_local_data.GOGReleaseDateProperty,
	//	vangogh_local_data.GOGOrderDateProperty,
	//	vangogh_local_data.TitleProperty,
	//	vangogh_local_data.RatingProperty,
	//	vangogh_local_data.DiscountPercentageProperty}
	//
	//digests[vangogh_local_data.DescendingProperty] = []string{
	//	vangogh_local_data.TrueValue,
	//	vangogh_local_data.FalseValue}

	//spvm.Digests = digests

	//gaugin_middleware.DefaultHeaders(st, w)

	if err := tmpl.ExecuteTemplate(w, "search-page", svm); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}
