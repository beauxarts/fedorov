package main

import (
	"fmt"
	"github.com/beauxarts/fedorov/cli"
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"strings"
)

func main() {

	nod.EnableStdOutPresenter()

	ns := nod.NewProgress("fedorov is serving your DRM-free books")
	defer ns.End()

	if err := cli.Sync(); err != nil {
		panic(err)
	}

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), nil, data.ReduxProperties()...)

	if err != nil {
		panic(err)
	}

	q := map[string][]string{
		data.SequenceNameProperty: {"эмиль и марго"},
	}

	fmt.Println()

	for id := range rxa.Match(q, true) {
		title, _ := rxa.GetFirstVal(data.TitleProperty, id)
		authors, _ := rxa.GetAllValues(data.AuthorsProperty, id)
		href, _ := rxa.GetFirstVal(data.HrefProperty, id)
		created, _ := rxa.GetFirstVal(data.DateCreatedProperty, id)
		//dls, _ := rxa.GetAllValues(data.DownloadLinksProperty, id)

		fmt.Println(id, title, strings.Join(authors, ","))
		fmt.Println(created)
		fmt.Println(href)

		//for _, dl := range dls {
		//	fmt.Println(dl)
		//}
	}

}
