package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/stencil_app"
	"github.com/boggydigital/nod"
	"golang.org/x/exp/maps"
	"net/http"
)

func GetLocalTagsEdit(w http.ResponseWriter, r *http.Request) {

	// GET /local-tags/edit?id

	id := r.URL.Query().Get("id")

	allValues := make(map[string]bool)
	for _, id := range rxa.Keys(data.LocalTagsProperty) {
		if values, ok := rxa.GetAllUnchangedValues(data.LocalTagsProperty, id); ok {
			for _, v := range values {
				allValues[v] = true
			}
		}
	}

	selectedValues := make(map[string]bool)
	if values, ok := rxa.GetAllUnchangedValues(data.LocalTagsProperty, id); ok {
		for _, v := range values {
			selectedValues[v] = true
		}
	}

	title, _ := rxa.GetFirstVal(data.TitleProperty, id)

	if err := app.RenderPropertyEditor(
		id,
		title,
		stencil_app.PropertyTitles[data.LocalTagsProperty],
		true,
		"",
		selectedValues,
		maps.Keys(allValues),
		true,
		"/local-tags/apply",
		w); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}
