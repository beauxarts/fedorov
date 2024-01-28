package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/pasu"
	"net/url"
	"strconv"
	"time"
)

func SyncHandler(u *url.URL) error {
	force := u.Query().Has("force")
	wu := u.Query().Get("webhook-url")

	return Sync(wu, force)
}
func Sync(webhookUrl string, force bool) error {

	if err := GetLitResHistoryLog(); err != nil {
		return err
	}

	if err := ReduceLitResHistoryLog(); err != nil {
		return err
	}

	if err := GetLitResArts(litres_integration.AllArtsTypes(), force); err != nil {
		return err
	}

	// add reduce arts

	// add get authors
	// add get series
	// add reduce authors
	// add reduce series

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

	absReduxDir, err := pasu.GetAbsRelDir(data.Redux)
	if err != nil {
		return err
	}

	rdx, err := kvas.NewReduxWriter(absReduxDir, data.SyncCompletedProperty)
	if err != nil {
		return err
	}

	tnu := time.Now().UTC().Unix()

	return rdx.ReplaceValues(data.SyncCompletedProperty, data.SyncCompletedProperty, strconv.FormatInt(tnu, 10))
}
