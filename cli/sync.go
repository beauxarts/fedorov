package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"net/url"
	"strconv"
	"time"
)

func SyncHandler(u *url.URL) error {
	force := u.Query().Has("force")

	return Sync(force)
}
func Sync(force bool) error {

	if err := GetLitResHistoryLog(); err != nil {
		return err
	}

	if err := ReduceLitResHistoryLog(); err != nil {
		return err
	}

	if err := GetLitResArts(litres_integration.AllArtsTypes(), force); err != nil {
		return err
	}

	if err := ReduceLitResArtsDetails(); err != nil {
		return err
	}

	if err := GetLitResAuthors(litres_integration.AllAuthorTypes(), force); err != nil {
		return err
	}

	if err := GetLitResSeries(litres_integration.AllSeriesTypes(), force); err != nil {
		return err
	}

	if err := GetLitresContents(force); err != nil {
		return err
	}

	if err := Cascade(); err != nil {
		return err
	}

	if err := DownloadLitResBooks(false); err != nil {
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
