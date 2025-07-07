package compton_fragments

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/boggydigital/redux"
)

var artTypeSection = map[litres_integration.ArtsType]string{
	litres_integration.ArtsTypeSimilar: compton_data.SimilarSection,
	litres_integration.ArtsTypeReviews: compton_data.ReviewsSection,
	litres_integration.ArtsTypeFiles:   compton_data.FilesSection,
}

func BookSections(id string, rdx redux.Readable) []string {
	hasSections := make([]string, 0, 3)
	// every book is expected to have at least those sections
	hasSections = append(hasSections, compton_data.InformationSection, compton_data.AnnotationSection, compton_data.ExternalLinksSection)

	artsTypes := []litres_integration.ArtsType{
		litres_integration.ArtsTypeSimilar,
		litres_integration.ArtsTypeReviews,
	}

	for _, at := range artsTypes {
		if hasArtsType(id, at) {
			hasSections = append(hasSections, artTypeSection[at])
		}
	}

	if videos, ok := rdx.GetAllValues(data.YouTubeVideosProperty, id); ok && len(videos) > 0 {
		hasSections = append(hasSections, compton_data.VideosSection)
	}

	if contentsUrl, ok := rdx.GetLastVal(data.ContentsUrlProperty, id); ok && contentsUrl != "" {
		hasSections = append(hasSections, compton_data.ContentsSection)
	}

	if hasArtsType(id, litres_integration.ArtsTypeFiles) {
		hasSections = append(hasSections, artTypeSection[litres_integration.ArtsTypeFiles])
	}

	return hasSections
}

func hasArtsType(id string, at litres_integration.ArtsType) bool {
	if reader, err := data.NewArtsReader(at); err == nil {
		if reader.Has(id) {
			return true
		}
	}
	return false
}
