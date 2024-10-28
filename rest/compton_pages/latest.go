package compton_pages

import (
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/beauxarts/fedorov/rest/compton_styles"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/input_types"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/kevlar"
)

const dehydratedCount = 10

func Latest(ids []string, total int, rdx kevlar.ReadableRedux) compton.PageElement {
	p := compton.Page(compton_data.AppNavLatest)
	p.RegisterStyles(compton_styles.Styles, "book-labels.css")

	pageStack := compton.FlexItems(p, direction.Column)
	p.Append(pageStack)

	appNav := compton_fragments.AppNavLinks(p, compton_data.AppNavLatest)

	showAllLink := compton.A("/latest?all")
	showAllLink.Append(compton.InputValue(p, input_types.Button, "Показать все"))

	topNav := compton.FICenter(p, appNav)
	if len(ids) < total {
		topNav.Append(showAllLink)
	}

	pageStack.Append(topNav)

	title := "Новинки"
	if len(ids) == total {
		title = "Все книги"
	}
	lpTitle := compton.DSTitle(p, title)

	latestPurchases := compton.DSLarge(p, lpTitle, true).
		BackgroundColor(color.Highlight).
		SummaryMarginBlockEnd(size.Normal).
		DetailsMarginBlockEnd(size.Unset).
		SummaryRowGap(size.XXSmall)

	cf := compton.NewCountFormatter(
		compton_data.SingleItem,
		compton_data.ManyItemsSinglePage,
		compton_data.ManyItemsManyPages)

	latestPurchases.AppendSummary(cf.TitleElement(p, 0, len(ids), total))

	pageStack.Append(latestPurchases)

	booksList := compton_fragments.BooksList(p, ids, 0, len(ids), rdx)
	latestPurchases.Append(booksList)

	if len(ids) < total {
		pageStack.Append(compton.FICenter(p, showAllLink))
	}

	pageStack.Append(compton.Br(),
		compton.Footer(p, "Tokyo", "https://github.com/beauxarts", "🇯🇵"))

	return p
}
