package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/nod"
	"net/http"
	"net/url"
)

func HasArtsHandler(u *url.URL) error {

	sessionId := u.Query().Get("session-id")

	return HasArts(sessionId, nil)
}
func HasArts(sessionId string, hc *http.Client) error {

	haa := nod.Begin("checking if this user has any arts...")
	defer haa.End()

	if hc == nil {
		var err error
		hc, err = getHttpClient()
		if err != nil {
			return haa.EndWithError(err)
		}
	}

	req, err := http.NewRequest(http.MethodGet,
		litres_integration.UserStatsUrl().String(), nil)
	if err != nil {
		return haa.EndWithError(err)
	}

	addHeaders(req, sessionId)

	resp, err := hc.Do(req)
	if err != nil {
		return haa.EndWithError(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return haa.EndWithError(errors.New(resp.Status))
	}

	var userStats litres_integration.UserStats
	if err := json.NewDecoder(resp.Body).Decode(&userStats); err != nil {
		return haa.EndWithError(err)
	}

	switch received := userStats.Received(); received {
	case 0:
		return haa.EndWithError(errors.New("no arts found. If this is not expected - please update cookie.txt"))
	default:
		haa.EndWithResult("found %d art(s)", received)
	}

	fmt.Println(userStats.Received())

	return nil
}
