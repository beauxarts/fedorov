package cli

import (
	"errors"
	"iter"
	"net/url"
	"strings"

	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathways"
	"github.com/boggydigital/redux"
)

func FreeArtsHandler(u *url.URL) error {

	q := u.Query()

	var add []string
	if q.Has("add") {
		add = strings.Split(q.Get("add"), ",")
	}

	var remove []string
	if q.Has("remove") {
		remove = strings.Split(q.Get("remove"), ",")
	}

	return FreeArts(add, remove)
}

func FreeArts(add []string, remove []string) error {

	faa := nod.Begin("managing free arts...")
	defer faa.Done()

	if len(add) == 0 && len(remove) == 0 {
		return errors.New("free-arts requires ids to add or remove")
	}

	reduxDir, err := pathways.GetAbsRelDir(data.Redux)
	if err != nil {
		return err
	}

	rdx, err := redux.NewWriter(reduxDir, data.FreeArtsProperty)
	if err != nil {
		return err
	}

	if err = rdx.CutKeys(data.FreeArtsProperty, remove...); err != nil {
		return err
	}

	if len(add) == 0 {
		return nil
	}

	addedFreeArts := make(map[string][]string)
	for _, artId := range add {
		addedFreeArts[artId] = []string{artId}
	}

	if err = rdx.BatchAddValues(data.FreeArtsProperty, addedFreeArts); err != nil {
		return err
	}

	return nil
}

func getFreeArts() (iter.Seq[string], error) {
	gfaa := nod.Begin("getting free arts...")
	defer gfaa.Done()

	reduxDir, err := pathways.GetAbsRelDir(data.Redux)
	if err != nil {
		return nil, err
	}

	rdx, err := redux.NewWriter(reduxDir, data.FreeArtsProperty)
	if err != nil {
		return nil, err
	}

	return rdx.Keys(data.FreeArtsProperty), nil
}
