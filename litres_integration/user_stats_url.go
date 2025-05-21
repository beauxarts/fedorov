package litres_integration

import "net/url"

func UserStatsUrl() *url.URL {
	return &url.URL{
		Scheme: httpsScheme,
		Host:   apiLitResHost,
		Path:   usersMeArtsStatsPath,
	}
}
