package cli

import (
	"errors"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/litres_integration"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"net/http"
	"path/filepath"
	"strings"
)

var unsupportedDownloadSuffixes = map[string]bool{
	".a4.pdf":   true,
	".a6.pdf":   true,
	".fb2.zip":  true,
	".txt":      true,
	".mp3.zip":  true,
	".fb3":      true,
	".rtf.zip":  true,
	".txt.zip":  true,
	".ios.epub": true,
}

func Download(hc *http.Client) error {

	da := nod.NewProgress("downloading books...")
	defer da.End()

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), nil,
		data.MyBooksIdsProperty,
		data.TitleProperty,
		data.AuthorsProperty,
		data.DownloadLinksProperty)

	if err != nil {
		return da.EndWithError(err)
	}

	ids, ok := rxa.GetAllUnchangedValues(data.MyBooksIdsProperty, data.MyBooksIdsProperty)
	if !ok {
		err = errors.New("no my books found")
		return da.EndWithError(err)
	}

	da.TotalInt(len(ids))

	dc := dolo.NewClient(hc, dolo.Defaults())

	for _, id := range ids {

		title, _ := rxa.GetFirstVal(data.TitleProperty, id)
		authors, _ := rxa.GetAllUnchangedValues(data.AuthorsProperty, id)

		dls, ok := rxa.GetAllUnchangedValues(data.DownloadLinksProperty, id)
		if !ok {
			nod.Log("book %s is missing download links", id)
			continue
		}

		bdla := nod.Begin("%s - %s", strings.Join(authors, ","), title)

		for _, link := range dls {
			if unsupportedLink(link) {
				continue
			}

			_, fname := filepath.Split(link)

			tpw := nod.NewProgress(" %s", fname)

			if err := dc.Download(litres_integration.HrefUrl(link), tpw, data.AbsDownloadsDir(), id, fname); err != nil {
				nod.Log(err.Error())
				continue
			}

			tpw.EndWithResult("done")
		}

		bdla.End()
		da.Increment()
	}

	da.EndWithResult("done")

	return nil
}

func unsupportedLink(link string) bool {
	for slp := range unsupportedDownloadSuffixes {
		if strings.HasSuffix(link, slp) {
			return true
		}
	}
	return false
}
