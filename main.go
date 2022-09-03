package main

import (
	"github.com/beauxarts/fedorov/cli"
	"github.com/boggydigital/nod"
)

func main() {

	nod.EnableStdOutPresenter()

	ns := nod.NewProgress("fedorov is serving your DRM-free books")
	defer ns.End()

	if err := cli.Sync(); err != nil {
		panic(err)
	}
}
