package rest

import (
	"crypto/sha256"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/view_models"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/middleware"
	"html/template"
	"io/fs"
)

var (
	rxa  kvas.ReduxAssets
	tmpl *template.Template
)

func SetUsername(u string) {
	middleware.SetUsername(sha256.Sum256([]byte(u)))
}

func SetPassword(p string) {
	middleware.SetPassword(sha256.Sum256([]byte(p)))
}

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
