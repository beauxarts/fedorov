package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/beauxarts/fedorov/rest/compton_styles"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/input_types"
	"github.com/boggydigital/compton/consts/size"
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

	p := compton.Page("Latest")
	p.RegisterStyles(compton_styles.Styles, "book-labels.css")

	stack := compton.FlexItems(p, direction.Column)
	p.Append(stack)

	appNav := compton_fragments.AppNavLinks(p, compton_data.AppNavLatest)

	showAllLink := compton.A("/latest?all")
	showAllLink.Append(compton.InputValue(p, input_types.Button, "Show All"))

	topNav := compton.FICenter(p, appNav)
	if !all {
		topNav.Append(showAllLink)
	}

	stack.Append(topNav)

	lpTitle := compton.DSTitle(p, "Последние приобретения")
	latestPurchases := compton.DSLarge(p, lpTitle, true).
		BackgroundColor(color.Highlight).
		SummaryMarginBlockEnd(size.Normal).
		DetailsMarginBlockEnd(size.Unset).
		SummaryRowGap(size.XXSmall)

	//itemsCount := compton_fragments.ItemsCount(p, 0, len(ids), updateTotals[section])
	//sectionDetailsToggle.AppendSummary(itemsCount)

	stack.Append(latestPurchases)

	gridItems := compton.GridItems(p).JustifyContent(align.Center)
	latestPurchases.Append(gridItems)

	if ahop, ok := rdx.GetAllValues(data.ArtsHistoryOrderProperty, data.ArtsHistoryOrderProperty); ok {

		if !all {
			ahop = ahop[:latestBooksLimit]
		}

		for _, id := range ahop {
			bookLink := compton.A("/new_book?id=" + id)
			bookCard := compton_fragments.BookCard(p, id, false, rdx)
			bookLink.Append(bookCard)
			gridItems.Append(bookLink)
		}
	}

	if !all {
		stack.Append(compton.FICenter(p, showAllLink))
	}

	stack.Append(compton.Br(),
		compton.Footer(p, "Tokyo", "https://github.com/beauxarts", "🇯🇵"))

	if err := p.WriteResponse(w); err != nil {
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
