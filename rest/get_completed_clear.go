package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/nod"
	"net/http"
)

func GetCompletedClear(w http.ResponseWriter, r *http.Request) {

	// GET /completed/clear?id

	id := r.URL.Query().Get(data.IdProperty)

	if err := rxa.CutVal(data.BookCompletedProperty, id, "true"); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/book?id="+id, http.StatusTemporaryRedirect)
}
