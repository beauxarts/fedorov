package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathways"
	"net/url"
	"path/filepath"
)

func MigrateHandler(_ *url.URL) error {
	return Migrate()
}

func Migrate() error {
	ma := nod.Begin("migrating data...")
	defer ma.End()

	dir, err := pathways.GetAbsDir(data.Metadata)
	if err != nil {
		return err
	}

	metadataDirs := []string{
		"_redux",
		"arts-files", "arts-reviews", "arts-details", "arts-quotes", "arts-similar",
		"author-details", "author-similar",
		"contents",
		"litres-operations", "litres-history-log",
		"series-similar", "series-details",
	}

	for _, md := range metadataDirs {
		if err := kevlar.Migrate(filepath.Join(dir, md)); err != nil {
			return err
		}
	}

	ma.EndWithResult("done")

	return nil
}
