package stencil_app

import (
	"github.com/beauxarts/fedorov/data"
	"net/url"
	"strings"
)

const (
	ScopeNewSearch      = "Новый поиск"
	ScopeCompletedBooks = "Прочитанные"
	ScopeKidsBooks      = "Детские"
)

var SearchScopes = []string{
	ScopeNewSearch,
	ScopeCompletedBooks,
	ScopeKidsBooks,
}

func SearchScopeQueries() map[string]string {
	scopeUrls := make(map[string]string, len(SearchScopes))

	scopeUrls[ScopeNewSearch] = ""

	q := url.Values{}
	q.Set(data.BookCompletedProperty, "true")
	q.Set(data.SortProperty, "date-created")
	q.Set(data.DescendingProperty, "true")
	scopeUrls[ScopeCompletedBooks] = q.Encode()

	q = url.Values{}
	q.Set(data.GenresProperty, "детские,сказки")
	q.Set(data.BookTypeProperty, strings.Join([]string{BookTypeText, BookTypePDF}, ","))
	q.Set(data.SortProperty, "date-created")
	q.Set(data.DescendingProperty, "true")
	scopeUrls[ScopeKidsBooks] = q.Encode()

	return scopeUrls
}
