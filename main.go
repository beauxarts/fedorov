package main

import (
	"embed"
	"github.com/beauxarts/fedorov/cli"
	"github.com/beauxarts/fedorov/rest"
	"github.com/boggydigital/nod"
	"sync"
)

var (
	once = sync.Once{}
	//go:embed "templates/*.gohtml"
	templates embed.FS
)

func main() {

	nod.EnableStdOutPresenter()

	once.Do(func() {
		rest.InitTemplates(templates)
	})

	ns := nod.NewProgress("fedorov is serving your DRM-free books")
	defer ns.End()

	if err := cli.Serve(1520, true); err != nil {
		panic(err)
	}

	//if err := cli.Sync(); err != nil {
	//	panic(err)
	//}
}
