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
	"github.com/boggydigital/pathology"
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

// Import adds external media to the library, e.g. DRM-free purchases from elsewhere.
// To do that you need to provide the following in the input directory:
// - BEFORE YOU START: book id should be a value that can be parsed into int64
//
// - book media files, e.g. `id.mp4`, `id.epub`, `id.txt`
// - book cover named `id.jpg`
// - import.txt with the media metadata
//   - id - ISBN is recommended or a LiveLib id
//   - title, ... - as needed
//   - data-source=litres - if the book (or similar book) is available on LitRes
//   - > href - e.g. `/book/mihail-shishkin/pismovnik-447855/`
//   - data-source=livelib - if the book (or similar book) is available on LiveLib
//   - > id of the book should match LiveLib id, e.g. `1003406901`
//   - download-links - filenames of the book media files, e.g. id.mp4, id.epub, id.txt
//   - download-titles - human-readable titles corresponding to the download-links
//   - > e.g. Audio-book (MP4), EPUB, Text file
//
// Upon placing those files, you can run import command. When import completes -
// all those input files are removed. import.txt is moved to the  _imported subdirectory
// of the input directory as YYYYMMDD-HHMM-import.txt, where YYYYMMDD-HHMM is the date
// and time of the import.
//
// If needed you can move YYYYMMDD-HHMM-import.txt back to the input directory, rename
// to import.txt, change some values and import again. New data will overwrite any
// existing data.
//
// Another suggestion is to consider exporting the book with similar metadata to use as a
// base for import.txt. export.txt and import.txt have identical structure.
func Import() error {

	ia := nod.Begin("importing books...")
	defer ia.End()

	absReduxDir, err := pathology.GetAbsRelDir(data.Redux)
	if err != nil {
		return ia.EndWithError(err)
	}

	rdx, err := kvas.NewReduxWriter(absReduxDir, data.ReduxProperties()...)
	if err != nil {
		return ia.EndWithError(err)
	}

	absImportFilename, err := data.AbsImportFilename()
	if err != nil {
		return ia.EndWithError(err)
	}

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

	absCookiesFilename, err := data.AbsCookiesFilename()
	if err != nil {
		return ia.EndWithError(err)
	}

	hc, err := coost.NewHttpClientFromFile(absCookiesFilename)
	if err != nil {
		return ia.EndWithError(err)
	}

	absInputDir, err := pathology.GetAbsDir(data.Input)
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
					if err := rdx.ReplaceValues(data.HrefProperty, idstr, hrefs...); err != nil {
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
			srcCoverFilename := filepath.Join(absInputDir, idstr+data.DefaultCoverExt)
			if _, err := os.Stat(srcCoverFilename); err == nil {
				absDestCoverFilename, err := data.AbsCoverImagePath(id, litres_integration.SizeMax)
				if err != nil {
					return ia.EndWithError(err)
				}
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
				absSrcFilename := filepath.Join(absInputDir, relSrcFilename)
				if _, err := os.Stat(absSrcFilename); err == nil {
					absDstFilename, err := data.AbsFileDownloadPath(id, relSrcFilename)
					if err != nil {
						return ia.EndWithError(err)
					}

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
			if err := rdx.ReplaceValues(prop, idstr, values...); err != nil {
				return ia.EndWithError(err)
			}
		}

		// add imported book id to my books
		if err := rdx.AddValues(data.MyBooksIdsProperty, data.MyBooksIdsProperty, idstr); err != nil {
			return ia.EndWithError(err)
		}

		// set imported book as... imported - primarily to be able to recreate upon sync
		if err := rdx.AddValues(data.ImportedProperty, idstr, "true"); err != nil {
			return ia.EndWithError(err)
		}
	}

	absImportedDir, err := pathology.GetAbsDir(data.Imported)
	if err != nil {
		return ia.EndWithError(err)
	}

	// move import declaration to imported dir when all is done
	if _, err := os.Stat(absImportedDir); os.IsNotExist(err) {
		if err := os.MkdirAll(absImportedDir, 0755); err != nil {
			return ia.EndWithError(err)
		}
	}

	_, dstImportFilename := filepath.Split(absImportFilename)
	dstImportFilename = fmt.Sprintf("%s_%s", time.Now().Format("20060102-1504"), dstImportFilename)
	dstImportFilename = filepath.Join(absImportedDir, dstImportFilename)
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

	absLitResMyBooksDetailsDir, err := data.AbsDataTypeDir(data.LitResMyBooksDetails)
	if err != nil {
		return nil, ilda.EndWithError(err)
	}

	kv, err := kvas.ConnectLocal(absLitResMyBooksDetailsDir, kvas.HtmlExt)
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

	absLiveLibDetailsDir, err := data.AbsDataTypeDir(data.LiveLibDetails)
	if err != nil {
		return nil, ilda.EndWithError(err)
	}

	kv, err := kvas.ConnectLocal(absLiveLibDetailsDir, kvas.HtmlExt)
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

func IsImported(id string, rdx kvas.ReadableRedux) bool {
	if val, ok := rdx.GetFirstVal(data.ImportedProperty, id); ok {
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
