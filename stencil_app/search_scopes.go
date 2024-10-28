package stencil_app

import (
	"github.com/beauxarts/fedorov/data"
	"net/url"
	"strings"
)

const (
	ScopeNewSearch        = "Новый"
	ScopeBacklogTextBooks = "Бэклог"
	ScopeKidsBooks        = "Детские"
)

var SearchScopes = []string{
	ScopeNewSearch,
	ScopeBacklogTextBooks,
	ScopeKidsBooks,
}

func SearchScopeQueries() map[string]string {
	scopeUrls := make(map[string]string, len(SearchScopes))

	scopeUrls[ScopeNewSearch] = ""

	q := url.Values{}
	//q.Set(data.BookCompletedProperty, "false")
	//q.Set(data.BookTypeProperty, BookTypeText)
	q.Set(data.SortProperty, data.ArtsHistoryOrderProperty)
	scopeUrls[ScopeBacklogTextBooks] = strings.ToLower(q.Encode())

	q = url.Values{}
	//q.Set(data.GenresProperty, "детские,сказки")
	//q.Set(data.BookTypeProperty, strings.Join([]string{BookTypeText, BookTypePDF}, ","))
	//q.Set(data.SortProperty, data.DateCreatedProperty)
	q.Set(data.DescendingProperty, "true")
	scopeUrls[ScopeKidsBooks] = strings.ToLower(q.Encode())

	return scopeUrls
}
