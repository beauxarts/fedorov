package compton_pages

import (
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/beauxarts/fedorov/rest/compton_styles"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/kevlar"
)

func Similar(id string, artsSimilar *litres_integration.ArtsSimilar, rdx kevlar.ReadableRedux) compton.PageElement {
	s := compton_fragments.ProductSection(compton_data.SimilarSection)
	s.RegisterStyles(compton_styles.Styles, "book-labels.css")

	if info := compton_fragments.Similar(s, id, artsSimilar, rdx); info != nil {
		s.Append(info)
	}
	return s
}
