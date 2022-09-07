package view_models

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kvas"
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

func NewShelf(ids []string, rxa kvas.ReduxAssets) *Shelf {
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
			Authors:     strings.Join(authors, ", "),
			DateCreated: created,
		})
	}

	return shelf
}
