package rest

import (
	"github.com/beauxarts/fedorov/view_models"
	"github.com/boggydigital/nod"
	"net/http"
)

func GetBook(w http.ResponseWriter, r *http.Request) {

	// GET /book?id

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

	bvm := view_models.NewBook(id, rxa)

	if err := tmpl.ExecuteTemplate(w, "book-page", bvm); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}
