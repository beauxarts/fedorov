package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/coost"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func SyncHandler(u *url.URL) error {
	force := u.Query().Has("force")

	return Sync(force)
}
func Sync(force bool) error {

	hc, err := getHttpClient()
	if err != nil {
		return err
	}

	sessionId, err := GetSessionId(hc)

	if err := HasArts(sessionId, hc); err != nil {
		return err
	}

	if err := GetLitResHistoryLog(sessionId, hc); err != nil {
		return err
	}

	if err := ReduceLitResHistoryLog(); err != nil {
		return err
	}

	if err := GetLitResArts(litres_integration.AllArtsTypes(), sessionId, hc, force); err != nil {
		return err
	}

	if err := ReduceLitResArtsDetails(); err != nil {
		return err
	}

	if err := GetLitResAuthors(litres_integration.AllAuthorTypes(), sessionId, hc, force); err != nil {
		return err
	}

	if err := GetLitResSeries(litres_integration.AllSeriesTypes(), sessionId, hc, force); err != nil {
		return err
	}

	if err := GetLitresContents(sessionId, hc, force); err != nil {
		return err
	}

	if err := Cascade(); err != nil {
		return err
	}

	if err := DownloadLitResBooks(sessionId, hc, false); err != nil {
		return err
	}

	if err := DownloadLitResCovers(true, false); err != nil {
		return err
	}

	if err := Dehydrate(map[string]bool{}, true, false); err != nil {
		return err
	}

	if err := Backup(); err != nil {
		return err
	}

	rdx, err := data.NewReduxWriter(data.SyncCompletedProperty)
	if err != nil {
		return err
	}

	return rdx.ReplaceValues(
		data.SyncCompletedProperty,
		data.SyncCompletedProperty,
		strconv.FormatInt(time.Now().UTC().Unix(), 10))
}

func getHttpClient() (*http.Client, error) {
	absCookiesFilename, err := data.AbsCookiesFilename()
	if err != nil {
		return nil, err
	}

	hc, err := coost.NewHttpClientFromFile(absCookiesFilename)
	if err != nil {
		return nil, err
	}

	return hc, err
}
