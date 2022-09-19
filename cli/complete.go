package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"net/url"
	"strings"
)

const (
	SetComplete   = "set"
	ClearComplete = "clear"
)

func CompleteHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	action := u.Query().Get("action")

	return Complete(ids, action)
}

func Complete(ids []string, action string) error {
	ca := nod.NewProgress("%s complete...", action)
	defer ca.End()

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), nil, data.TitleProperty, data.BookCompletedProperty)
	if err != nil {
		return err
	}

	ca.TotalInt(len(ids))

	for _, id := range ids {

		value := "true"
		if action == ClearComplete {
			value = "false"
		}

		if err := rxa.ReplaceValues(data.BookCompletedProperty, id, value); err != nil {
			return err
		}

		ca.Increment()
	}

	ca.EndWithResult("done")
	return nil
}
