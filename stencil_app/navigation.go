package stencil_app

import "github.com/boggydigital/stencil"

const (
	NavLatestBooks = "Новинки"
	NavSearch      = "Поиск"
)

var NavItems = []string{NavLatestBooks, NavSearch}

var NavIcons = map[string]string{
	NavLatestBooks: stencil.IconStack,
	NavSearch:      stencil.IconSearch,
}

var NavHrefs = map[string]string{
	NavLatestBooks: BooksPath,
	NavSearch:      SearchPath,
}
