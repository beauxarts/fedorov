package cli

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathology"
	"net/url"
)

func CascadeHandler(_ *url.URL) error {
	return Cascade()
}

func Cascade() error {

	ca := nod.Begin("cascading reductions...")
	defer ca.End()

	props := []string{data.TitleProperty, data.BookCompletedProperty, data.MyBooksIdsProperty, data.MyBooksOrderProperty}

	absReduxDir, err := pathology.GetAbsRelDir(data.Redux)
	if err != nil {
		return ca.EndWithError(err)
	}

	rdx, err := kvas.ReduxWriter(absReduxDir, props...)
	if err != nil {
		return ca.EndWithError(err)
	}

	// cascading data.BookCompletedProperty
	bcpa := nod.NewProgress(" " + data.BookCompletedProperty)
	defer bcpa.End()

	ids := rdx.Keys(data.TitleProperty)
	bcpa.TotalInt(len(ids))

	for _, id := range ids {
		bcpa.Increment()
		if val, ok := rdx.GetFirstVal(data.BookCompletedProperty, id); ok && val != "" {
			continue
		}
		if err := rdx.ReplaceValues(data.BookCompletedProperty, id, "false"); err != nil {
			return ca.EndWithError(err)
		}
	}

	bcpa.EndWithResult("done")

	// cascading data.MyBooksOrderProperty

	mboa := nod.NewProgress(" " + data.MyBooksOrderProperty)
	defer mboa.End()

	myBooksIds, _ := rdx.GetAllValues(data.MyBooksIdsProperty, data.MyBooksIdsProperty)
	mboa.TotalInt(len(myBooksIds))

	order := make(map[string][]string)
	for i, id := range myBooksIds {
		order[id] = []string{fmt.Sprintf("%9d", i)}
	}
	if err := rdx.BatchReplaceValues(data.MyBooksOrderProperty, order); err != nil {
		return mboa.EndWithError(err)
	}
	mboa.EndWithResult("done")

	ca.EndWithResult("done")

	return nil
}
