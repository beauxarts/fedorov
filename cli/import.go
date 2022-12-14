package cli

import (
	"errors"
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/litres_integration"
	"github.com/boggydigital/coost"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/wits"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
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

	hc, err := coost.NewHttpClientFromFile(data.AbsCookiesFilename(), litres_integration.LitResHost)
	if err != nil {
		return err
	}

	// import the data, copy files, etc

	for idstr := range skv {

		// import additional data from external data source
		if ds, ok := skv[idstr][data.DataSourceProperty]; ok {
			if len(ds) == 0 {
				break
			}
			switch ds[0] {
			case "litres":

				if hrefs, ok := skv[idstr][data.HrefProperty]; ok {
					rxa.ReplaceValues(data.HrefProperty, idstr, hrefs...)
				} else {
					return ia.EndWithError(errors.New("href is required for litres data import"))
				}

				if rdx, err := importLitresData(idstr, hc); err != nil {
					return ia.EndWithError(err)
				} else {
					// rdx -> skv
					for p := range rdx {
						if p == data.DownloadLinksProperty {
							continue
						}
						if len(rdx[p][idstr]) == 0 {
							continue
						}
						skv[idstr][p] = rdx[p][idstr]
					}
				}
			default:
				//unknown data source - ignore
			}
		}

		if id, err := strconv.ParseInt(idstr, 10, 64); err == nil {

			// move cover into destination folder
			srcCoverFilename := filepath.Join(data.Pwd(), idstr+data.CoverExt)
			if _, err := os.Stat(srcCoverFilename); err == nil {
				absDestCoverFilename := data.AbsCoverPath(id, litres_integration.SizeMax)
				if err := os.Rename(srcCoverFilename, absDestCoverFilename); err != nil {
					return ia.EndWithError(err)
				}
			}

			// move download files into destination folder
			for _, link := range skv[idstr][data.DownloadLinksProperty] {
				_, relSrcFilename := filepath.Split(link)
				absSrcFilename := filepath.Join(data.Pwd(), relSrcFilename)
				if _, err := os.Stat(absSrcFilename); err == nil {
					absDstFilename := data.AbsDownloadPath(id, relSrcFilename)

					absDstDir, _ := filepath.Split(absDstFilename)
					if _, err := os.Stat(absDstDir); os.IsNotExist(err) {
						if err := os.MkdirAll(absDstDir, 0755); err != nil {
							return ia.EndWithError(err)
						}
					}

					if err := os.Rename(absSrcFilename, absDstFilename); err != nil {
						return ia.EndWithError(err)
					}
				} else {
					nod.Log(err.Error())
				}
			}
		}

		// replace redux values with imported data
		for prop, values := range skv[idstr] {
			if len(values) == 0 {
				continue
			}
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

	// move import declaration to imported dir when all is done
	if _, err := os.Stat(data.AbsImportedDir()); os.IsNotExist(err) {
		if err := os.MkdirAll(data.AbsImportedDir(), 0755); err != nil {
			return ia.EndWithError(err)
		}
	}

	_, dstImportFilename := filepath.Split(absImportFilename)
	dstImportFilename = fmt.Sprintf("%s_%s", time.Now().Format("20060102-1504"), dstImportFilename)
	dstImportFilename = filepath.Join(data.AbsImportedDir(), dstImportFilename)
	if err := os.Rename(absImportFilename, dstImportFilename); err != nil {
		return ia.EndWithError(err)
	}

	ia.EndWithResult("done")

	return nil
}

func importLitresData(id string, hc *http.Client) (map[string]map[string][]string, error) {

	ilda := nod.Begin("importing data from LitRes...")
	defer ilda.End()

	if err := GetDetails([]string{id}, hc); err != nil {
		return nil, ilda.EndWithError(err)
	}

	if err := GetCovers([]string{id}, true); err != nil {
		return nil, ilda.EndWithError(err)
	}

	kv, err := kvas.ConnectLocal(data.AbsMyBooksDetailsDir(), kvas.HtmlExt)
	if err != nil {
		return nil, ilda.EndWithError(err)
	}

	lrdx, err := ReduceBookDetails(id, kv)
	if err != nil {
		return nil, ilda.EndWithError(err)
	}

	rdx := make(map[string]map[string][]string)
	for _, p := range data.ReduxProperties() {
		rdx[p] = make(map[string][]string)
	}

	MapLitresToFedorov(id, lrdx, rdx)

	return rdx, nil
}

func IsImported(id string, rxa kvas.ReduxAssets) bool {
	if val, ok := rxa.GetFirstVal(data.ImportedProperty, id); ok {
		return val == "true"
	}
	return false
}
