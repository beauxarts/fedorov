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
	"github.com/boggydigital/pasu"
	"os"
	"sync"
)

const (
	dirsOverrideFilename = "directories.txt"
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
)

func main() {

	nod.EnableStdOutPresenter()

	ns := nod.NewProgress("fedorov is serving your DRM-free books")
	defer ns.End()

	once.Do(func() {
		rest.InitTemplates(templates, stencilAppStyles)
	})

	if err := pasu.Setup(
		dirsOverrideFilename,
		data.DefaultFedorovRootDir,
		data.RelToAbsDirs,
		data.AllAbsDirs...); err != nil {
		_ = ns.EndWithError(err)
		os.Exit(1)
	}

	defs, err := clo.Load(
		bytes.NewBuffer(cliCommands),
		bytes.NewBuffer(cliHelp),
		clo_delegates.Values)
	if err != nil {
		_ = ns.EndWithError(err)
		os.Exit(1)
	}

	clo.HandleFuncs(map[string]clo.Handler{
		"backup":                    cli.BackupHandler,
		"cascade":                   cli.CascadeHandler,
		"complete":                  cli.CompleteHandler,
		"dehydrate":                 cli.DehydrateHandler,
		"download-litres":           cli.DownloadLitResHandler,
		"export":                    cli.ExportHandler,
		"get-litres-arts":           cli.GetLitResArtsHandler,
		"get-litres-authors":        cli.GetLitResAuthorsHandler,
		"get-litres-covers":         cli.GetLitResCoversHandler,
		"get-litres-history-log":    cli.GetLitResHistoryLogHandler,
		"get-litres-series":         cli.GetLitResSeriesHandler,
		"get-livelib-details":       cli.GetLiveLibDetailsHandler,
		"import":                    cli.ImportHandler,
		"post-completion":           cli.PostCompletionHandler,
		"purge":                     cli.PurgeHandler,
		"reduce-litres-history-log": cli.ReduceLitResHistoryLogHandler,
		"serve":                     cli.ServeHandler,
		"sync":                      cli.SyncHandler,
		"version":                   cli.VersionHandler,
	})

	if err := defs.AssertCommandsHaveHandlers(); err != nil {
		_ = ns.EndWithError(err)
		os.Exit(1)
	}

	//atr, err := data.NewArtsTypesReader(litres_integration.ArtsTypeFiles)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, id := range atr.Keys() {
	//
	//	fmt.Println(id)
	//
	//	ad, err := atr.ReadValue(id)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	fmt.Println(ad)
	//
	//}

	if err := defs.Serve(os.Args[1:]); err != nil {
		_ = ns.EndWithError(err)
		os.Exit(1)
	}
}
