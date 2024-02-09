package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pasu"
	"net/http"
)

func GetCompletedSet(w http.ResponseWriter, r *http.Request) {

	// GET /completed/set?id

	id := r.URL.Query().Get(data.IdProperty)

	absReduxDir, err := pasu.GetAbsRelDir(data.Redux)
	if err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	bcRdx, err := kvas.NewReduxWriter(absReduxDir, data.BookCompletedProperty)
	if err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	if err := bcRdx.ReplaceValues(data.BookCompletedProperty, id, "true"); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/book?id="+id, http.StatusTemporaryRedirect)
}
