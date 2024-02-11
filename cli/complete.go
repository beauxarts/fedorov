package cli

import (
	"errors"
	"github.com/beauxarts/fedorov/data"
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

	rdx, err := data.NewReduxWriter(data.BookCompletedProperty)
	if err != nil {
		return err
	}

	ca.TotalInt(len(ids))

	for _, id := range ids {

		switch action {
		case SetComplete:
			if err := rdx.ReplaceValues(data.BookCompletedProperty, id, "true"); err != nil {
				return err
			}
		case ClearComplete:
			if err := rdx.CutValues(data.BookCompletedProperty, id, "true"); err != nil {
				return err
			}
		default:
			return errors.New("unknown compelte action " + action)
		}

		ca.Increment()
	}

	ca.EndWithResult("done")
	return nil
}
