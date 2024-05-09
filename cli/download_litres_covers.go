package cli

import (
	"errors"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pasu"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func DownloadLitResCoversHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	forceImported := u.Query().Has("force-imported")
	skipExisting := u.Query().Has("skip-existing")

	return DownloadLitResCovers(skipExisting, forceImported, ids...)
}

func DownloadLitResCovers(skipExisting, forceImported bool, ids ...string) error {

	gca := nod.NewProgress("downloading LitRes covers...")
	defer gca.End()

	rdx, err := data.NewReduxReader(
		data.ArtsHistoryOrderProperty)
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

		idn, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			gca.Error(err)
			gca.ProgressInt(len(sizes))
			continue
		}

		for _, size := range sizes {
			relFn := data.RelCoverFilename(id, size)
			cu := litres_integration.CoverUrl(idn, size)

			if skipExisting {
				absFn := filepath.Join(absCoversDir, relFn)
				if _, err := os.Stat(absFn); err == nil {
					gca.Increment()
					continue
				}
			}

			if err := dc.Download(cu, nil, absCoversDir, relFn); err != nil {
				gca.Error(err)
			}

			gca.Increment()
		}
	}

	gca.EndWithResult("done")

	return nil
}
