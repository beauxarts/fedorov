package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathways"
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

	force := u.Query().Has("force")
	skipExisting := u.Query().Has("skip-existing")

	return DownloadLitResCovers(skipExisting, force, ids...)
}

func DownloadLitResCovers(skipExisting, force bool, artsIds ...string) error {

	gca := nod.NewProgress("downloading LitRes covers...")
	defer gca.Done()

	if len(artsIds) == 0 {
		var err error
		artsIds, err = GetRecentArts(force)
		if err != nil {
			return err
		}
	}

	sizes := litres_integration.AllCoverSizes()

	gca.TotalInt(len(artsIds) * len(sizes))

	dc := dolo.DefaultClient

	absCoversDir, err := pathways.GetAbsDir(data.Covers)
	if err != nil {
		return err
	}

	for _, id := range artsIds {

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

			if err := dc.Download(cu, force, nil, absCoversDir, relFn); err != nil {
				gca.Error(err)
			}

			gca.Increment()
		}
	}

	return nil
}
