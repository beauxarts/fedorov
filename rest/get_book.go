package rest

import (
	"github.com/beauxarts/fedorov/rest/compton_pages"
	"github.com/boggydigital/nod"
	"net/http"
)

func GetBook(w http.ResponseWriter, r *http.Request) {

	// GET /book?id

	id := r.URL.Query().Get("id")

	var err error
	if rdx, err = rdx.RefreshReader(); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	if p := compton_pages.Book(id, rdx); p != nil {
		if err := p.WriteResponse(w); err != nil {
			http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
			return
		}
	}
}
