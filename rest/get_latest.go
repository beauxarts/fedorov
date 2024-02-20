package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"net/http"
)

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

	lbvm := make([]*LatestBookViewModel, 0, latestBooksLimit)

	if ahop, ok := rdx.GetAllValues(data.ArtsHistoryOrderProperty, data.ArtsHistoryOrderProperty); ok {

		if !all {
			ahop = ahop[:latestBooksLimit]
		}

		for _, id := range ahop {
			title := ""
			if t, ok := rdx.GetLastVal(data.TitleProperty, id); ok {
				title = t
			}

			var authors []string
			if aus, err := authorsFullNames(id, rdx); err == nil {
				authors = aus
			}

			lbvm = append(lbvm, &LatestBookViewModel{
				Id:      id,
				Title:   title,
				Authors: authors,
			})
		}
	}

	DefaultHeaders(w)

	if err := tmpl.ExecuteTemplate(w, "latest", lbvm); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
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
