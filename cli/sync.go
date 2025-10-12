package cli

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/boggydigital/coost"
	"github.com/boggydigital/pathways"
	"github.com/boggydigital/redux"
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

	if err = HasArts(sessionId, hc); err != nil {
		return err
	}

	if err = GetLitResOperations(sessionId, hc); err != nil {
		return err
	}

	if err = ReduceLitResOperations(); err != nil {
		return err
	}

	recentArtsIds, err := GetRecentArts(force)
	if err != nil {
		return err
	}

	freeArtsIds, err := getFreeArts()
	if err != nil {
		return err
	}

	for freeArtId := range freeArtsIds {
		recentArtsIds = append(recentArtsIds, freeArtId)
	}

	if err = GetLitResArts(litres_integration.AllArtsTypes(), hc, force, recentArtsIds...); err != nil {
		return err
	}

	if err = ReduceLitResArtsDetails(); err != nil {
		return err
	}

	recentPersonsIds, err := GetRecentPersons(force, recentArtsIds...)
	if err != nil {
		return err
	}

	if err = GetLitResAuthors(litres_integration.AllAuthorTypes(), hc, force, recentPersonsIds...); err != nil {
		return err
	}

	recentSeriesIds, err := GetRecentSeries(force, recentArtsIds...)
	if err != nil {
		return err
	}

	if err = GetLitResSeries(litres_integration.AllSeriesTypes(), hc, force, recentSeriesIds...); err != nil {
		return err
	}

	if err = GetLitresContents(hc, force, recentArtsIds...); err != nil {
		return err
	}

	if err = Cascade(); err != nil {
		return err
	}

	if err = DownloadLitResCovers(true, force, recentArtsIds...); err != nil {
		return err
	}

	if err = Dehydrate(force, recentArtsIds...); err != nil {
		return err
	}

	if err = DownloadLitResBooks(hc, force, recentArtsIds...); err != nil {
		return err
	}

	if err = GetVideosMetadata(force); err != nil {
		return err
	}

	if err = Backup(); err != nil {
		return err
	}

	reduxDir, err := pathways.GetAbsRelDir(data.Redux)
	if err != nil {
		return err
	}

	rdx, err := redux.NewWriter(reduxDir, data.SyncCompletedProperty)
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
