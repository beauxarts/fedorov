package cli

import (
	"errors"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pasu"
	"net/url"
	"strconv"
	"strings"
)

func DownloadLitResCoversHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	forceImported := u.Query().Has("force-imported")

	return DownloadLitResCovers(ids, forceImported)
}

func DownloadLitResCovers(ids []string, forceImported bool) error {

	gca := nod.NewProgress("downloading LitRes covers...")
	defer gca.End()

	rdx, err := data.NewReduxReader(
		data.ArtsHistoryOrderProperty,
		data.ImportedProperty)
	if err != nil {
		return gca.EndWithError(err)
	}

	if len(ids) == 0 {
		var ok bool
		ids, ok = rdx.GetAllValues(data.ArtsHistoryOrderProperty, data.ArtsHistoryOrderProperty)
		if !ok {
			err = errors.New("no arts history order found")
			return gca.EndWithError(err)
		}
	}

	sizes := litres_integration.AllCoverSizes()

	gca.TotalInt(len(ids) * len(sizes))

	dc := dolo.DefaultClient

	absCoversDir, err := pasu.GetAbsDir(data.Covers)
	if err != nil {
		return gca.EndWithError(err)
	}

	for _, id := range ids {

		// don't attempt downloading covers for imported books
		if !forceImported && IsImported(id, rdx) {
			continue
		}

		idn, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			gca.Error(err)
			gca.ProgressInt(len(sizes))
			continue
		}

		for _, size := range sizes {
			fn := data.RelCoverFilename(id, size)
			cu := litres_integration.CoverUrl(idn, size)

			if err := dc.Download(cu, nil, absCoversDir, fn); err != nil {
				gca.Error(err)
			}

			gca.Increment()
		}
	}

	gca.EndWithResult("done")

	return nil
}
