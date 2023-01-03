package cli

import (
	"errors"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/litres_integration"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"net/url"
	"strconv"
	"strings"
)

func GetCoversHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	return GetCovers(ids, false)
}

func GetCovers(ids []string, forceImported bool) error {

	gca := nod.NewProgress("fetching covers...")
	defer gca.End()

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), nil,
		data.MyBooksIdsProperty,
		data.ImportedProperty)
	if err != nil {
		return gca.EndWithError(err)
	}

	if len(ids) == 0 {
		var ok bool
		ids, ok = rxa.GetAllUnchangedValues(data.MyBooksIdsProperty, data.MyBooksIdsProperty)
		if !ok {
			err = errors.New("no my books found")
			return gca.EndWithError(err)
		}
	}

	gca.TotalInt(len(ids))

	dc := dolo.DefaultClient

	for _, id := range ids {

		// don't attempt downloading covers for imported books
		if !forceImported && IsImported(id, rxa) {
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

			if err := dc.Download(cu, nil, data.AbsCoverDir(), fn); err != nil {
				//attempting partner url if default url fails
				pcu := litres_integration.PartnerCoverUrl(idn, size)
				if err := dc.Download(pcu, nil, data.AbsCoverDir(), fn); err != nil {
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
