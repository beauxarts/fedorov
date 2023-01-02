package rest

import (
	"encoding/json"
	"github.com/beauxarts/fedorov/data"
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

	var digests map[string][]string

	switch property {
	case data.SortProperty:
		digests = map[string][]string{data.SortProperty: {
			data.TitleProperty,
			data.DateCreatedProperty,
			data.DateTranslatedProperty,
			data.DateReleasedProperty}}
	case data.DescendingProperty:
		digests = map[string][]string{data.DescendingProperty: {
			"true",
			"false"}}
	default:
		digests = getDigests(property)
	}

	DefaultHeaders(w)

	if err := json.NewEncoder(w).Encode(digests[property]); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}
