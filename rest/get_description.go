package rest

import (
	"github.com/beauxarts/fedorov/stencil_app"
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
	if rdx, err = rdx.RefreshReader(); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	desc, ok := rdx.GetLastVal("" /*data.DescriptionProperty*/, id)

	if !ok {
		http.Error(w, nod.ErrorStr("book has no downloads"), http.StatusInternalServerError)
		return
	}

	DefaultHeaders(w)

	if err := app.RenderSection(id, stencil_app.DescriptionSection, desc, w); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}
