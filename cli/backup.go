package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/konpo"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pasu"
	"net/url"
)

func BackupHandler(_ *url.URL) error {
	return Backup()
}

func Backup() error {

	ba := nod.NewProgress("backing up metadata...")
	defer ba.End()

	absBackupsDir, err := pasu.GetAbsDir(data.Backups)
	if err != nil {
		return ba.EndWithError(err)
	}

	absMetadataDir, err := pasu.GetAbsDir(data.Metadata)
	if err != nil {
		return ba.EndWithError(err)
	}

	if err := konpo.Compress(absMetadataDir, absBackupsDir); err != nil {
		return ba.EndWithError(err)
	}

	ba.EndWithResult("done")

	cba := nod.NewProgress("cleaning up old backups...")
	defer cba.End()

	if err := konpo.Cleanup(absBackupsDir, true, cba); err != nil {
		return cba.EndWithError(err)
	}

	cba.EndWithResult("done")

	return nil
}
