package rest

import (
	"encoding/json"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/stencil_app"
	"github.com/boggydigital/nod"
	"net/http"
)

func GetDigest(w http.ResponseWriter, r *http.Request) {
	// GET /digest?property

	property := r.URL.Query().Get("property")
	if property == "" {
		http.Error(w, nod.ErrorStr("missing digest property"), http.StatusBadRequest)
		return
	}

	var values []string
	valueTitles := make(map[string]string)

	switch property {
	case data.SortProperty:
		values = []string{
			data.TitleProperty,
			data.DateCreatedProperty,
			data.DateTranslatedProperty,
			data.DateReleasedProperty,
			data.MyBooksOrderProperty}
	case data.DescendingProperty:
		values = []string{
			"true",
			"false"}
	default:
		values = getDigests(property)[property]
	}

	for _, v := range values {
		if title, ok := stencil_app.PropertyTitles[v]; ok {
			valueTitles[v] = title
		} else {
			valueTitles[v] = v
		}
	}

	DefaultHeaders(w)

	if err := json.NewEncoder(w).Encode(valueTitles); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}
