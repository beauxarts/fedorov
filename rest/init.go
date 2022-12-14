package rest

import (
	"crypto/sha256"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/stencil_app"
	"github.com/beauxarts/fedorov/view_models"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/middleware"
	"github.com/boggydigital/stencil"
	"html/template"
	"io/fs"
)

const (
	SearchResultsLimit = 20
)

var (
	rxa  kvas.ReduxAssets
	tmpl *template.Template
	app  *stencil.AppConfiguration
)

func SetUsername(role, u string) {
	middleware.SetUsername(role, sha256.Sum256([]byte(u)))
}

func SetPassword(role, p string) {
	middleware.SetPassword(role, sha256.Sum256([]byte(p)))
}

func InitTemplates(templatesFS fs.FS, stencilAppStyles fs.FS) {
	tmpl = template.Must(
		template.
			New("").
			Funcs(view_models.FuncMap()).
			ParseFS(templatesFS, "templates/*.gohtml"))

	stencil.InitAppTemplates(stencilAppStyles, "stencil_app/styles/css.gohtml")
}

func Init() error {

	fbr := &kvas.ReduxFabric{
		Aggregates: map[string][]string{
			data.AnyTextProperty: data.AnyTextProperties(),
		},
		Transitives: nil,
		Atomics:     nil,
	}

	var err error
	if rxa, err = kvas.ConnectReduxAssets(data.AbsReduxDir(), fbr, data.ReduxProperties()...); err != nil {
		return err
	}

	if app, err = stencil_app.Init(rxa); err != nil {
		return err
	}

	return err
}
