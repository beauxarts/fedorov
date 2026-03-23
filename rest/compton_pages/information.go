package compton_pages

import (
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/redux"
)

func Information(id string, rdx redux.Readable) compton.PageElement {
	s := compton_fragments.ProductSection(compton_data.InformationSection)
	if info := compton_fragments.Information(s, id, rdx); info != nil {
		s.Append(info)
	}
	return s
}
