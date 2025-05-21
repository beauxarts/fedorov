package litres_integration

import "net/url"

func HrefUrl(href string) *url.URL {
	return &url.URL{
		Scheme: httpsScheme,
		Host:   wwwLitResHost,
		Path:   href,
	}
}
