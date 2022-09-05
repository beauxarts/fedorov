package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/nod"
	"net/http"
	"strings"
)

var booksListProperties = []string{
	data.TitleProperty,
	data.AuthorsProperty,
	data.DateCreatedProperty,
}

type ListBookViewModel struct {
	Id          string
	Title       string
	Authors     string
	DateCreated string
}

type BooksViewModel struct {
	Books []*ListBookViewModel
}

func GetBooks(w http.ResponseWriter, r *http.Request) {

	myBooks, ok := rxa.GetAllValues(data.MyBooksIdsProperty, data.MyBooksIdsProperty)
	if !ok {
		http.Error(w, nod.ErrorStr("no my books found"), http.StatusInternalServerError)
		return
	}

	bvm := &BooksViewModel{
		Books: make([]*ListBookViewModel, 0, len(myBooks)),
	}

	for _, id := range myBooks {
		title, _ := rxa.GetFirstVal(data.TitleProperty, id)
		authors, _ := rxa.GetAllUnchangedValues(data.AuthorsProperty, id)
		created, _ := rxa.GetFirstVal(data.DateCreatedProperty, id)

		bvm.Books = append(bvm.Books, &ListBookViewModel{
			Id:          id,
			Title:       title,
			Authors:     strings.Join(authors, ","),
			DateCreated: created,
		})
	}

	if err := tmpl.ExecuteTemplate(w, "books-page", bvm); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}
