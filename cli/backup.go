package cli

import (
	"net/url"

	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/camino"
	"github.com/boggydigital/nod"
)

func BackupHandler(_ *url.URL) error {
	return Backup()
}

func Backup() error {

	ba := nod.NewProgress("backing up metadata...")
	defer ba.Done()

	if err := camino.Compress(data.Metadata, data.Backups); err != nil {
		return err
	}

	cba := nod.NewProgress("cleaning up old backups...")
	defer cba.Done()

	if err := camino.CleanupTimed(data.Backups, true); err != nil {
		return err
	}

	return nil
}
