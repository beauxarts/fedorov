package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/beauxarts/fedorov/rest/compton_pages"
	"github.com/boggydigital/nod"
	"net/http"
)

func GetSimilar(w http.ResponseWriter, r *http.Request) {

	// GET /similar?id

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

	saReader, err := data.NewArtsReader(litres_integration.ArtsTypeSimilar)
	if err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	var similarArts *litres_integration.ArtsSimilar
	if sa, err := saReader.ArtsSimilar(id); err == nil {
		similarArts = sa
	} else {
		_ = nod.Error(err)
	}

	if p := compton_pages.Similar(id, similarArts, rdx); p != nil {
		if err = p.WriteResponse(w); err != nil {
			http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		}
	}
}
