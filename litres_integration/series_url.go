package litres_integration

import (
	"net/url"
	"strings"
)

var seriesPathTemplates = map[SeriesType]string{
	SeriesDetails: seriesDetailsPathTemplate,
	SeriesArts:    seriesArtsPathTemplate,
}

func SeriesUrl(st SeriesType, id string) *url.URL {
	pathTemplate := seriesPathTemplates[st]
	if pathTemplate == "" {
		return nil
	}

	seriesPath := strings.Replace(pathTemplate, "{id}", id, -1)

	return &url.URL{
		Scheme: httpsScheme,
		Host:   apiLitResHost,
		Path:   seriesPath,
	}
}
