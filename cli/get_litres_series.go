package cli

import (
	"net/url"
	"strings"
)

func GetLitResSeriesHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	return GetLitResSeries(ids)
}

func GetLitResSeries(ids []string) error {
	return nil
}
