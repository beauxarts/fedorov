package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/view_models"
	"github.com/boggydigital/nod"
	"net/http"
	"os"
	"path/filepath"
)

func GetDownloads(w http.ResponseWriter, r *http.Request) {

	// GET /downloads?id

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, nod.ErrorStr("missing required book id"), http.StatusInternalServerError)
		return
	}

	var err error
	if rxa, err = rxa.RefreshReduxAssets(); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	links, ok := rxa.GetAllUnchangedValues(data.DownloadLinksProperty, id)

	if !ok {
		http.Error(w, nod.ErrorStr("book has no downloads"), http.StatusInternalServerError)
		return
	}

	files := make([]string, 0, len(links))
	for _, link := range links {
		_, filename := filepath.Split(link)
		if _, err := os.Stat(data.AbsDownloadPath(id, filename)); err == nil {
			files = append(files, filename)
		}
	}

	dvm := view_models.NewDownloads(id, files)

	DefaultHeaders(w)

	if err := tmpl.ExecuteTemplate(w, "downloads-page", dvm); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

}
