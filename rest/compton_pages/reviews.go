package compton_pages

import (
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/boggydigital/compton"
)

func Reviews(artsReviews *litres_integration.ArtsReviews) compton.PageElement {
	s := compton_fragments.ProductSection(compton_data.ReviewsSection)
	if info := compton_fragments.Reviews(s, artsReviews); info != nil {
		s.Append(info)
	}
	return s
}
