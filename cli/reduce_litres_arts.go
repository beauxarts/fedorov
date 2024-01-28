package cli

import (
	"github.com/beauxarts/scrinium/litres_integration"
	"net/url"
	"strings"
)

func ReduceLitResArtsHandler(u *url.URL) error {

	var artsTypesStr []string
	if ats := u.Query().Get("arts-type"); ats != "" {
		artsTypesStr = strings.Split(ats, ",")
	}

	artsTypes := make([]litres_integration.ArtsType, 0, len(artsTypesStr))

	if allArtsTypes := u.Query().Has("all-arts-types"); allArtsTypes {
		artsTypes = litres_integration.AllArtsTypes()
	} else {
		for _, ats := range artsTypesStr {
			artsTypes = append(artsTypes, litres_integration.ParseArtsType(ats))
		}
	}

	return ReduceLitResArts(artsTypes)
}

func ReduceLitResArts(artsTypes []litres_integration.ArtsType) error {
	return nil
}
