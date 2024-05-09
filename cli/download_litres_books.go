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
	"os"
	"path/filepath"
	"strings"
)

func DownloadLitResBooksHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	force := u.Query().Has("force")

	return DownloadLitResBooks(force, ids...)
}

func DownloadLitResBooks(force bool, ids ...string) error {

	da := nod.NewProgress("downloading LitRes books...")
	defer da.End()

	rdx, err := data.NewReduxReader(
		data.ArtsHistoryOrderProperty,
		data.TitleProperty,
		data.PersonsIdsProperty,
		data.PersonsRolesProperty,
		data.PersonFullNameProperty,
	)
	if err != nil {
		return da.EndWithError(err)
	}

	kv, err := data.NewArtsReader(litres_integration.ArtsTypeFiles)
	if err != nil {
		return da.EndWithError(err)
	}

	if ids == nil {
		var ok bool
		ids, ok = rdx.GetAllValues(data.ArtsHistoryOrderProperty, data.ArtsHistoryOrderProperty)
		if !ok {
			err = errors.New("no arts history order found")
			return da.EndWithError(err)
		}
	}

	da.TotalInt(len(ids))

	absCookiesFilename, err := data.AbsCookiesFilename()
	if err != nil {
		return da.EndWithError(err)
	}

	hc, err := coost.NewHttpClientFromFile(absCookiesFilename)
	if err != nil {
		return da.EndWithError(err)
	}

	dc := dolo.NewClient(hc, dolo.Defaults())

	absDownloadsDir, err := pasu.GetAbsDir(data.Downloads)
	if err != nil {
		return da.EndWithError(err)
	}

	for _, id := range ids {

		title, _ := rdx.GetFirstVal(data.TitleProperty, id)
		authorsNames, err := authorsFullNames(id, rdx)
		if err != nil {
			return da.EndWithError(err)
		}

		bdla := nod.Begin("%s %s - %s", id, strings.Join(authorsNames, ","), title)

		if !kv.Has(id) {
			continue
		}

		artFiles, err := kv.ArtsFiles(id)
		if err != nil {
			return da.EndWithError(err)
		}

		for _, afd := range artFiles.DownloadsTypes() {

			u := afd.Url(id)

			_, relFn := filepath.Split(u.Path)

			tpw := nod.NewProgress(" %s", relFn)

			if !force {
				absFn := filepath.Join(absDownloadsDir, id, relFn)
				if _, err := os.Stat(absFn); err == nil {
					tpw.EndWithResult("already exists")
					continue
				}
			}

			if err := dc.Download(afd.Url(id), tpw, absDownloadsDir, id, relFn); err != nil {
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

func authorsFullNames(id string, rdx kvas.ReadableRedux) ([]string, error) {

	if err := rdx.MustHave(
		data.PersonsIdsProperty,
		data.PersonsRolesProperty,
		data.PersonFullNameProperty,
	); err != nil {
		return nil, err
	}

	authorsNames := make([]string, 0)

	if pids, ok := rdx.GetAllValues(data.PersonsIdsProperty, id); ok && len(pids) > 0 {
		if prs, sure := rdx.GetAllValues(data.PersonsRolesProperty, id); sure && len(prs) == len(pids) {

			for i := 0; i < len(prs); i++ {
				if prs[i] != "author" {
					continue
				}
				if afn, fine := rdx.GetLastVal(data.PersonFullNameProperty, pids[i]); fine {
					authorsNames = append(authorsNames, afn)
				}
			}
		}
	}

	return authorsNames, nil
}
