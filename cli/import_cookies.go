package cli

import (
	"net/url"

	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/boggydigital/coost"
	"github.com/boggydigital/nod"
)

func ImportCookiesHandler(u *url.URL) error {

	ica := nod.Begin("importing cookies...")
	defer ica.Done()

	cookieStr := u.Query().Get("cookies")

	absCookiesFilename, err := data.AbsCookiesFilename()
	if err != nil {
		return err
	}

	return coost.Import(cookieStr, litres_integration.DefaultUrl(), absCookiesFilename)
}
