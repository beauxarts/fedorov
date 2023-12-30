package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathology"
	"net/url"
	"os"
	"strings"
)

func ExportHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	return Export(ids)
}

func Export(ids []string) error {

	ea := nod.Begin("exporting books...")
	defer ea.End()

	absReduxDir, err := pathology.GetAbsRelDir(data.Redux)
	if err != nil {
		return ea.EndWithError(err)
	}

	rdx, err := kvas.NewReduxReader(absReduxDir, data.ReduxProperties()...)
	if err != nil {
		return ea.EndWithError(err)
	}

	absExportFilename, err := data.AbsExportFilename()
	if err != nil {
		return ea.EndWithError(err)
	}

	file, err := os.Create(absExportFilename)
	defer file.Close()
	if err != nil {
		return ea.EndWithError(err)
	}

	if err := rdx.Export(file, ids...); err != nil {
		return ea.EndWithError(err)
	}

	ea.EndWithResult("done")

	return nil
}
