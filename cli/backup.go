package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/packer"
	"github.com/boggydigital/pasu"
	"net/url"
)

func BackupHandler(_ *url.URL) error {
	return Backup()
}

func Backup() error {

	ba := nod.NewProgress("creating metadata backup...")
	defer ba.End()

	absBackupsDir, err := pasu.GetAbsDir(data.Backups)
	if err != nil {
		return ba.EndWithError(err)
	}

	absReduxDir, err := pasu.GetAbsRelDir(data.Redux)
	if err != nil {
		return ba.EndWithError(err)
	}

	if err := packer.Pack(absReduxDir, absBackupsDir, ba); err != nil {
		return ba.EndWithError(err)
	}

	ba.EndWithResult("done")

	return nil
}
