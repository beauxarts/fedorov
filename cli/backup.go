package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/packer"
	"github.com/boggydigital/pathology"
	"net/url"
	"os"
)

func BackupHandler(_ *url.URL) error {
	return Backup()
}

func Backup() error {

	ba := nod.NewProgress("creating metadata backup...")
	defer ba.End()

	absBackupsDir, err := pathology.GetAbsDir(data.Backups)
	if err != nil {
		return ba.EndWithError(err)
	}

	absReduxDir, err := pathology.GetAbsRelDir(data.Redux)
	if err != nil {
		return ba.EndWithError(err)
	}

	if _, err := os.Stat(absBackupsDir); os.IsNotExist(err) {
		if err := os.MkdirAll(absBackupsDir, 0755); err != nil {
			return ba.EndWithError(err)
		}
	}

	if err := packer.Pack(absReduxDir, absBackupsDir, ba); err != nil {
		return ba.EndWithError(err)
	}

	ba.EndWithResult("done")

	return nil
}
