package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/view_models"
	"github.com/boggydigital/nod"
	"net/http"
	"os"
)

func GetDownloads(w http.ResponseWriter, r *http.Request) {

	// GET /downloads?id

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, nod.ErrorStr("missing required book id"), http.StatusInternalServerError)
		return
	}

	links, ok := rxa.GetAllUnchangedValues(data.DownloadLinksProperty, id)

	if !ok {
		http.Error(w, nod.ErrorStr("book has no downloads"), http.StatusInternalServerError)
		return
	}

	availability := make(map[string]bool)
	for _, link := range links {
		path := data.AbsDownloadPath(id, link)
		_, err := os.Stat(path)
		availability[link] = err == nil
	}

	dvm := view_models.NewDownloads(links, availability)

	if err := tmpl.ExecuteTemplate(w, "downloads-page", dvm); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

}
