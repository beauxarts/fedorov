package cli

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

func PurgeHandler(u *url.URL) error {
	id := u.Query().Get("id")
	confirm := u.Query().Has("confirm")
	return Purge(id, confirm)
}

// Purge will remove all book artefacts from the system:
// - details
// - covers
// - downloads
// - reductions (must be last to allow downloads to be resolved)
func Purge(id string, confirm bool) error {

	wa := nod.Begin("purge removes all book data, restoring that data will require an earlier backup")
	wa.End()

	pa := nod.Begin("purging book %s data...", id)
	defer pa.End()

	// covers

	var idi int64 = 0
	if pid, err := strconv.ParseInt(id, 10, 64); err == nil {
		idi = pid
	}

	for _, cs := range data.CoverSizesAsc {
		cfn := data.AbsCoverPath(idi, cs)
		if _, err := os.Stat(cfn); err == nil {
			rca := nod.Begin(" found cover %s...", filepath.Base(cfn))
			if confirm {
				if err := os.Remove(cfn); err != nil {
					return rca.EndWithError(err)
				}
				rca.EndWithResult("removed")
			}
			rca.End()
		}
	}

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), data.ReduxProperties()...)
	if err != nil {
		return pa.EndWithError(err)
	}

	// downloads

	if links, ok := rxa.GetAllValues(data.DownloadLinksProperty, id); ok {
		for _, link := range links {
			lfn := data.AbsDownloadPath(idi, filepath.Base(link))
			if _, err := os.Stat(lfn); err == nil {
				rda := nod.Begin(" found download %s...", filepath.Base(lfn))
				if confirm {
					if err := os.Remove(lfn); err != nil {
						return rda.EndWithError(err)
					}
					rda.EndWithResult("removed")
				}
				rda.End()
			}
		}
	}

	// details

	detailsDirs := []string{
		data.AbsLitResMyBooksFreshDir(),
		data.AbsLitResMyBooksDetailsDir(),
		data.AbsLiveLibDetailsDir(),
	}

	for _, d := range detailsDirs {
		kv, err := kvas.ConnectLocal(d, kvas.HtmlExt)
		if err != nil {
			return pa.EndWithError(err)
		}

		if kv.Has(id) {
			cda := nod.Begin(" found %s in %s...", id, filepath.Base(d))
			if confirm {
				if _, err := kv.Cut(id); err != nil {
					return cda.EndWithError(err)
				}
				cda.EndWithResult("removed")
			}
			cda.End()
		}
	}

	// reductions

	for _, p := range data.ReduxProperties() {
		if rxa.HasKey(p, id) {
			cra := nod.Begin(" found %s in %s...", id, p)
			if confirm {
				if values, ok := rxa.GetAllValues(p, id); ok {
					for _, val := range values {
						if err := rxa.CutVal(p, id, val); err != nil {
							return cra.EndWithError(err)
						}
					}
				}
				cra.EndWithResult("removed")
			}
			cra.End()
		}
	}

	result := fmt.Sprintf("run `purge %s -confirm` to delete all", id)
	if confirm {
		result = "done"
	}
	pa.EndWithResult(result)

	return nil
}
