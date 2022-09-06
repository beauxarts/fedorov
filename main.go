package main

import (
	"embed"
	"github.com/beauxarts/fedorov/cli"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/wits"
	"os"
	"sync"
)

const (
	directoriesFilename = "directories.txt"
)

var (
	once = sync.Once{}
	//go:embed "templates/*.gohtml"
	templates embed.FS
	rootDir   = "/var/lib/fedorov"
)

func main() {

	nod.EnableStdOutPresenter()

	once.Do(func() {
		rest.InitTemplates(templates)
	})

	ns := nod.NewProgress("fedorov is serving your DRM-free books")
	defer ns.End()

	if err := readUserDirectories(); err != nil {
		_ = ns.EndWithError(err)
		os.Exit(1)
	}

	data.ChRoot(rootDir)

	if err := cli.Serve(1520, true); err != nil {
		panic(err)
	}

	//if err := cli.Sync(); err != nil {
	//	panic(err)
	//}
}

func readUserDirectories() error {
	if _, err := os.Stat(directoriesFilename); os.IsNotExist(err) {
		return nil
	}

	udFile, err := os.Open(directoriesFilename)
	if err != nil {
		return err
	}

	dirs, err := wits.ReadKeyValue(udFile)
	if err != nil {
		return err
	}

	if sd, ok := dirs["root"]; ok {
		rootDir = sd
	}

	//validate that directories actually exist
	if _, err := os.Stat(rootDir); err != nil {
		return err
	}

	return nil
}
