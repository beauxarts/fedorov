package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathology"
	"net/http"
)

func GetCompletedClear(w http.ResponseWriter, r *http.Request) {

	// GET /completed/clear?id

	id := r.URL.Query().Get(data.IdProperty)

	absReduxDir, err := pathology.GetAbsRelDir(data.Redux)
	if err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	bcRdx, err := kvas.NewReduxWriter(absReduxDir, data.BookCompletedProperty)
	if err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	if err := bcRdx.CutValues(data.BookCompletedProperty, id, "true"); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	if err := updatePrerender(); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/book?id="+id, http.StatusTemporaryRedirect)
}
