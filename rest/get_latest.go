package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_pages"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/redux"
	"net/http"
)

const latestBooksLimit = 60

type LatestBookViewModel struct {
	Id      string
	Title   string
	Authors []string
}

func GetLatest(w http.ResponseWriter, r *http.Request) {

	// GET /latest

	var err error
	if rdx, err = rdx.RefreshReader(); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	all := r.URL.Query().Has("all")

	var ids []string
	var total int

	if ahop, ok := rdx.GetAllValues(data.ArtsOperationsOrderProperty, data.ArtsOperationsOrderProperty); ok {
		total = len(ahop)
		if !all {
			ahop = ahop[:latestBooksLimit]
		}
		ids = ahop
	}

	if p := compton_pages.Latest(ids, total, rdx); p != nil {
		if err := p.WriteResponse(w); err != nil {
			http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
			return
		}
	}
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
