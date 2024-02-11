package cli

import (
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

	props := []string{data.TitleProperty, data.BookCompletedProperty}

	rdx, err := data.NewReduxWriter(props...)
	if err != nil {
		return ca.EndWithError(err)
	}

	if err := cascadeBookCompletedProperty(rdx); err != nil {
		return ca.EndWithError(err)
	}

	ca.EndWithResult("done")

	return nil
}

func cascadeBookCompletedProperty(rdx kvas.WriteableRedux) error {

	bca := nod.NewProgress(" " + data.BookCompletedProperty)
	defer bca.End()

	if err := rdx.MustHave(data.TitleProperty, data.BookCompletedProperty); err != nil {
		return bca.EndWithError(err)
	}

	ids := rdx.Keys(data.TitleProperty)
	bca.TotalInt(len(ids))

	completed := make(map[string][]string)

	for _, id := range ids {
		bca.Increment()
		if val, ok := rdx.GetFirstVal(data.BookCompletedProperty, id); ok && val != "" {
			completed[id] = []string{"true"}
		}
		completed[id] = []string{"false"}
	}

	if err := rdx.BatchReplaceValues(data.BookCompletedProperty, completed); err != nil {
		return bca.EndWithError(err)
	}

	bca.EndWithResult("done")

	return nil
}
