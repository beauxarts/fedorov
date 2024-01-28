package cli

import (
	"errors"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/coost"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pasu"
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

func DownloadLitResHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	return DownloadLitRes(ids)
}

func DownloadLitRes(ids []string) error {

	da := nod.NewProgress("downloading LitRes books...")
	defer da.End()

	absReduxDir, err := pasu.GetAbsRelDir(data.Redux)
	if err != nil {
		return da.EndWithError(err)
	}

	rdx, err := kvas.NewReduxReader(absReduxDir,
		data.ArtsHistoryOrderProperty,
		data.TitleProperty,
		data.AuthorsProperty,
		data.ImportedProperty,
		data.DownloadLinksProperty)

	if err != nil {
		return da.EndWithError(err)
	}

	if ids == nil {
		var ok bool
		ids, ok = rdx.GetAllValues(data.ArtsHistoryOrderProperty, data.ArtsHistoryOrderProperty)
		if !ok {
			err = errors.New("no my books found")
			return da.EndWithError(err)
		}
	}

	da.TotalInt(len(ids))

	absCookiesFilename, err := data.AbsCookiesFilename()
	if err != nil {
		return da.EndWithError(err)
	}

	cj, err := coost.NewJar(absCookiesFilename)
	if err != nil {
		return da.EndWithError(err)
	}

	hc := cj.NewHttpClient()

	dc := dolo.NewClient(hc, dolo.Defaults())

	absDownloadsDir, err := pasu.GetAbsDir(data.Downloads)
	if err != nil {
		return da.EndWithError(err)
	}

	for _, id := range ids {

		// don't attempt downloading imported books
		if IsImported(id, rdx) {
			continue
		}

		title, _ := rdx.GetFirstVal(data.TitleProperty, id)
		authors, _ := rdx.GetAllValues(data.AuthorsProperty, id)

		dls, ok := rdx.GetAllValues(data.DownloadLinksProperty, id)
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

			if err := dc.Download(litres_integration.HrefUrl(link), tpw, absDownloadsDir, id, fname); err != nil {
				nod.Log(err.Error())
				continue
			}

			if err := cj.Store(absCookiesFilename); err != nil {
				return da.EndWithError(err)
			}

			tpw.EndWithResult("done")
		}

		bdla.End()
		da.Increment()
	}

	da.EndWithResult("done")

	return nil
}
