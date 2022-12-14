package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
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

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), nil, data.ReduxProperties()...)
	if err != nil {
		return ea.EndWithError(err)
	}

	file, err := os.Create(data.AbsExportFilename())
	defer file.Close()
	if err != nil {
		return ea.EndWithError(err)
	}

	if err := rxa.Export(file, ids...); err != nil {
		return ea.EndWithError(err)
	}

	ea.EndWithResult("done")

	return nil
}
