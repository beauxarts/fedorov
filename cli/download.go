package cli

import (
	"errors"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/litres_integration"
	"github.com/boggydigital/coost"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

var skipFormatDownloads = map[string]bool{
	// download
	data.FormatEPUB: false,
	data.FormatAZW3: false,
	data.FormatMOBI: false, // should be replaced by AZW3 - hasn't happened yet
	data.FormatMP4:  false,
	data.FormatZIP:  false,
	// don't download
	data.FormatPDFA4:   true,
	data.FormatPDFA6:   true,
	data.FormatTXT:     true,
	data.FormatFB2:     true,
	data.FormatFB3:     true,
	data.FormatRTF:     true,
	data.FormatTXTZIP:  true,
	data.FormatIOSEPUB: true,
	data.FormatMP3:     true,
}

func DownloadHandler(u *url.URL) error {
	hc, err := coost.NewHttpClientFromFile(data.AbsCookiesFilename(), litres_integration.LitResHost)
	if err != nil {
		return err
	}

	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	return Download(ids, hc)
}

func Download(ids []string, hc *http.Client) error {

	da := nod.NewProgress("downloading books...")
	defer da.End()

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), nil,
		data.MyBooksIdsProperty,
		data.TitleProperty,
		data.AuthorsProperty,
		data.ImportedProperty,
		data.DownloadLinksProperty)

	if err != nil {
		return da.EndWithError(err)
	}

	if ids == nil {
		var ok bool
		ids, ok = rxa.GetAllUnchangedValues(data.MyBooksIdsProperty, data.MyBooksIdsProperty)
		if !ok {
			err = errors.New("no my books found")
			return da.EndWithError(err)
		}
	}

	da.TotalInt(len(ids))

	dc := dolo.NewClient(hc, dolo.Defaults())

	for _, id := range ids {

		// don't attempt downloading imported books
		if IsImported(id, rxa) {
			continue
		}

		title, _ := rxa.GetFirstVal(data.TitleProperty, id)
		authors, _ := rxa.GetAllUnchangedValues(data.AuthorsProperty, id)

		dls, ok := rxa.GetAllUnchangedValues(data.DownloadLinksProperty, id)
		if !ok {
			nod.Log("book %s is missing download links", id)
			continue
		}

		bdla := nod.Begin("%s - %s", strings.Join(authors, ","), title)

		for _, link := range dls {

			if f := data.LinkFormat(link); skipFormatDownloads[f] {
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
