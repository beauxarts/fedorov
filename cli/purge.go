package cli

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/beauxarts/scrinium/livelib_integration"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pasu"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

func PurgeHandler(u *url.URL) error {
	id := u.Query().Get("id")
	confirm := u.Query().Has("confirm")
	wu := u.Query().Get("webhook-url")
	return Purge(id, wu, confirm)
}

// Purge will remove all book artefacts from the system:
// - details
// - covers
// - downloads
// - reductions (must be last to allow downloads to be resolved)
func Purge(id string, webhookUrl string, confirm bool) error {

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
		cfn, err := data.AbsCoverImagePath(idi, cs)
		if err != nil {
			return pa.EndWithError(err)
		}
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

	props := data.ReduxProperties()
	props = append(props, data.ImportedProperties()...)

	absReduxDir, err := pasu.GetAbsRelDir(data.Redux)
	if err != nil {
		return pa.EndWithError(err)
	}

	rdx, err := kvas.NewReduxWriter(absReduxDir, props...)
	if err != nil {
		return pa.EndWithError(err)
	}

	// downloads

	//if links, ok := rdx.GetAllValues(data.DownloadLinksProperty, id); ok {
	//	for _, link := range links {
	//		lfn, err := data.AbsFileDownloadPath(idi, filepath.Base(link))
	//		if err != nil {
	//			return pa.EndWithError(err)
	//		}
	//		if _, err := os.Stat(lfn); err == nil {
	//			rda := nod.Begin(" found download %s...", filepath.Base(lfn))
	//			if confirm {
	//				if err := os.Remove(lfn); err != nil {
	//					return rda.EndWithError(err)
	//				}
	//				rda.EndWithResult("removed")
	//			}
	//			rda.End()
	//		}
	//	}
	//}

	// details

	//TODO: Likely need to add Arts data types
	dataTypes := []fmt.Stringer{
		// LitResDataTypes
		litres_integration.LitResHistoryLog,

		// LiveLibDataTypes
		livelib_integration.LiveLibDetails,

		// ArtsType
		litres_integration.ArtsTypeDetails,
		litres_integration.ArtsTypeSimilar,
		litres_integration.ArtsTypeQuotes,
		litres_integration.ArtsTypeFiles,
		litres_integration.ArtsTypeReviews,

		// AuthorTypes
		litres_integration.AuthorDetails,
		litres_integration.AuthorArts,

		// SeriesTypes
		litres_integration.SeriesDetails,
		litres_integration.SeriesArts,
	}

	for _, dt := range dataTypes {

		adtd, err := data.AbsDataTypeDir(dt)
		if err != nil {
			return pa.EndWithError(err)
		}

		kv, err := kvas.ConnectLocal(adtd, kvas.HtmlExt)
		if err != nil {
			return pa.EndWithError(err)
		}

		if kv.Has(id) {
			cda := nod.Begin(" found %s in %s...", id, filepath.Base(adtd))
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

	for _, p := range props {
		if rdx.HasKey(p, id) {
			cra := nod.Begin(" found %s in %s...", id, p)
			if confirm {
				if values, ok := rdx.GetAllValues(p, id); ok {
					for _, val := range values {
						if err := rdx.CutValues(p, id, val); err != nil {
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
