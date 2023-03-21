package rest

import (
	"github.com/beauxarts/fedorov/stencil_app"
	"github.com/boggydigital/stencil/stencil_rest"
	"net/http"
)

func PostPrerender(w http.ResponseWriter, r *http.Request) {

	// POST /prerender

	// the following pages will be pre-rendered:
	// - default path (/updates)
	// - every top-level search route

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

	stencil_rest.Prerender(paths, port, w)

}
