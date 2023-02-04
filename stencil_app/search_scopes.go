package stencil_app

import (
	"github.com/beauxarts/fedorov/data"
	"net/url"
	"strings"
)

const (
	ScopeNewSearch     = "Новый"
	ScopeUnreadBooks   = "Не прочитанные"
	ScopeKidsBooks     = "Детские"
	ScopeImportedBooks = "Импорт"
)

var SearchScopes = []string{
	ScopeNewSearch,
	ScopeUnreadBooks,
	ScopeKidsBooks,
	ScopeImportedBooks,
}

func SearchScopeQueries() map[string]string {
	scopeUrls := make(map[string]string, len(SearchScopes))

	scopeUrls[ScopeNewSearch] = ""

	q := url.Values{}
	q.Set(data.BookCompletedProperty, "false")
	q.Set(data.BookTypeProperty, BookTypeText)
	q.Set(data.SortProperty, data.MyBooksOrderProperty)
	scopeUrls[ScopeUnreadBooks] = q.Encode()

	q = url.Values{}
	q.Set(data.GenresProperty, "детские,сказки")
	q.Set(data.BookTypeProperty, strings.Join([]string{BookTypeText, BookTypePDF}, ","))
	q.Set(data.SortProperty, data.DateCreatedProperty)
	q.Set(data.DescendingProperty, "true")
	scopeUrls[ScopeKidsBooks] = q.Encode()

	q = url.Values{}
	q.Set(data.ImportedProperty, "true")
	q.Set(data.SortProperty, data.DateCreatedProperty)
	q.Set(data.DescendingProperty, "true")
	scopeUrls[ScopeImportedBooks] = q.Encode()

	return scopeUrls
}
