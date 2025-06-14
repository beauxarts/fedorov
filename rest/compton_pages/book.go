package compton_pages

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/beauxarts/fedorov/rest/compton_styles"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/issa"
	"github.com/boggydigital/redux"
	"slices"
	"strings"
)

var openSections = []string{
	compton_data.InformationSection,
	compton_data.FilesSection,
}

func Book(id string, hasSections []string, rdx redux.Readable) compton.PageElement {

	title, _ := rdx.GetLastVal(data.TitleProperty, id)

	p, pageStack := compton_fragments.AppPage(title)
	p.RegisterStyles(compton_styles.Styles, "book.css")

	// tinting document background color to the representative product color
	if repColor, ok := rdx.GetLastVal(data.RepItemImageColorProperty, id); ok && repColor != issa.NeutralRepColor {
		p.SetAttribute("style", "--c-rep:"+repColor)
	}

	appNav := compton_fragments.AppNavLinks(p, "")

	showTocNavLinks := compton.NavLinks(p)
	showTocLink := showTocNavLinks.AppendLink(p, &compton.NavTarget{
		Href:   "#",
		Symbol: compton.DownwardArrow,
	})

	topLevelNav := []compton.Element{appNav, showTocNavLinks}

	if bookSectionsLinks := compton.SectionsLinks(p, hasSections, compton_data.SectionTitles); bookSectionsLinks != nil {
		pageStack.Append(compton.Attach(p, showTocLink, bookSectionsLinks))
		topLevelNav = append(topLevelNav, bookSectionsLinks)
	}

	pageStack.Append(compton.FICenter(p, topLevelNav...))

	if cover := compton_fragments.BookCover(p, id, rdx); cover != nil {
		pageStack.Append(compton.FICenter(p, cover))
	}

	productTitle := compton.Heading(2)
	productTitle.Append(compton.Fspan(p, title).TextAlign(align.Center))
	productTitle.SetAttribute("style", "view-transition-name:product-title-"+id)

	pageStack.Append(compton.FICenter(p, productTitle))

	if subtitle, ok := rdx.GetLastVal(data.SubtitleProperty, id); ok {
		productSubtitle := compton.Fspan(p, subtitle).
			ForegroundColor(color.Gray).
			FontSize(size.XSmall).
			TextAlign(align.Center)
		pageStack.Append(productSubtitle)
	}

	summaryRow := compton.Frow(p).
		FontSize(size.XSmall)
	properties, values := compton_fragments.SummarizeBookProperties(id, rdx)
	for _, p := range properties {
		summaryRow.PropVal(compton_data.PropertyTitles[p], strings.Join(values[p], ", "))
	}
	pageStack.Append(compton.FICenter(p, summaryRow))

	for ii, section := range hasSections {

		sectionTitle := compton_data.SectionTitles[section]
		detailsSummary := compton.DSLarge(p, sectionTitle, slices.Contains(openSections, section)).
			BackgroundColor(color.Highlight).
			MarkerColor(color.Gray).
			SummaryMarginBlockEnd(size.Normal).
			DetailsMarginBlockEnd(size.Unset)
		detailsSummary.SetId(sectionTitle)
		detailsSummary.SetTabIndex(ii + 1)

		switch section {
		case compton_data.InformationSection:
			productBadges := compton.FlexItems(p, direction.Row).ColumnGap(size.XSmall)
			for _, fmtBadge := range compton_fragments.FormatBadges(id, rdx) {
				badge := compton.Badge(p, fmtBadge.Title, fmtBadge.Background, color.Highlight)
				badge.AddClass(fmtBadge.Class)
				productBadges.Append(badge)
			}
			detailsSummary.AppendBadges(productBadges)
		case compton_data.ReviewsSection:
			if ratingAvg := compton_fragments.RatingAvg(id, rdx); ratingAvg != "" {
				ratingBadge := compton.Badge(p, ratingAvg, color.Background, color.Foreground)
				detailsSummary.AppendBadges(ratingBadge)
			}
		}

		ifh := compton.IframeExpandHost(p, section, "/"+section+"?id="+id)
		detailsSummary.Append(ifh)

		pageStack.Append(detailsSummary)
	}

	pageStack.Append(compton.Br(),
		compton.Footer(p, "Tokyo", "https://github.com/beauxarts", "🇯🇵"))

	return p
}
