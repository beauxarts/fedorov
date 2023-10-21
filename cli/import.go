package cli

import (
	"errors"
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/beauxarts/scrinium/livelib_integration"
	"github.com/boggydigital/coost"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/wits"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	LitResDataSource  = "litres"
	LiveLibDataSource = "livelib"
)

func ImportHandler(_ *url.URL) error {
	return Import()
}

func Import() error {

	ia := nod.Begin("importing books...")
	defer ia.End()

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), data.ReduxProperties()...)
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

	hc, err := coost.NewHttpClientFromFile(data.AbsCookiesFilename())
	if err != nil {
		return ia.EndWithError(err)
	}

	// import the data, copy files, etc

	for idstr := range skv {

		// import additional data from external data source
		if ds, ok := skv[idstr][data.DataSourceProperty]; ok {
			if len(ds) == 0 {
				break
			}
			switch ds[0] {
			case LitResDataSource:

				if hrefs, ok := skv[idstr][data.HrefProperty]; ok {
					if err := rxa.ReplaceValues(data.HrefProperty, idstr, hrefs...); err != nil {
						return ia.EndWithError(err)
					}
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

			case LiveLibDataSource:

				if rdx, err := importLiveLibData(idstr, hc); err != nil {
					return ia.EndWithError(err)
				} else {
					// rdx -> skv
					for p := range rdx {
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
					if strings.Contains(err.Error(), "invalid cross-device link") {
						if err := copyDelete(srcCoverFilename, absDestCoverFilename); err != nil {
							return ia.EndWithError(err)
						}
					} else {
						return ia.EndWithError(err)
					}
				}
			}

			// move download files into destination folder and infer book type

			for _, link := range skv[idstr][data.DownloadLinksProperty] {

				if len(skv[idstr][data.BookTypeProperty]) == 0 {
					skv[idstr][data.BookTypeProperty] = []string{data.FormatBookType(data.LinkFormat(link))}
				}

				_, relSrcFilename := filepath.Split(link)
				if relSrcFilename == "" {
					continue
				}
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
						if strings.Contains(err.Error(), "invalid cross-device link") {
							if err := copyDelete(absSrcFilename, absDstFilename); err != nil {
								return ia.EndWithError(err)
							}
						} else {
							return ia.EndWithError(err)
						}
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
		if strings.Contains(err.Error(), "invalid cross-device link") {
			if err := copyDelete(absImportFilename, dstImportFilename); err != nil {
				return ia.EndWithError(err)
			}
		} else {
			return ia.EndWithError(err)
		}
	}

	ia.EndWithResult("done")

	return nil
}

func importLitresData(id string, hc *http.Client) (map[string]map[string][]string, error) {

	ilda := nod.Begin("importing data from LitRes...")
	defer ilda.End()

	if err := GetLitResDetails([]string{id}, false, false); err != nil {
		return nil, ilda.EndWithError(err)
	}

	if err := GetLitResCovers([]string{id}, true); err != nil {
		return nil, ilda.EndWithError(err)
	}

	kv, err := kvas.ConnectLocal(data.AbsLitResMyBooksDetailsDir(), kvas.HtmlExt)
	if err != nil {
		return nil, ilda.EndWithError(err)
	}

	lrdx, err := ReduceLitResBookDetails(id, kv)
	if err != nil {
		return nil, ilda.EndWithError(err)
	}

	rdx := make(map[string]map[string][]string)
	for _, p := range data.ReduxProperties() {
		rdx[p] = make(map[string][]string)
	}

	MapLitResToFedorov(id, lrdx, rdx)

	return rdx, nil
}

func importLiveLibData(id string, hc *http.Client) (map[string]map[string][]string, error) {
	ilda := nod.Begin("importing data from LiveLib...")
	defer ilda.End()

	if err := GetLiveLibDetails([]string{id}, hc, false); err != nil {
		return nil, ilda.EndWithError(err)
	}

	kv, err := kvas.ConnectLocal(data.AbsLiveLibDetailsDir(), kvas.HtmlExt)
	if err != nil {
		return nil, ilda.EndWithError(err)
	}

	lrdx, err := ReduceLiveLibBookDetails(id, kv)
	if err != nil {
		return nil, ilda.EndWithError(err)
	}

	if srcs, ok := lrdx[livelib_integration.ImageProperty]; ok && len(srcs) > 0 {
		if err := GetLiveLibCover(id, srcs[0]); err != nil {
			return nil, ilda.EndWithError(err)
		}
	}

	rdx := make(map[string]map[string][]string)
	for _, p := range data.ReduxProperties() {
		rdx[p] = make(map[string][]string)
	}

	MapLiveLibToFedorov(id, lrdx, rdx)

	return rdx, nil
}

func IsImported(id string, rxa kvas.ReduxAssets) bool {
	if val, ok := rxa.GetFirstVal(data.ImportedProperty, id); ok {
		return val == "true"
	}
	return false
}

func copyDelete(src, dst string) error {
	srcf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcf.Close()

	dstf, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstf.Close()

	_, err = io.Copy(dstf, srcf)
	if err != nil {
		return err
	}

	return os.Remove(src)
}
