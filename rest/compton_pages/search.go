package compton_pages

import (
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/kevlar"
	"strconv"
)

const filterSearchTitle = "Фильтр и поиск"

func Search(query map[string][]string, ids []string, from, to int, rdx kevlar.ReadableRedux) compton.PageElement {

	p, pageStack := compton_fragments.AppPage(compton_data.AppNavSearch)

	appNav := compton_fragments.AppNavLinks(p, compton_data.AppNavSearch)

	searchScope := compton_data.SearchScopeFromQuery(query)
	searchLinks := compton_fragments.SearchLinks(p, searchScope)

	pageStack.Append(compton.FICenter(p, appNav, searchLinks))

	filterSearchHeading := compton.DSTitle(p, filterSearchTitle)

	filterSearchDetails := compton.DSLarge(p, filterSearchHeading, len(query) == 0).
		BackgroundColor(color.Highlight).
		SummaryMarginBlockEnd(size.Normal).
		DetailsMarginBlockEnd(size.Unset).
		SummaryRowGap(size.XXSmall)

	if len(query) > 0 {

		cf := compton.NewCountFormatter(
			compton_data.SingleItem,
			compton_data.ManyItemsSinglePage,
			compton_data.ManyItemsManyPages)

		filterSearchDetails.AppendSummary(cf.TitleElement(p, from, to, len(ids)))
	}

	searchQuery := compton.Query(p, query,
		compton_data.PropertyTitles, "/search", "Очистить")

	filterSearchDetails.Append(compton_fragments.SearchForm(p, query, searchQuery, rdx))
	pageStack.Append(filterSearchDetails)

	if searchQuery != nil {
		pageStack.Append(searchQuery)
	}

	if len(ids) > 0 {
		booksList := compton_fragments.BooksList(p, ids, from, to, rdx)
		pageStack.Append(booksList)
	}

	if to < len(ids) {
		query["from"] = []string{strconv.Itoa(to)}
		enq := compton_data.EncodeQuery(query)

		href := "/search?" + enq

		pageStack.Append(compton_fragments.Button(p, "След. страница", href))
	}

	pageStack.Append(compton.Br(),
		compton.Footer(p, "Tokyo", "https://github.com/beauxarts", "🇯🇵"))

	return p
}
