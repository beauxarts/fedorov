package compton_data

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"net/url"
	"strconv"
	"strings"
)

var (
	SearchNew   = "Новый"
	SearchText  = litres_integration.ArtTypeText.String()
	SearchAudio = litres_integration.ArtTypeAudio.String()
	SearchPDF   = litres_integration.ArtTypePDF.String()
)

var SearchOrder = []string{
	SearchNew,
	SearchText,
	SearchAudio,
	SearchPDF,
}

func SearchScopes() map[string]string {
	queries := make(map[string]string, len(SearchOrder))

	queries[SearchNew] = ""

	q := url.Values{}
	q.Set(data.ArtTypeProperty, strconv.Itoa(int(litres_integration.ArtTypeText)))
	q.Set(data.SortProperty, data.ArtsOperationsEventTimeProperty)
	q.Set(data.DescendingProperty, "true")
	queries[SearchText] = q.Encode()

	q = url.Values{}
	q.Set(data.ArtTypeProperty, strconv.Itoa(int(litres_integration.ArtTypeAudio)))
	q.Set(data.SortProperty, data.ArtsOperationsEventTimeProperty)
	q.Set(data.DescendingProperty, "true")
	queries[SearchAudio] = q.Encode()

	q = url.Values{}
	q.Set(data.ArtTypeProperty, strconv.Itoa(int(litres_integration.ArtTypePDF)))
	q.Set(data.SortProperty, data.ArtsOperationsEventTimeProperty)
	q.Set(data.DescendingProperty, "true")
	queries[SearchPDF] = q.Encode()

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
