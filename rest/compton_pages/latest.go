package compton_pages

import (
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
)

const dehydratedCount = 10

func Latest(ids []string, total int, rdx kevlar.ReadableRedux) compton.PageElement {
	p := compton.Page(compton_data.AppNavLatest)
	p.RegisterStyles(compton_styles.Styles, "book-labels.css")

	stack := compton.FlexItems(p, direction.Column)
	p.Append(stack)

	appNav := compton_fragments.AppNavLinks(p, compton_data.AppNavLatest)

	showAllLink := compton.A("/latest?all")
	showAllLink.Append(compton.InputValue(p, input_types.Button, "Показать все"))

	topNav := compton.FICenter(p, appNav)
	if len(ids) < total {
		topNav.Append(showAllLink)
	}

	stack.Append(topNav)

	title := "Последние приобретения"
	if len(ids) == total {
		title = "Все книги"
	}
	lpTitle := compton.DSTitle(p, title)

	latestPurchases := compton.DSLarge(p, lpTitle, true).
		BackgroundColor(color.Highlight).
		SummaryMarginBlockEnd(size.Normal).
		DetailsMarginBlockEnd(size.Unset).
		SummaryRowGap(size.XXSmall)

	itemsCount := compton_fragments.ItemsCount(p, 0, len(ids), total)
	latestPurchases.AppendSummary(itemsCount)

	stack.Append(latestPurchases)

	gridItems := compton.GridItems(p).JustifyContent(align.Center)
	latestPurchases.Append(gridItems)

	for ii, id := range ids {
		bookLink := compton.A("/book?id=" + id)
		bookCard := compton_fragments.BookCard(p, id, ii < dehydratedCount, rdx)
		bookLink.Append(bookCard)
		gridItems.Append(bookLink)
	}

	if len(ids) < total {
		stack.Append(compton.FICenter(p, showAllLink))
	}

	stack.Append(compton.Br(),
		compton.Footer(p, "Tokyo", "https://github.com/beauxarts", "🇯🇵"))

	return p
}
