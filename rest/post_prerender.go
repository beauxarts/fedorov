package rest

import (
	"github.com/beauxarts/fedorov/stencil_app"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/stencil/stencil_rest"
	"net/http"
)

func PostPrerender(w http.ResponseWriter, _ *http.Request) {

	// POST /prerender

	// the following pages will be pre-rendered:
	// - default path (/updates)
	// - every top-level search route

	if err := updatePrerender(); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}

func updatePrerender() error {
	paths := []string{
		"/books",
	}

	for _, p := range stencil_app.SearchScopeQueries() {
		sp := stencil_app.SearchPath
		if p != "" {
			sp = sp + "?" + p
		}
		paths = append(paths, sp)
	}

	return stencil_rest.Prerender(paths, port)
}
