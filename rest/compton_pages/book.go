package compton_pages

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/beauxarts/fedorov/rest/compton_styles"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/kevlar"
)

func Book(id string, hasSections []string, rdx kevlar.ReadableRedux) compton.PageElement {

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

	for ii, section := range hasSections {

		sectionTitle := compton_data.SectionTitles[section]
		summaryHeading := compton.DSTitle(p, sectionTitle)
		detailsSummary := compton.DSLarge(p, summaryHeading, ii == 0).
			BackgroundColor(color.Highlight).
			ForegroundColor(color.Foreground).
			MarkerColor(color.Gray).
			SummaryMarginBlockEnd(size.Normal).
			DetailsMarginBlockEnd(size.Unset)
		detailsSummary.SetId(sectionTitle)

		switch section {
		case compton_data.InformationSection:
			detailsSummary.Append(compton_fragments.BookProperties(p, id, rdx))
		case compton_data.ExternalLinksSection:
		default:
			ifh := compton.IframeExpandHost(p, section, "/"+section+"?id="+id)
			detailsSummary.Append(ifh)
		}

		pageStack.Append(detailsSummary)
	}

	pageStack.Append(compton.Br(),
		compton.Footer(p, "Tokyo", "https://github.com/beauxarts", "🇯🇵"))

	return p
}
