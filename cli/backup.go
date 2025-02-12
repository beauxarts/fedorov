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
	defer ba.Done()

	absBackupsDir, err := pathways.GetAbsDir(data.Backups)
	if err != nil {
		return err
	}

	absMetadataDir, err := pathways.GetAbsDir(data.Metadata)
	if err != nil {
		return err
	}

	if err := backups.Compress(absMetadataDir, absBackupsDir); err != nil {
		return err
	}

	cba := nod.NewProgress("cleaning up old backups...")
	defer cba.Done()

	if err := backups.Cleanup(absBackupsDir, true, cba); err != nil {
		return err
	}

	return nil
}
