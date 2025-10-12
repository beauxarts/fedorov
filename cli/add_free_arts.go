package cli

import (
	"errors"
	"iter"
	"net/url"
	"strings"
	"time"

	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathways"
	"github.com/boggydigital/redux"
)

func AddFreeArtsHandler(u *url.URL) error {

	q := u.Query()

	var ids []string
	if q.Has("id") {
		ids = strings.Split(q.Get("id"), ",")
	}

	return AddFreeArts(ids...)
}

func AddFreeArts(ids ...string) error {

	afaa := nod.Begin("adding free arts...")
	defer afaa.Done()

	if len(ids) == 0 {
		return errors.New("add-free-arts requires ids to add")
	}

	reduxDir, err := pathways.GetAbsRelDir(data.Redux)
	if err != nil {
		return err
	}

	rdx, err := redux.NewWriter(reduxDir, data.FreeArtsProperty)
	if err != nil {
		return err
	}

	freeArtsAdded := make(map[string][]string)
	for _, artId := range ids {
		freeArtsAdded[artId] = []string{time.Now().UTC().Format(time.RFC3339)}
	}

	if err = rdx.BatchReplaceValues(data.FreeArtsProperty, freeArtsAdded); err != nil {
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
