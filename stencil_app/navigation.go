package stencil_app

import "github.com/boggydigital/stencil"

const (
	NavBooks  = "Книги"
	NavSearch = "Поиск"
)

var NavItems = []string{NavBooks, NavSearch}

var NavIcons = map[string]string{
	NavBooks:  stencil.IconStack,
	NavSearch: stencil.IconSearch,
}

var NavHrefs = map[string]string{
	NavBooks:  BooksPath,
	NavSearch: SearchPath,
}
