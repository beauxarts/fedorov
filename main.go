package main

import (
	"bytes"
	_ "embed"
	"github.com/beauxarts/fedorov/cli"
	"github.com/beauxarts/fedorov/clo_delegates"
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/clo"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathways"
	"log"
	"os"
)

const (
	dirsOverrideFilename = "directories.txt"
)

var (
	//go:embed "cli-commands.txt"
	cliCommands []byte
	//go:embed "cli-help.txt"
	cliHelp []byte
)

func main() {

	nod.EnableStdOutPresenter()

	ns := nod.NewProgress("fedorov is serving your DRM-free books")
	defer ns.Done()

	if err := pathways.Setup(
		dirsOverrideFilename,
		data.DefaultFedorovRootDir,
		data.RelToAbsDirs,
		data.AllAbsDirs...); err != nil {
		log.Fatalln(err)
	}

	defs, err := clo.Load(
		bytes.NewBuffer(cliCommands),
		bytes.NewBuffer(cliHelp),
		clo_delegates.Values)
	if err != nil {
		log.Fatalln(err)
	}

	clo.HandleFuncs(map[string]clo.Handler{
		"backup":                     cli.BackupHandler,
		"cascade":                    cli.CascadeHandler,
		"dehydrate":                  cli.DehydrateHandler,
		"download-litres-books":      cli.DownloadLitResBooksHandler,
		"download-litres-covers":     cli.DownloadLitResCoversHandler,
		"get-litres-arts":            cli.GetLitResArtsHandler,
		"get-litres-authors":         cli.GetLitResAuthorsHandler,
		"get-litres-contents":        cli.GetLitResContentsHandler,
		"get-litres-operations":      cli.GetLitResOperationsHandler,
		"get-litres-series":          cli.GetLitResSeriesHandler,
		"get-recent-arts":            cli.GetRecentArtsHandler,
		"get-recent-persons":         cli.GetRecentPersonsHandler,
		"get-recent-series":          cli.GetRecentSeriesHandler,
		"get-session-id":             cli.GetSessionIdHandler,
		"get-videos-metadata":        cli.GetVideosMetadataHandler,
		"has-arts":                   cli.HasArtsHandler,
		"migrate":                    cli.MigrateHandler,
		"reduce-litres-arts-details": cli.ReduceLitResArtsDetailsHandler,
		"reduce-litres-operations":   cli.ReduceLitResOperationsHandler,
		"serve":                      cli.ServeHandler,
		"sync":                       cli.SyncHandler,
		"version":                    cli.VersionHandler,
	})

	if err = defs.AssertCommandsHaveHandlers(); err != nil {
		log.Fatalln(err)
	}

	if err = defs.Serve(os.Args[1:]); err != nil {
		log.Fatalln(err)
	}
}
