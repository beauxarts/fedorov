package compton_pages

import (
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/redux"
	"maps"
	"slices"
	"strconv"
)

const filterSearchTitle = "Фильтр и поиск"

func Search(query map[string][]string, ids []string, from, to int, rdx redux.Readable) compton.PageElement {

	p, pageStack := compton_fragments.AppPage(compton_data.AppNavSearch)
	p.AppendSpeculationRules("/book?id=*")

	p.SetAttribute("style", "--c-rep:var(--c-background)")

	searchScope := compton_data.SearchScopeFromQuery(query)
	if searchScope == "" {
		searchScope = "Поиск"
	}

	menuNav := compton_fragments.MenuNav(p, searchScope, "", rdx)
	pageStack.Append(menuNav)

	filterSearchDetails := compton.DSLarge(p, filterSearchTitle, len(query) == 0).
		BackgroundColor(color.Highlight).
		SummaryMarginBlockEnd(size.Normal).
		DetailsMarginBlockEnd(size.Unset).
		SummaryRowGap(size.XXSmall)

	if len(query) > 0 {

		cf := compton.NewCountFormatter(
			compton_data.SingleItem,
			compton_data.ManyItemsSinglePage,
			compton_data.ManyItemsManyPages)

		resultsBadge := compton.Badge(p, cf.Title(from, to, len(ids)), color.Background, color.Foreground)
		filterSearchDetails.AppendBadges(resultsBadge)
	}

	var queryFrow *compton.FrowElement
	if len(query) > 0 {
		queryFrow = compton.Frow(p).FontSize(size.Small)
		fq := compton_fragments.FormatQuery(query)
		props := maps.Keys(query)
		sortedProps := slices.Sorted(props)
		for _, prop := range sortedProps {
			vals := fq[prop]
			queryFrow.PropVal(compton_data.PropertyTitles[prop], vals...)
		}
		queryFrow.LinkColor("Очистить", "/search", color.Blue)
	}

	filterSearchDetails.Append(compton_fragments.SearchForm(p, query, queryFrow, rdx))
	pageStack.Append(filterSearchDetails)

	if queryFrow != nil {
		pageStack.Append(compton.FICenter(p, queryFrow))
	}

	if len(ids) > 0 {
		booksList := compton_fragments.BooksList(p, ids, from, to, rdx)
		pageStack.Append(booksList)
	}

	if to < len(ids) {
		query["from"] = []string{strconv.Itoa(to)}
		enq := compton_data.EncodeQuery(query)

		nextPageNavLink := compton.NavLinks(p)
		nextPageNavLink.AppendSubmitLink(p, &compton.NavTarget{
			Href:  "/search?" + enq,
			Title: "След. страница",
		})

		pageStack.Append(nextPageNavLink)
	}

	pageStack.Append(compton.Br(),
		compton.Footer(p, "Tokyo", "https://github.com/beauxarts", "🇯🇵"))

	return p
}
