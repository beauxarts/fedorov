package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kvas"
	"net/url"
	"strconv"
	"time"
)

func SyncHandler(u *url.URL) error {
	newOnly := u.Query().Has("new-only")
	noThrottle := u.Query().Has("no-throttle")
	wu := u.Query().Get("webhook-url")

	return Sync(wu, newOnly, noThrottle)
}
func Sync(webhookUrl string, newOnly, noThrottle bool) error {

	if err := GetLitResMyBooks(); err != nil {
		return err
	}

	if err := ReduceLitResMyBooks(); err != nil {
		return err
	}

	if err := GetLitResDetails(nil, newOnly, noThrottle); err != nil {
		return err
	}

	if err := ReduceLitResBooksDetails(true); err != nil {
		return err
	}

	if err := Cascade(); err != nil {
		return err
	}

	if err := DownloadLitRes(nil); err != nil {
		return err
	}

	if err := GetLitResCovers(nil, false); err != nil {
		return err
	}

	if err := Dehydrate(map[string]bool{}, true, false); err != nil {
		return err
	}

	if err := Backup(); err != nil {
		return err
	}

	if err := PostCompletion(webhookUrl); err != nil {
		return err
	}

	rdx, err := kvas.ReduxWriter(data.AbsReduxDir(), data.SyncCompletedProperty)
	if err != nil {
		return err
	}

	tnu := time.Now().UTC().Unix()

	return rdx.ReplaceValues(data.SyncCompletedProperty, data.SyncCompletedProperty, strconv.FormatInt(tnu, 10))
}
