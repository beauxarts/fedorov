package cli

import (
	"net/http"
	"net/url"

	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/boggydigital/nod"
)

func GetSessionIdHandler(u *url.URL) error {
	_, err := GetSessionId(nil)
	return err
}
func GetSessionId(hc *http.Client) (string, error) {
	gsia := nod.Begin("getting session-id...")
	defer gsia.Done()

	if hc == nil {
		var err error
		if hc, err = getHttpClient(); err != nil {
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
