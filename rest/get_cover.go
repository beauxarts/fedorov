package rest

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/nod"
	"net/http"
	"os"
	"path/filepath"
)

func GetCover(w http.ResponseWriter, r *http.Request) {

	// GET /cover?id

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, nod.ErrorStr("missing required book id"), http.StatusInternalServerError)
		return
	}

	coverPath := filepath.Join(data.AbsCoverDir(), id+data.CoverExt)
	if _, err := os.Stat(coverPath); err == nil {
		w.Header().Set("Cache-Control", "max-age=31536000")
		http.ServeFile(w, r, coverPath)
	} else {
		_ = nod.Error(fmt.Errorf("no cover for id %s", id))
		http.NotFound(w, r)
	}

}
