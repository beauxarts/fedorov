package litres_integration

import (
	"net/url"
	"strings"
)

var authorPathTemplates = map[AuthorType]string{
	AuthorDetails: authorDetailsPathTemplate,
	AuthorArts:    authorArtsPathTemplate,
}

func AuthorUrl(at AuthorType, id string) *url.URL {
	pathTemplate := authorPathTemplates[at]
	if pathTemplate == "" {
		return nil
	}

	authorPath := strings.Replace(pathTemplate, "{id}", id, -1)

	return &url.URL{
		Scheme: httpsScheme,
		Host:   apiLitResHost,
		Path:   authorPath,
	}
}
