package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/backups"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathways"
	"net/url"
)

func BackupHandler(_ *url.URL) error {
	return Backup()
}

func Backup() error {

	ba := nod.NewProgress("backing up metadata...")
	defer ba.End()

	absBackupsDir, err := pathways.GetAbsDir(data.Backups)
	if err != nil {
		return ba.EndWithError(err)
	}

	absMetadataDir, err := pathways.GetAbsDir(data.Metadata)
	if err != nil {
		return ba.EndWithError(err)
	}

	if err := backups.Compress(absMetadataDir, absBackupsDir); err != nil {
		return ba.EndWithError(err)
	}

	ba.EndWithResult("done")

	cba := nod.NewProgress("cleaning up old backups...")
	defer cba.End()

	if err := backups.Cleanup(absBackupsDir, true, cba); err != nil {
		return cba.EndWithError(err)
	}

	cba.EndWithResult("done")

	return nil
}
