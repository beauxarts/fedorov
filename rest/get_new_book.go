package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/nod"
	"net/http"
	"strings"
)

type NewBookViewModel struct {
	Id        string
	Title     string
	Authors   []string
	Downloads []*DownloadViewModel
}

type DownloadViewModel struct {
	Id          string
	Filename    string
	Description string
}

func GetNewBook(w http.ResponseWriter, r *http.Request) {

	// GET /new_book?id

	id := r.URL.Query().Get("id")

	var err error
	if rdx, err = rdx.RefreshReader(); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	title := ""
	if t, ok := rdx.GetLastVal(data.TitleProperty, id); ok {
		title = t
	}

	var authors []string
	if aus, err := authorsFullNames(id, rdx); err == nil {
		authors = aus
	}

	kv, err := data.NewArtsReader(litres_integration.ArtsTypeFiles)
	if err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	artFiles, err := kv.ArtsFiles(id)
	if err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	nbvm := &NewBookViewModel{
		Id:      id,
		Title:   title,
		Authors: authors,
	}

	for _, dt := range artFiles.DownloadsTypes() {

		fn := dt.Filename
		if ext := dt.Extension; ext != nil {
			fn = strings.Replace(fn, "zip", *ext, 1)
		}

		dvm := &DownloadViewModel{
			Id:          id,
			Filename:    fn,
			Description: dt.TypeDescription(),
		}

		nbvm.Downloads = append(nbvm.Downloads, dvm)
	}

	DefaultHeaders(w)

	if err := tmpl.ExecuteTemplate(w, "new_book", nbvm); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

}
