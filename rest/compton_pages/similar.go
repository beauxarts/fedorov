package compton_pages

import (
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/redux"
)

func Similar(id string, artsSimilar *litres_integration.ArtsSimilar, rdx redux.Readable) compton.PageElement {
	s := compton_fragments.ProductSection(compton_data.SimilarSection)

	if info := compton_fragments.Similar(s, id, artsSimilar, rdx); info != nil {
		s.Append(info)
	}
	return s
}
