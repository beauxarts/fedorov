package compton_pages

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/redux"
)

func Reviews(id string, rdx redux.Readable) compton.PageElement {

	s := compton_fragments.ProductSection(compton_data.ReviewsSection, id, rdx)

	raReader, err := data.NewArtsReader(litres_integration.ArtsTypeReviews)
	if err != nil {
		s.Error(err)
		return s
	}

	artsReviews, err := raReader.ArtsReviews(id)
	if err != nil {
		s.Error(err)
		return s
	}

	if info := compton_fragments.Reviews(s, artsReviews); info != nil {
		s.Append(info)
	}
	return s
}
