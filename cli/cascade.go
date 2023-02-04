package cli

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"net/url"
)

func CascadeHandler(_ *url.URL) error {
	return Cascade()
}

func Cascade() error {

	ca := nod.Begin("cascading reductions...")
	defer ca.End()

	props := []string{data.TitleProperty, data.BookCompletedProperty, data.MyBooksIdsProperty, data.MyBooksOrderProperty}

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), nil, props...)
	if err != nil {
		return ca.EndWithError(err)
	}

	// cascading data.BookCompletedProperty
	bcpa := nod.NewProgress(" " + data.BookCompletedProperty)
	defer bcpa.End()

	ids := rxa.Keys(data.TitleProperty)
	bcpa.TotalInt(len(ids))

	for _, id := range ids {
		bcpa.Increment()
		if val, ok := rxa.GetFirstVal(data.BookCompletedProperty, id); ok && val != "" {
			continue
		}
		if err := rxa.ReplaceValues(data.BookCompletedProperty, id, "false"); err != nil {
			return ca.EndWithError(err)
		}
	}

	bcpa.EndWithResult("done")

	// cascading data.MyBooksOrderProperty

	mboa := nod.NewProgress(" " + data.MyBooksOrderProperty)
	defer mboa.End()

	myBooksIds, _ := rxa.GetAllUnchangedValues(data.MyBooksIdsProperty, data.MyBooksIdsProperty)
	mboa.TotalInt(len(myBooksIds))

	order := make(map[string][]string)
	for i, id := range myBooksIds {
		order[id] = []string{fmt.Sprintf("%9d", i)}
	}
	if err := rxa.BatchReplaceValues(data.MyBooksOrderProperty, order); err != nil {
		return mboa.EndWithError(err)
	}
	mboa.EndWithResult("done")

	ca.EndWithResult("done")

	return nil
}
