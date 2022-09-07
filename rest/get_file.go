package rest

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/nod"
	"net/http"
	"os"
)

func GetFile(w http.ResponseWriter, r *http.Request) {

	// GET /file?id&href

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, nod.ErrorStr("missing required book id"), http.StatusInternalServerError)
		return
	}

	file := r.URL.Query().Get("file")

	if file == "" {
		http.Error(w, nod.ErrorStr("missing file"), http.StatusInternalServerError)
		return
	}

	localFilepath := data.AbsDownloadPath(id, file)
	if _, err := os.Stat(localFilepath); err == nil {
		w.Header().Set("Cache-Control", "max-age=31536000")
		w.Header().Set("Content-Disposition", "attachment; filename=\""+file+"\"")
		http.ServeFile(w, r, localFilepath)
	} else {
		_ = nod.Error(fmt.Errorf("no file for id %s, file %s", id, file))
		http.NotFound(w, r)
	}

}
