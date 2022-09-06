package rest

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/nod"
	"net/http"
	"os"
	"path/filepath"
)

func GetFile(w http.ResponseWriter, r *http.Request) {

	// GET /file?id&href

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, nod.ErrorStr("missing required book id"), http.StatusInternalServerError)
		return
	}

	href := r.URL.Query().Get("href")

	if href == "" {
		http.Error(w, nod.ErrorStr("missing required file href"), http.StatusInternalServerError)
		return
	}

	filePath := data.AbsDownloadPath(id, href)
	_, filename := filepath.Split(filePath)
	if _, err := os.Stat(filePath); err == nil {
		w.Header().Set("Cache-Control", "max-age=31536000")
		w.Header().Set("Content-Disposition", "attachment; filename=\""+filename+"\"")
		http.ServeFile(w, r, filePath)
	} else {
		_ = nod.Error(fmt.Errorf("no file for id %s, href %s", id, href))
		http.NotFound(w, r)
	}

}
