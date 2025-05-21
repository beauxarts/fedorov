package compton_pages

import (
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/boggydigital/compton"
)

func Contents(contents *litres_integration.Contents) compton.PageElement {
	s := compton_fragments.ProductSection(compton_data.ContentsSection)
	if contents := compton_fragments.Contents(s, contents); contents != nil {
		s.Append(contents)
	}
	return s
}
