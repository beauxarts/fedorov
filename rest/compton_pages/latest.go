package compton_pages

import (
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/redux"
)

func Latest(ids []string, total int, rdx redux.Readable) compton.PageElement {

	p, pageStack := compton_fragments.AppPage(compton_data.AppNavLatest)
	p.AppendSpeculationRules("/book?id=*")

	p.SetAttribute("style", "--c-rep:var(--c-background)")

	appNav := compton_fragments.AppNavLinks(p, compton_data.AppNavLatest)

	showAllNavLinks := compton.NavLinks(p)
	showAllNavLinks.AppendLink(p, &compton.NavTarget{
		Href:  "/latest?all",
		Title: "Показать все",
	})
	showAllNavLinks.SetAttribute("style", "view-transition-name:secondary-nav")

	topNav := compton.FICenter(p, appNav)
	if len(ids) < total {
		topNav.Append(showAllNavLinks)
	}

	pageStack.Append(topNav)

	title := "Новинки"
	if len(ids) == total {
		title = "Все книги"
	}

	latestPurchases := compton.DSLarge(p, title, true).
		BackgroundColor(color.Highlight).
		SummaryMarginBlockEnd(size.Normal).
		DetailsMarginBlockEnd(size.Unset).
		SummaryRowGap(size.XXSmall)

	cf := compton.NewCountFormatter(
		compton_data.SingleItem,
		compton_data.ManyItemsSinglePage,
		compton_data.ManyItemsManyPages)

	latestBadge := compton.Badge(p, cf.Title(0, len(ids), total), color.Background, color.Foreground)
	latestPurchases.AppendBadges(latestBadge)

	pageStack.Append(latestPurchases)

	booksList := compton_fragments.BooksList(p, ids, 0, len(ids), rdx)
	latestPurchases.Append(booksList)

	pageStack.Append(compton.Br(),
		compton.Footer(p, "Tokyo", "https://github.com/beauxarts", "🇯🇵"))

	return p
}
