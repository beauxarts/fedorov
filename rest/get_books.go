package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/view_models"
	"github.com/boggydigital/nod"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {

	// GET /books

	myBooks, ok := rxa.GetAllValues(data.MyBooksIdsProperty, data.MyBooksIdsProperty)
	if !ok {
		http.Error(w, nod.ErrorStr("no my books found"), http.StatusInternalServerError)
		return
	}

	shelf := view_models.NewShelf(myBooks, rxa)

	if err := tmpl.ExecuteTemplate(w, "books-page", shelf); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}
