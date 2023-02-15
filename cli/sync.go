package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/coost"
	"github.com/boggydigital/kvas"
	"net/url"
	"strconv"
	"time"
)

func SyncHandler(u *url.URL) error {
	newOnly := u.Query().Has("new-only")

	return Sync(newOnly)
}
func Sync(newOnly bool) error {

	hc, err := coost.NewHttpClientFromFile(data.AbsCookiesFilename(), litres_integration.LitResHost)
	if err != nil {
		return err
	}

	if err := GetLitResMyBooks(hc); err != nil {
		return err
	}

	if err := ReduceLitResMyBooks(); err != nil {
		return err
	}

	if err := GetLitResDetails(nil, hc, newOnly); err != nil {
		return err
	}

	if err := ReduceLitResBooksDetails(true); err != nil {
		return err
	}

	if err := Cascade(); err != nil {
		return err
	}

	if err := DownloadLitRes(nil, hc); err != nil {
		return err
	}

	if err := GetLitResCovers(nil, false); err != nil {
		return err
	}

	if err := Backup(); err != nil {
		return err
	}

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), nil, data.SyncCompletedProperty)
	if err != nil {
		return err
	}

	tnu := time.Now().UTC().Unix()

	return rxa.ReplaceValues(data.SyncCompletedProperty, data.SyncCompletedProperty, strconv.FormatInt(tnu, 10))
}
