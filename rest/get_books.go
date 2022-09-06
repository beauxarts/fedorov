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

type ListBook struct {
	Id          string
	Title       string
	Authors     string
	DateCreated string
}

type Shelf struct {
	Books []*ListBook
}

func NewShelf(ids []string) *Shelf {
	shelf := &Shelf{
		Books: make([]*ListBook, 0, len(ids)),
	}

	for _, id := range ids {
		title, _ := rxa.GetFirstVal(data.TitleProperty, id)
		authors, _ := rxa.GetAllUnchangedValues(data.AuthorsProperty, id)
		created, _ := rxa.GetFirstVal(data.DateCreatedProperty, id)

		shelf.Books = append(shelf.Books, &ListBook{
			Id:          id,
			Title:       title,
			Authors:     strings.Join(authors, ","),
			DateCreated: created,
		})
	}

	return shelf
}

func GetBooks(w http.ResponseWriter, r *http.Request) {

	// GET /books

	myBooks, ok := rxa.GetAllValues(data.MyBooksIdsProperty, data.MyBooksIdsProperty)
	if !ok {
		http.Error(w, nod.ErrorStr("no my books found"), http.StatusInternalServerError)
		return
	}

	shelf := NewShelf(myBooks)

	if err := tmpl.ExecuteTemplate(w, "books-page", shelf); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}
