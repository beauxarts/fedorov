package cli

import (
	"net/url"

	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/boggydigital/coost"
)

func ImportCookiesHandler(u *url.URL) error {

	cookieStr := u.Query().Get("cookies")

	absCookiesFilename, err := data.AbsCookiesFilename()
	if err != nil {
		return err
	}

	return coost.Import(cookieStr, litres_integration.DefaultUrl(), absCookiesFilename)
}
