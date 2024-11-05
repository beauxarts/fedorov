package compton_pages

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/beauxarts/fedorov/rest/compton_styles"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/input_types"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/kevlar"
)

func Book(id string, hasSections []string, rdx kevlar.ReadableRedux) compton.PageElement {

	title, _ := rdx.GetLastVal(data.TitleProperty, id)

	p, pageStack := compton_fragments.AppPage(title)
	p.RegisterStyles(compton_styles.Styles, "book.css")

	// tinting document background color to the representative product color
	alpha := "5%"
	if repColor, ok := rdx.GetLastVal(data.RepItemImageColorProperty, id); ok {
		p.SetAttribute("style", "background-color:color-mix(in display-p3,"+repColor+" "+alpha+",var(--c-background))")
	}

	appNav := compton_fragments.AppNavLinks(p, "")
	showToc := compton.InputValue(p, input_types.Button, "Разделы")

	pageStack.Append(compton.FICenter(p, appNav, showToc))

	productSectionsLinks := compton_fragments.BookSectionsLinks(p, hasSections)
	pageStack.Append(productSectionsLinks)

	pageStack.Append(compton.Attach(p, showToc, productSectionsLinks))

	if cover := compton_fragments.BookCover(p, id, rdx); cover != nil {
		pageStack.Append(compton.FICenter(p, cover))
	}

	productTitle := compton.Heading(1)
	productTitle.Append(compton.Fspan(p, title).TextAlign(align.Center))

	fmtLabels := compton_fragments.FormatLabels(id, rdx)
	productLabels := compton.Labels(p, fmtLabels...).FontSize(size.Small).RowGap(size.XSmall).ColumnGap(size.XSmall)
	pageStack.Append(compton.FICenter(p, productTitle, productLabels))

	if subtitle, ok := rdx.GetLastVal(data.SubtitleProperty, id); ok {
		productSubtitle := compton.Fspan(p, subtitle).
			ForegroundColor(color.Gray).
			FontSize(size.Small).
			TextAlign(align.Center)
		pageStack.Append(productSubtitle)
	}

	for _, section := range hasSections {

		sectionTitle := compton_data.SectionTitles[section]
		summaryHeading := compton.DSTitle(p, sectionTitle)
		detailsSummary := compton.DSLarge(p, summaryHeading, false).
			BackgroundColor(color.Highlight).
			ForegroundColor(color.Foreground).
			MarkerColor(color.Gray).
			SummaryMarginBlockEnd(size.Normal).
			DetailsMarginBlockEnd(size.Unset)
		detailsSummary.SetId(sectionTitle)

		ifh := compton.IframeExpandHost(p, section, "/"+section+"?id="+id)
		detailsSummary.Append(ifh)

		pageStack.Append(detailsSummary)
	}

	pageStack.Append(compton.Br(),
		compton.Footer(p, "Tokyo", "https://github.com/beauxarts", "🇯🇵"))

	return p
}
