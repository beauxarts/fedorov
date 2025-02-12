package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/coost"
	"github.com/boggydigital/nod"
	"net/http"
	"net/url"
)

func GetSessionIdHandler(u *url.URL) error {
	_, err := GetSessionId(nil)
	return err
}
func GetSessionId(hc *http.Client) (string, error) {
	gsia := nod.Begin("getting session-id...")
	defer gsia.Done()

	if hc == nil {
		absCookiesFilename, err := data.AbsCookiesFilename()
		if err != nil {
			return "", err
		}

		hc, err = coost.NewHttpClientFromFile(absCookiesFilename)
		if err != nil {
			return "", err
		}
	}

	sessionId, err := litres_integration.GetSessionId(hc)
	if err != nil {
		return "", err
	}

	gsia.EndWithResult(sessionId)
	return sessionId, nil
}
