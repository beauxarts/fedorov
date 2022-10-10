package stencil_app

import "github.com/beauxarts/fedorov/data"

var BooksProperties = []string{
	data.BookTypeProperty,
	data.BookCompletedProperty,
	data.LocalTagsProperty,
	data.AuthorsProperty,
	data.DateCreatedProperty,
}

var BooksLabels = []string{
	data.BookTypeProperty,
	data.BookCompletedProperty,
	data.LocalTagsProperty,
}
