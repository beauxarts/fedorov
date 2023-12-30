package main

import (
	"bytes"
	"embed"
	"github.com/beauxarts/fedorov/cli"
	"github.com/beauxarts/fedorov/clo_delegates"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest"
	"github.com/boggydigital/clo"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathology"
	"github.com/boggydigital/wits"
	"log"
	"os"
	"sync"
)

const (
	userDirsFilename = "directories.txt"
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

	rootDir   = "/var/lib/fedorov"
	reduxDir  = rootDir + "/_redux"
	coversDir = rootDir + "/covers"
)

func main() {

	// setup pathology dirs
	pathology.SetDefaultRootDir(data.DefaultFedorovRootDir)
	if err := pathology.SetAbsDirs(data.AllAbsDirs...); err != nil {
		panic(err)
	}
	if _, err := os.Stat(userDirsFilename); err == nil {
		udFile, err := os.Open(userDirsFilename)
		if err != nil {
			panic(err)
		}
		userDirs, err := wits.ReadKeyValue(udFile)
		if err != nil {
			panic(err)
		}
		pathology.SetUserDirsOverrides(userDirs)
	}
	pathology.SetRelToAbsDir(data.RelToAbsDirs)

	nod.EnableStdOutPresenter()

	once.Do(func() {
		rest.InitTemplates(templates, stencilAppStyles)
	})

	ns := nod.NewProgress("fedorov is serving your DRM-free books")
	defer ns.End()

	defs, err := clo.Load(
		bytes.NewBuffer(cliCommands),
		bytes.NewBuffer(cliHelp),
		clo_delegates.Values)
	if err != nil {
		log.Fatalln(err)
	}

	clo.HandleFuncs(map[string]clo.Handler{
		"backup":                 cli.BackupHandler,
		"cascade":                cli.CascadeHandler,
		"complete":               cli.CompleteHandler,
		"dehydrate":              cli.DehydrateHandler,
		"download-litres":        cli.DownloadLitResHandler,
		"export":                 cli.ExportHandler,
		"get-litres-arts":        cli.GetLitResArtsHandler,
		"get-litres-authors":     cli.GetLitResAuthorsHandler,
		"get-litres-covers":      cli.GetLitResCoversHandler,
		"get-litres-series":      cli.GetLitResSeriesHandler,
		"get-livelib-details":    cli.GetLiveLibDetailsHandler,
		"get-litres-my-books":    cli.GetLitResMyBooksHandler,
		"import":                 cli.ImportHandler,
		"post-completion":        cli.PostCompletionHandler,
		"purge":                  cli.PurgeHandler,
		"reduce-litres-my-books": cli.ReduceLitResMyBooksHandler,
		"serve":                  cli.ServeHandler,
		"sync":                   cli.SyncHandler,
		"version":                cli.VersionHandler,
	})

	if err := defs.AssertCommandsHaveHandlers(); err != nil {
		log.Fatalln(err)
	}

	if err := defs.Serve(os.Args[1:]); err != nil {
		log.Fatalln(err)
	}
}
