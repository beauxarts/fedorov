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
	"github.com/boggydigital/issa"
	"github.com/boggydigital/kevlar"
	"strings"
)

const colorBlendClass = "color-blend"

func Book(id string, hasSections []string, rdx kevlar.ReadableRedux) compton.PageElement {

	title, _ := rdx.GetLastVal(data.TitleProperty, id)

	p, pageStack := compton_fragments.AppPage(title)
	p.RegisterStyles(compton_styles.Styles, "book.css")

	// tinting document background color to the representative product color
	if repColor, ok := rdx.GetLastVal(data.RepItemImageColorProperty, id); ok && repColor != issa.NeutralRepColor {
		compton.SetTint(p, repColor)
	}

	appNav := compton_fragments.AppNavLinks(p, "")
	appNav.AddClass(colorBlendClass)
	showToc := compton.InputValue(p, input_types.Button, "Разделы")
	showToc.AddClass(colorBlendClass)

	pageStack.Append(compton.FICenter(p, appNav, showToc))

	productSectionsLinks := compton_fragments.BookSectionsLinks(p, hasSections)
	pageStack.Append(productSectionsLinks)

	pageStack.Append(compton.Attach(p, showToc, productSectionsLinks))

	if cover := compton_fragments.BookCover(p, id, rdx); cover != nil {
		pageStack.Append(compton.FICenter(p, cover))
	}

	productTitle := compton.Heading(1)
	productTitle.AddClass(colorBlendClass)
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

	summaryRow := compton.Frow(p).
		FontSize(size.Small)
	properties, values := compton_fragments.SummarizeBookProperties(id, rdx)
	for _, p := range properties {
		summaryRow.PropVal(compton_data.PropertyTitles[p], strings.Join(values[p], ", "))
	}
	pageStack.Append(compton.FICenter(p, summaryRow))

	for _, section := range hasSections {

		sectionTitle := compton_data.SectionTitles[section]
		summaryHeading := compton.DSTitle(p, sectionTitle)
		detailsSummary := compton.DSLarge(p, summaryHeading, false).
			BackgroundColor(color.Highlight).
			ForegroundColor(color.Foreground).
			MarkerColor(color.Gray).
			SummaryMarginBlockEnd(size.Normal).
			DetailsMarginBlockEnd(size.Unset)
		detailsSummary.AddClassSummary(colorBlendClass)
		detailsSummary.SetId(sectionTitle)

		ifh := compton.IframeExpandHost(p, section, "/"+section+"?id="+id)
		detailsSummary.Append(ifh)

		pageStack.Append(detailsSummary)
	}

	pageStack.Append(compton.Br(),
		compton.Footer(p, "Tokyo", "https://github.com/beauxarts", "🇯🇵"))

	return p
}
