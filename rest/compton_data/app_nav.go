package compton_data

import "github.com/boggydigital/compton"

const (
	AppNavLatest = "Книги"
	AppNavSearch = "Поиск"
)

var AppNavOrder = []string{AppNavLatest, AppNavSearch}

var AppNavIcons = map[string]compton.Symbol{
	AppNavLatest: compton.Stack,
	AppNavSearch: compton.Search,
}

var AppNavLinks = map[string]string{
	AppNavLatest: "/latest",
	AppNavSearch: "/search",
}
