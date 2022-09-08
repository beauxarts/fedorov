package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/view_models"
	"github.com/boggydigital/nod"
	"net/http"
)

func GetDescription(w http.ResponseWriter, r *http.Request) {

	// GET /description?id

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, nod.ErrorStr("missing required book id"), http.StatusInternalServerError)
		return
	}

	var err error
	if rxa, err = rxa.RefreshReduxAssets(); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	desc, ok := rxa.GetFirstVal(data.DescriptionProperty, id)

	if !ok {
		http.Error(w, nod.ErrorStr("book has no downloads"), http.StatusInternalServerError)
		return
	}

	dvm := view_models.NewDescription(id, desc)

	if err := tmpl.ExecuteTemplate(w, "description-page", dvm); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}
