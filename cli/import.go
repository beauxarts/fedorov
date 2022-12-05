package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/wits"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

func ImportHandler(_ *url.URL) error {
	return Import()
}

func Import() error {

	ia := nod.Begin("importing books...")
	defer ia.End()

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), nil, data.ReduxProperties()...)
	if err != nil {
		return ia.EndWithError(err)
	}

	absImportFilename := data.AbsImportFilename()
	if _, err := os.Stat(absImportFilename); err != nil {
		return ia.EndWithError(err)
	}

	file, err := os.Open(absImportFilename)
	defer file.Close()

	if err != nil {
		return ia.EndWithError(err)
	}

	skv, err := wits.ReadSectionKeyValues(file)
	if err != nil {
		return ia.EndWithError(err)
	}

	// import the data, copy files, etc

	for idstr := range skv {

		if id, err := strconv.ParseInt(idstr, 10, 64); err == nil {

			// move cover into destination folder
			srcCoverFilename := filepath.Join(data.Pwd(), idstr+data.CoverExt)
			if _, err := os.Stat(srcCoverFilename); err == nil {
				absDestCoverFilename := data.AbsCoverPath(id)
				if err := os.Rename(srcCoverFilename, absDestCoverFilename); err != nil {
					return ia.EndWithError(err)
				}
			}

			// move download files into destination folder
			for _, link := range skv[idstr][data.DownloadLinksProperty] {
				_, relSrcFilename := filepath.Split(link)
				if _, err := os.Stat(relSrcFilename); err == nil {
					absDstFilename := data.AbsDownloadPath(id, relSrcFilename)

					absDstDir, _ := filepath.Split(absDstFilename)
					if _, err := os.Stat(absDstDir); os.IsNotExist(err) {
						if err := os.MkdirAll(absDstDir, 0755); err != nil {
							return ia.EndWithError(err)
						}
					}

					absSrcFilename := filepath.Join(data.Pwd(), relSrcFilename)
					if err := os.Rename(absSrcFilename, absDstFilename); err != nil {
						return ia.EndWithError(err)
					}
				}
			}
		}

		// replace redux values with imported data
		for prop, values := range skv[idstr] {
			if err := rxa.ReplaceValues(prop, idstr, values...); err != nil {
				return ia.EndWithError(err)
			}
		}

		// add imported book id to my books
		if err := rxa.AddVal(data.MyBooksIdsProperty, data.MyBooksIdsProperty, idstr); err != nil {
			return ia.EndWithError(err)
		}

		// set imported book as... imported - primarily to be able to recreate upon sync
		if err := rxa.AddVal(data.ImportedProperty, idstr, "true"); err != nil {
			return ia.EndWithError(err)
		}
	}

	// delete import declaration when all is complete
	if err := os.Remove(absImportFilename); err != nil {
		return ia.EndWithError(err)
	}

	ia.EndWithResult("done")

	return nil
}

func IsImported(id string, rxa kvas.ReduxAssets) bool {
	if val, ok := rxa.GetFirstVal(data.ImportedProperty, id); ok {
		return val == "true"
	}
	return false
}
