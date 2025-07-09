package compton_fragments

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_styles"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/issa"
	"github.com/boggydigital/redux"
)

func ProductSection(section string, id string, rdx redux.Readable) compton.PageElement {

	title := compton_data.SectionTitles[section]
	ifc := compton.IframeExpandContent(section, title)

	if style, ok := compton_data.SectionStyles[section]; ok && style != "" {
		ifc.RegisterStyles(compton_styles.Styles, style)
	}

	if repColor, ok := rdx.GetLastVal(data.RepItemImageColorProperty, id); ok && repColor != issa.NeutralRepColor {
		ifc.SetAttribute("style", "--c-rep:"+repColor)
	}

	return ifc
}
