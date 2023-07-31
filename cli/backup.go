package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/packer"
	"net/url"
	"os"
)

func BackupHandler(_ *url.URL) error {
	return Backup()
}

func Backup() error {

	ba := nod.NewProgress("creating metadata backup...")
	defer ba.End()

	if _, err := os.Stat(data.AbsBackupDir()); os.IsNotExist(err) {
		if err := os.MkdirAll(data.AbsBackupDir(), 0755); err != nil {
			return ba.EndWithError(err)
		}
	}

	if err := packer.Pack(data.Pwd(), data.AbsReduxDir(), data.AbsBackupDir(), ba); err != nil {
		return ba.EndWithError(err)
	}

	ba.EndWithResult("done")

	return nil
}
