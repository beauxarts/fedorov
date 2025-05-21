package litres_integration

import "net/url"

func ContentsUrl(path string) *url.URL {
	return &url.URL{
		Scheme: httpsScheme,
		Host:   LitResHost,
		Path:   path,
	}
}
