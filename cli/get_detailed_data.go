package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/litres_integration"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/nod"
	"net/url"
)

func GetDetailedDataHandler(u *url.URL) error {
	return GetDetailedData()
}

func GetDetailedData() error {
	gdda := nod.NewProgress("fetching detailed data...")
	defer gdda.End()

	dc := dolo.DefaultClient

	ddu, addrd := litres_integration.DetailedDataUrl(), data.AbsDetailedDataRemoteDir()

	if err := dc.Download(ddu, gdda, addrd); err != nil {
		return gdda.EndWithError(err)
	}

	gdda.EndWithResult("done")

	return nil
}
