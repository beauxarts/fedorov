package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathways"
	"github.com/boggydigital/redux"
	"net/http"
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

	return DownloadLitResBooks(nil, force, ids...)
}

func DownloadLitResBooks(hc *http.Client, force bool, artsIds ...string) error {

	da := nod.NewProgress("downloading LitRes books...")
	defer da.Done()

	rdx, err := data.NewReduxReader(
		data.ArtsOperationsOrderProperty,
		data.TitleProperty,
		data.PersonsIdsProperty,
		data.PersonsRolesProperty,
		data.PersonFullNameProperty,
	)
	if err != nil {
		return err
	}

	kv, err := data.NewArtsReader(litres_integration.ArtsTypeFiles)
	if err != nil {
		return err
	}

	if artsIds == nil {
		artsIds, err = GetRecentArts(force)
		if err != nil {
			return err
		}
	}

	da.TotalInt(len(artsIds))

	if hc == nil {
		hc, err = getHttpClient()
		if err != nil {
			return err
		}
	}

	dc := dolo.NewClient(hc, dolo.Defaults())

	absDownloadsDir, err := pathways.GetAbsDir(data.Downloads)
	if err != nil {
		return err
	}

	for _, id := range artsIds {

		title, _ := rdx.GetLastVal(data.TitleProperty, id)
		authorsNames, err := authorsFullNames(id, rdx)
		if err != nil {
			return err
		}

		bdla := nod.Begin("%s %s - %s", id, strings.Join(authorsNames, ","), title)

		if !kv.Has(id) {
			continue
		}

		artFiles, err := kv.ArtsFiles(id)
		if err != nil {
			return err
		}

		for _, afd := range artFiles.PreferredDownloadsTypes() {

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

			if err := dc.Download(afd.Url(id), force, tpw, absDownloadsDir, id, relFn); err != nil {
				nod.Log(err.Error())
				continue
			}

			tpw.Done()
		}

		bdla.Done()
		da.Increment()
	}

	return nil
}

func authorsFullNames(id string, rdx redux.Readable) ([]string, error) {

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
