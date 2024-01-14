package cli

import (
	"errors"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pasu"
	"net/url"
	"strconv"
	"strings"
)

func GetLitResCoversHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	return GetLitResCovers(ids, false)
}

func GetLitResCovers(ids []string, forceImported bool) error {

	gca := nod.NewProgress("fetching LitRes covers...")
	defer gca.End()

	absReduxDir, err := pasu.GetAbsRelDir(data.Redux)
	if err != nil {
		return gca.EndWithError(err)
	}

	rdx, err := kvas.NewReduxReader(absReduxDir,
		data.MyBooksIdsProperty,
		data.ImportedProperty)
	if err != nil {
		return gca.EndWithError(err)
	}

	if len(ids) == 0 {
		var ok bool
		ids, ok = rdx.GetAllValues(data.MyBooksIdsProperty, data.MyBooksIdsProperty)
		if !ok {
			err = errors.New("no my books found")
			return gca.EndWithError(err)
		}
	}

	gca.TotalInt(len(ids))

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
			gca.Increment()
			continue
		}

		sizes := []litres_integration.CoverSize{
			litres_integration.Size330,
			litres_integration.Size415,
			litres_integration.SizeMax,
		}

		for _, size := range sizes {
			fn := data.RelCoverFilename(id, size)
			cu := litres_integration.CoverUrl(idn, size)

			if err := dc.Download(cu, nil, absCoversDir, fn); err != nil {
				//attempting partner url if default url fails
				pcu := litres_integration.PartnerCoverUrl(idn, size)
				if err := dc.Download(pcu, nil, absCoversDir, fn); err != nil {
					gca.Error(err)
					gca.Increment()
					continue
				}
			}
		}

		gca.Increment()
	}

	gca.EndWithResult("done")

	return nil
}
