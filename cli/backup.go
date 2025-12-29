package cli

import (
	"net/url"

	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/backups"
	"github.com/boggydigital/nod"
)

func BackupHandler(_ *url.URL) error {
	return Backup()
}

func Backup() error {

	ba := nod.NewProgress("backing up metadata...")
	defer ba.Done()

	absBackupsDir := data.Pwd.AbsDirPath(data.Backups)
	absMetadataDir := data.Pwd.AbsDirPath(data.Metadata)

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
