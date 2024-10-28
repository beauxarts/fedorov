package compton_data

import (
	"github.com/beauxarts/fedorov/data"
	"net/url"
	"strings"
)

const (
	SearchNew     = "Новый"
	SearchBacklog = "Бэклог"
	SearchKids    = "Детские"
)

var SearchOrder = []string{
	SearchNew,
	SearchBacklog,
	SearchKids,
}

func SearchScopes() map[string]string {
	queries := make(map[string]string, len(SearchOrder))

	queries[SearchNew] = ""

	q := url.Values{}
	//q.Set(data.BookCompletedProperty, "false")
	//q.Set(data.BookTypeProperty, BookTypeText)
	q.Set(data.SortProperty, data.ArtsHistoryOrderProperty)
	queries[SearchBacklog] = q.Encode()

	q = url.Values{}
	//q.Set(data.GenresProperty, "детские,сказки")
	//q.Set(data.BookTypeProperty, strings.Join([]string{BookTypeText, BookTypePDF}, ","))
	//q.Set(data.SortProperty, data.DateCreatedProperty)
	q.Set(data.DescendingProperty, "true")
	queries[SearchKids] = q.Encode()

	return queries
}

func EncodeQuery(query map[string][]string) string {
	q := &url.Values{}
	for property, values := range query {
		q.Set(property, strings.Join(values, ", "))
	}

	return q.Encode()
}

func SearchScopeFromQuery(query map[string][]string) string {
	enq := EncodeQuery(query)

	searchScope := SearchNew
	for st, sq := range SearchScopes() {
		if sq == enq {
			searchScope = st
		}
	}

	return searchScope
}
