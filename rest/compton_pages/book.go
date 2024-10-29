package compton_pages

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/beauxarts/fedorov/rest/compton_styles"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/kevlar"
)

func Book(id string, rdx kevlar.ReadableRedux) compton.PageElement {

	title, _ := rdx.GetLastVal(data.TitleProperty, id)

	p, pageStack := compton_fragments.AppPage(title)
	p.RegisterStyles(compton_styles.Styles, "app.css")

	appNav := compton_fragments.AppNavLinks(p, "")
	topNav := compton.FICenter(p, appNav)
	pageStack.Append(topNav)

	if cover := compton_fragments.BookCover(p, id, rdx); cover != nil {
		pageStack.Append(cover)
	}

	productTitle := compton.HeadingText(title, 1)
	productTitle.AddClass("product-title")

	fmtLabels := compton_fragments.FormatLabels(id, rdx)
	productLabels := compton.Labels(p, fmtLabels...).FontSize(size.Small).RowGap(size.XSmall).ColumnGap(size.XSmall)
	pageStack.Append(compton.FICenter(p, productTitle, productLabels))

	return p
}
