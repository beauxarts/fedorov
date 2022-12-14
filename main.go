package main

import (
	"bytes"
	"embed"
	"github.com/beauxarts/fedorov/cli"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest"
	"github.com/boggydigital/clo"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/wits"
	"log"
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
	//go:embed "stencil_app/styles/css.gohtml"
	stencilAppStyles embed.FS
	//go:embed "cli-commands.txt"
	cliCommands []byte
	//go:embed "cli-help.txt"
	cliHelp []byte
	rootDir = "/var/lib/fedorov"
)

func main() {

	nod.EnableStdOutPresenter()

	once.Do(func() {
		rest.InitTemplates(templates, stencilAppStyles)
	})

	ns := nod.NewProgress("fedorov is serving your DRM-free books")
	defer ns.End()

	if err := readUserDirectories(); err != nil {
		_ = ns.EndWithError(err)
		os.Exit(1)
	}

	data.ChRoot(rootDir)

	defs, err := clo.Load(
		bytes.NewBuffer(cliCommands),
		bytes.NewBuffer(cliHelp),
		nil)
	if err != nil {
		log.Fatalln(err)
	}

	clo.HandleFuncs(map[string]clo.Handler{
		"backup":          cli.BackupHandler,
		"complete":        cli.CompleteHandler,
		"download":        cli.DownloadHandler,
		"export":          cli.ExportHandler,
		"get-covers":      cli.GetCoversHandler,
		"get-details":     cli.GetDetailsHandler,
		"get-my-books":    cli.GetMyBooksHandler,
		"import":          cli.ImportHandler,
		"reduce-details":  cli.ReduceDetailsHandler,
		"reduce-my-books": cli.ReduceMyBooksHandler,
		"serve":           cli.ServeHandler,
		"sync":            cli.SyncHandler,
		"version":         cli.VersionHandler,
	})

	if err := defs.AssertCommandsHaveHandlers(); err != nil {
		log.Fatalln(err)
	}

	if err := defs.Serve(os.Args[1:]); err != nil {
		log.Fatalln(err)
	}
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
