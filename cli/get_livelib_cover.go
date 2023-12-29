package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathology"
	"net/url"
)

func GetLiveLibCover(id, src string) error {

	gca := nod.NewProgress("fetching LitRes covers...")
	defer gca.End()

	dc := dolo.DefaultClient

	fn := data.RelCoverFilename(id, litres_integration.SizeMax)

	cu, err := url.Parse(src)
	if err != nil {
		return gca.EndWithError(err)
	}

	absCoverDir, err := pathology.GetAbsDir(data.Covers)
	if err != nil {
		return gca.EndWithError(err)
	}

	if err := dc.Download(cu, nil, absCoverDir, fn); err != nil {
		return gca.EndWithError(err)
	}

	return nil
}
