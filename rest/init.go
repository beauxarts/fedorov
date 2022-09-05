package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/view_models"
	"github.com/boggydigital/kvas"
	"html/template"
	"io/fs"
)

var (
	rxa  kvas.ReduxAssets
	tmpl *template.Template
)

func InitTemplates(templatesFS fs.FS) {
	tmpl = template.Must(
		template.
			New("").
			Funcs(view_models.FuncMap()).
			ParseFS(templatesFS, "templates/*.gohtml"))
}

func Init() error {
	var err error
	rxa, err = kvas.ConnectReduxAssets(data.AbsReduxDir(), nil, data.ReduxProperties()...)
	return err
}
