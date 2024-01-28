package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pasu"
	"net/url"
)

func CascadeHandler(_ *url.URL) error {
	return Cascade()
}

func Cascade() error {

	ca := nod.Begin("cascading reductions...")
	defer ca.End()

	props := []string{data.TitleProperty, data.BookCompletedProperty, data.ArtsHistoryOrderProperty, data.MyBooksOrderProperty}

	absReduxDir, err := pasu.GetAbsRelDir(data.Redux)
	if err != nil {
		return ca.EndWithError(err)
	}

	rdx, err := kvas.NewReduxWriter(absReduxDir, props...)
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

	//mboa := nod.NewProgress(" " + data.MyBooksOrderProperty)
	//defer mboa.End()
	//
	//artsIds, _ := rdx.GetAllValues(data.ArtsHistoryOrderProperty, data.ArtsHistoryOrderProperty)
	//mboa.TotalInt(len(artsIds))
	//
	//order := make(map[string][]string)
	//for i, id := range artsIds {
	//	order[id] = []string{fmt.Sprintf("%9d", i)}
	//}
	//if err := rdx.BatchReplaceValues(data.MyBooksOrderProperty, order); err != nil {
	//	return mboa.EndWithError(err)
	//}
	//mboa.EndWithResult("done")

	ca.EndWithResult("done")

	return nil
}
