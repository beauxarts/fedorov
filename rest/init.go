package rest

import (
	"crypto/sha256"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/stencil_app"
	"github.com/beauxarts/fedorov/view_models"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/middleware"
	"github.com/boggydigital/pathology"
	"github.com/boggydigital/stencil"
	"html/template"
	"io/fs"
)

const (
	SearchResultsLimit = 24 // divides by 2,3,4,6 to allow that many columns
)

var (
	rdx  kvas.ReadableRedux
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

	absReduxDir, err := pathology.GetAbsRelDir(data.Redux)
	if err != nil {
		return err
	}

	if rdx, err = kvas.ReduxReader(absReduxDir, data.ReduxProperties()...); err != nil {
		return err
	}

	if app, err = stencil_app.Init(rdx); err != nil {
		return err
	}

	return err
}
