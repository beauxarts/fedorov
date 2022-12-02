package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/litres_integration"
	"github.com/boggydigital/coost"
	"github.com/boggydigital/kvas"
	"net/url"
	"strconv"
	"time"
)

func SyncHandler(u *url.URL) error {
	return Sync()
}
func Sync() error {

	hc, err := coost.NewHttpClientFromFile(data.AbsCookiesFilename(), litres_integration.LitResHost)
	if err != nil {
		return err
	}

	if err := GetMyBooks(hc); err != nil {
		return err
	}

	if err := ReduceMyBooks(); err != nil {
		return err
	}

	if err := GetDetails(nil, hc); err != nil {
		return err
	}

	if err := ReduceDetails(true); err != nil {
		return err
	}

	if err := Download(nil, hc); err != nil {
		return err
	}

	if err := GetCovers(nil); err != nil {
		return err
	}

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), nil, data.SyncCompletedProperty)
	if err != nil {
		return err
	}

	tnu := time.Now().UTC().Unix()

	return rxa.ReplaceValues(data.SyncCompletedProperty, data.SyncCompletedProperty, strconv.FormatInt(tnu, 10))
}
