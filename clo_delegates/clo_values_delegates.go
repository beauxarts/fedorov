package clo_delegates

import (
	"fmt"
	"github.com/beauxarts/fedorov/litres_integration"
)

var Values = map[string]func() []string{
	"arts-types":   allArtsTypesStrings,
	"author-types": allAuthorTypesStrings,
	"series-types": allSeriesTypesStrings,
}

func toStrings[T fmt.Stringer](stringers ...T) []string {
	strings := make([]string, 0, len(stringers))
	for _, s := range stringers {
		strings = append(strings, s.String())
	}
	return strings
}

func allArtsTypesStrings() []string {
	return toStrings(litres_integration.AllArtsTypes()...)
}

func allAuthorTypesStrings() []string {
	return toStrings(litres_integration.AllAuthorTypes()...)
}

func allSeriesTypesStrings() []string {
	return toStrings(litres_integration.AllSeriesTypes()...)
}
