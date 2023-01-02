package cli

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/nod"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"time"
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

	bfn := fmt.Sprintf(
		"backup-%s.tar.gz",
		time.Now().Format("2006-01-02-15-04-05"))

	absBackupPath := filepath.Join(data.AbsBackupDir(), bfn)

	if _, err := os.Stat(absBackupPath); os.IsExist(err) {
		return ba.EndWithError(err)
	}

	file, err := os.Create(absBackupPath)
	if err != nil {
		return ba.EndWithError(err)
	}
	defer file.Close()

	gw := gzip.NewWriter(file)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	files := make([]string, 0)

	from := data.AbsReduxDir()

	if err := filepath.Walk(from, func(f string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		files = append(files, f)
		return nil
	}); err != nil {
		return ba.EndWithError(err)
	}

	ba.TotalInt(len(files))

	for _, f := range files {

		fi, err := os.Stat(f)
		if err != nil {
			return ba.EndWithError(err)
		}

		header, err := tar.FileInfoHeader(fi, f)
		if err != nil {
			return ba.EndWithError(err)
		}

		rp, err := filepath.Rel(data.Pwd(), f)
		if err != nil {
			return ba.EndWithError(err)
		}

		header.Name = filepath.ToSlash(rp)

		if err := tw.WriteHeader(header); err != nil {
			return ba.EndWithError(err)
		}

		of, err := os.Open(f)
		if err != nil {
			return ba.EndWithError(err)
		}

		if _, err := io.Copy(tw, of); err != nil {
			return ba.EndWithError(err)
		}

		ba.Increment()
	}

	ba.EndWithResult("%s is ready", absBackupPath)

	return nil
}
