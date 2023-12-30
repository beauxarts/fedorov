package cli

import (
	"net/url"
	"strings"
)

func GetLitResArtsHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	return GetLitResArts(ids)
}

func GetLitResArts(ids []string) error {
	return nil
}
