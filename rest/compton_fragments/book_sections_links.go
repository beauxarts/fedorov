package compton_fragments

import (
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/size"
)

func BookSectionsLinks(r compton.Registrar, sections []string) compton.Element {

	linksStack := compton.FlexItems(r, direction.Row).
		JustifyContent(align.Center).
		FontSize(size.Small).
		FontWeight(font_weight.Bolder).
		RowGap(size.Small)

	for _, s := range sections {
		title := compton_data.SectionTitles[s]
		link := compton.A("#" + title)
		linkText := compton.Fspan(r, title)
		link.Append(linkText)
		linksStack.Append(link)
	}

	wrapper := compton.FICenter(r, linksStack)
	wrapper.SetId("product-sections-links")

	return wrapper
}
