package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/litres_integration"
	"github.com/boggydigital/coost"
	"net/url"
)

func SyncHandler(u *url.URL) error {
	return Sync()
}
func Sync() error {

	hc, err := coost.NewHttpClientFromFile(data.AbsCookiesFilename(), litres_integration.LitResHost)
	if err != nil {
		return err
	}

	if err := GetMyBooksFresh(hc); err != nil {
		return err
	}

	if err := ReduceMyBooksFresh(); err != nil {
		return err
	}

	if err := GetMyBooksDetails(nil, hc); err != nil {
		return err
	}

	if err := ReduceMyBooksDetails(); err != nil {
		return err
	}

	if err := Download(hc); err != nil {
		return err
	}

	if err := GetCovers(nil); err != nil {
		return err
	}

	return nil
}
