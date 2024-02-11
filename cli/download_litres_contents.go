package cli

import (
	"net/url"
	"strings"
)

func DownloadLitResContentsHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	return DownloadLitresContents(ids...)
}

func DownloadLitresContents(ids ...string) error {
	return nil
}
