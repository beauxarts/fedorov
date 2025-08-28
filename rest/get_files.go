package rest

import (
	"net/http"

	"github.com/beauxarts/fedorov/rest/compton_pages"
	"github.com/boggydigital/nod"
)

func GetFiles(w http.ResponseWriter, r *http.Request) {

	// GET /files?id

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, nod.ErrorStr("missing book id"), http.StatusInternalServerError)
		return
	}

	var err error
	if rdx, err = rdx.RefreshReader(); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	if p := compton_pages.Files(id, rdx); p != nil {
		if err = p.WriteResponse(w); err != nil {
			http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		}
	}
}
