package clo_delegates

import (
	"github.com/beauxarts/scrinium/litres_integration"
)

var Values = map[string]func() []string{
	"arts-types": allArtsTypesStr,
}

func artsTypeStr(artsTypes ...litres_integration.ArtsType) []string {
	atStr := make([]string, 0, len(artsTypes))
	for _, at := range artsTypes {
		atStr = append(atStr, at.String())
	}
	return atStr
}

func allArtsTypesStr() []string {
	return artsTypeStr(litres_integration.AllArtsTypes()...)
}
