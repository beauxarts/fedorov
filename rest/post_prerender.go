package rest

import (
	"bytes"
	"fmt"
	"github.com/beauxarts/fedorov/stencil_app"
	"github.com/boggydigital/middleware"
	"github.com/boggydigital/nod"
	"io"
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

	// we don't want to accumulate existing static content over the lifetime of the app
	middleware.ClearStaticContent()

	host := fmt.Sprintf("http://localhost:%d", port)

	for _, p := range paths {
		if err := setStaticContent(host, p); err != nil {
			http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
			return
		}
	}

	if _, err := io.WriteString(w, "ok"); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

}

func setStaticContent(host, p string) error {
	resp, err := http.DefaultClient.Get(host + p)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bs := make([]byte, 0, 1024*1024)
	bb := bytes.NewBuffer(bs)

	if _, err := io.Copy(bb, resp.Body); err != nil {
		return err
	}

	middleware.SetStaticContent(p, bb.Bytes())

	return nil
}
