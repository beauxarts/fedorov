package litres_integration

import (
	"net/url"
	"strings"
)

var artsPathTemplates = map[ArtsType]string{
	ArtsTypeDetails: artsDetailsPathTemplate,
	ArtsTypeSimilar: artsSimilarPathTemplate,
	ArtsTypeQuotes:  artsQuotesPathTemplate,
	ArtsTypeFiles:   artsFilesPathTemplate,
	ArtsTypeReviews: artsReviewsPathTemplate,
}

func ArtsTypeUrl(at ArtsType, id string) *url.URL {
	pathTemplate := artsPathTemplates[at]
	if pathTemplate == "" {
		return nil
	}

	artsPath := strings.Replace(pathTemplate, "{id}", id, -1)

	return &url.URL{
		Scheme: httpsScheme,
		Host:   apiLitResHost,
		Path:   artsPath,
	}
}
