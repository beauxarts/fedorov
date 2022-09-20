package stencil_app

const (
	NavBooks  = "Книги"
	NavSearch = "Поиск"
)

var NavItems = []string{NavBooks, NavSearch}

var NavIcons = map[string]string{
	NavBooks:  "stack",
	NavSearch: "search",
}

var NavHrefs = map[string]string{
	NavBooks:  BooksPath,
	NavSearch: SearchPath,
}
