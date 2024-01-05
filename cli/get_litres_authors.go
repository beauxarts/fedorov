package cli

import (
	"net/url"
	"strings"
)

func GetLitResAuthorsHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	return GetLitResAuthors(ids)
}

func GetLitResAuthors(ids []string) error {
	return nil
}
