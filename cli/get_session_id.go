package cli

import (
	"errors"
	"net/url"
	"os"

	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/boggydigital/coost"
	"github.com/boggydigital/nod"
)

func GetSessionIdHandler(_ *url.URL) error {
	_, err := GetSessionId()
	return err
}
func GetSessionId() (string, error) {
	gsia := nod.Begin("getting session-id...")
	defer gsia.Done()

	absCookiesFilename, err := data.AbsCookiesFilename()

	if err != nil {
		return "", err
	}

	jar, err := coost.Read(litres_integration.DefaultUrl(), absCookiesFilename)
	if os.IsNotExist(err) {
		return "", errors.New("cookies file not found, use import-cookies command to add")
	} else if err != nil {
		return "", err
	}

	for _, cookie := range jar.Cookies(litres_integration.DefaultUrl()) {
		if cookie.Name == "SID" {
			gsia.EndWithResult(cookie.Value)
			return cookie.Value, nil
		}
	}

	return "", errors.New("sessionId is empty")
}
