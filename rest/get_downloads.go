package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/stencil_app"
	"github.com/beauxarts/fedorov/view_models"
	"github.com/boggydigital/nod"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func GetDownloads(w http.ResponseWriter, r *http.Request) {

	// GET /downloads?id

	idstr := r.URL.Query().Get("id")

	if idstr == "" {
		http.Error(w, nod.ErrorStr("missing required book id"), http.StatusInternalServerError)
		return
	}

	var err error
	if rxa, err = rxa.RefreshReduxAssets(); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	links, ok := rxa.GetAllValues(data.DownloadLinksProperty, idstr)
	titles, _ := rxa.GetAllValues(data.DownloadTitlesProperty, idstr)

	if !ok {
		http.Error(w, nod.ErrorStr("book has no downloads"), http.StatusInternalServerError)
		return
	}

	files := make([]string, 0, len(links))

	if id, err := strconv.ParseInt(idstr, 10, 64); err == nil {
		for _, link := range links {
			_, filename := filepath.Split(link)
			absDownloadFilename := data.AbsDownloadPath(id, filename)
			if _, err := os.Stat(absDownloadFilename); err == nil {
				files = append(files, filename)
			} else {
				nod.Log(err.Error())
			}
		}
	}

	sb := &strings.Builder{}
	dvm := view_models.NewDownloads(idstr, files, titles)

	if err := tmpl.ExecuteTemplate(sb, "downloads", dvm); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	DefaultHeaders(w)

	if err := app.RenderSection(idstr, stencil_app.DownloadsSection, sb.String(), w); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

}
