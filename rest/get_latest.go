package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/beauxarts/fedorov/rest/compton_styles"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/input_types"
	"github.com/boggydigital/kevlar"
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

	page := compton.Page("Latest")
	page.RegisterStyles(compton_styles.Styles, "book-labels.css")

	stack := compton.FlexItems(page, direction.Column)
	page.Append(stack)

	gridItems := compton.GridItems(page).JustifyContent(align.Center)
	stack.Append(gridItems)

	if ahop, ok := rdx.GetAllValues(data.ArtsHistoryOrderProperty, data.ArtsHistoryOrderProperty); ok {

		if !all {
			ahop = ahop[:latestBooksLimit]
		}

		for _, id := range ahop {
			bookLink := compton.A("/new_book?id=" + id)
			bookCard := compton_fragments.BookCard(page, id, false, rdx)
			bookLink.Append(bookCard)
			gridItems.Append(bookLink)
		}
	}

	showAllLink := compton.A("/latest?all")
	showAllLink.Append(compton.InputValue(page, input_types.Button, "Show All"))
	stack.Append(compton.FICenter(page, showAllLink))

	if err := page.WriteResponse(w); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}

func authorsFullNames(id string, rdx kevlar.ReadableRedux) ([]string, error) {

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
