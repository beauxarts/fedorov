package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/nod"
	"net/http"
)

type Book struct {
	Id                string
	Title             string
	Authors           []string
	AdditionalDetails map[string][]string
	DownloadLinks     []string
}

func NewBook(id string) *Book {

	bvm := &Book{
		Id:                id,
		AdditionalDetails: make(map[string][]string),
	}

	for _, p := range data.ReduxProperties() {

		switch p {
		case data.TitleProperty:
			bvm.Title, _ = rxa.GetFirstVal(p, id)
		case data.AuthorsProperty:
			bvm.Authors, _ = rxa.GetAllUnchangedValues(p, id)
		case data.DownloadLinksProperty:
			bvm.DownloadLinks, _ = rxa.GetAllUnchangedValues(p, id)
		default:
			values, _ := rxa.GetAllUnchangedValues(p, id)
			if len(values) > 0 {
				bvm.AdditionalDetails[p] = values
			}
		}
	}

	return bvm
}

func GetBook(w http.ResponseWriter, r *http.Request) {

	// GET /book?id

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, nod.ErrorStr("missing required book id"), http.StatusInternalServerError)
		return
	}

	bvm := NewBook(id)

	if err := tmpl.ExecuteTemplate(w, "book-page", bvm); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}
