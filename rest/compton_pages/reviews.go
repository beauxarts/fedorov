package compton_pages

import (
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/compton"
)

func Reviews(id string, artsReviews *litres_integration.ArtsReviews) compton.PageElement {
	s := compton_fragments.ProductSection(compton_data.ReviewsSection)
	if info := compton_fragments.Reviews(s, id, artsReviews); info != nil {
		s.Append(info)
	}
	return s
}
