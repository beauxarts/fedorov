package cli

import (
	"errors"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/view_models"
	"github.com/beauxarts/litres_integration"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"net/http"
	"path/filepath"
	"strings"
)

var skipFormatDownloads = map[string]bool{
	// download
	view_models.FormatEPUB: false,
	view_models.FormatAZW3: false,
	view_models.FormatMOBI: false, // should be replaced by AZW3 - hasn't happened yet
	view_models.FormatFB2:  false,
	view_models.FormatMP4:  false,
	view_models.FormatZIP:  false,
	// don't download
	view_models.FormatPDFA4:   true,
	view_models.FormatPDFA6:   true,
	view_models.FormatTXT:     true,
	view_models.FormatFB3:     true,
	view_models.FormatRTF:     true,
	view_models.FormatTXTZIP:  true,
	view_models.FormatIOSEPUB: true,
	view_models.FormatMP3:     true,
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

			if f := view_models.LinkFormat(link); skipFormatDownloads[f] {
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
