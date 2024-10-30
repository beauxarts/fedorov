package compton_fragments

import (
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_styles"
	"github.com/boggydigital/compton"
)

func ProductSection(section string) compton.PageElement {

	title := compton_data.SectionTitles[section]
	ifc := compton.IframeExpandContent(section, title)

	if style, ok := compton_data.SectionStyles[section]; ok && style != "" {
		ifc.RegisterStyles(compton_styles.Styles, style)
	}

	return ifc
}
