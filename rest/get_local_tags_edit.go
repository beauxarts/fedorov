package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/stencil_app"
	"github.com/boggydigital/nod"
	"net/http"
)

func GetLocalTagsEdit(w http.ResponseWriter, r *http.Request) {

	// GET /local-tags/edit?id

	id := r.URL.Query().Get("id")

	allValues := make(map[string]string)
	for _, id := range rdx.Keys(data.LocalTagsProperty) {
		if values, ok := rdx.GetAllValues(data.LocalTagsProperty, id); ok {
			for _, v := range values {
				allValues[v] = v
			}
		}
	}

	selectedValues := make(map[string]bool)
	if values, ok := rdx.GetAllValues(data.LocalTagsProperty, id); ok {
		for _, v := range values {
			selectedValues[v] = true
		}
	}

	title, _ := rdx.GetLastVal(data.TitleProperty, id)

	if err := app.RenderPropertyEditor(
		id,
		title,
		stencil_app.PropertyTitles[data.LocalTagsProperty],
		true,
		"",
		selectedValues,
		allValues,
		true,
		"/local-tags/apply",
		w); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}
