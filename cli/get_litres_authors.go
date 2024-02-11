package cli

import (
	"github.com/beauxarts/scrinium/litres_integration"
	"net/url"
	"strings"
)

func GetLitResAuthorsHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	var authorTypesStr []string
	if ats := u.Query().Get("author-type"); ats != "" {
		authorTypesStr = strings.Split(ats, ",")
	}

	authorTypes := make([]litres_integration.AuthorType, 0, len(authorTypesStr))

	if allAuthorTypes := u.Query().Has("all-author-types"); allAuthorTypes {
		authorTypes = litres_integration.AllAuthorTypes()
	} else {
		for _, ats := range authorTypesStr {
			authorTypes = append(authorTypes, litres_integration.ParseAuthorType(ats))
		}
	}

	force := u.Query().Has("force")

	return GetLitResAuthors(authorTypes, force, ids...)
}

func GetLitResAuthors(authorTypes []litres_integration.AuthorType, force bool, ids ...string) error {
	return nil
}
