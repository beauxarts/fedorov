package compton_pages

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/redux"
)

func Similar(id string, rdx redux.Readable) compton.PageElement {
	s := compton_fragments.ProductSection(compton_data.SimilarSection, id, rdx)

	saReader, err := data.NewArtsReader(litres_integration.ArtsTypeSimilar)
	if err != nil {
		s.Error(err)
		return s
	}

	var similarArts *litres_integration.ArtsSimilar

	if sa, err := saReader.ArtsSimilar(id); err == nil {
		similarArts = sa
	} else {
		s.Error(err)
		return s
	}

	if info := compton_fragments.Similar(s, id, similarArts, rdx); info != nil {
		s.Append(info)
	}
	return s
}
