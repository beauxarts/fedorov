package rest

import (
	"crypto/sha256"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/view_models"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/middleware"
	"github.com/boggydigital/stencil"
	"html/template"
	"io/fs"
)

var (
	rxa  kvas.ReduxAssets
	tmpl *template.Template
	app  *stencil.ReduxApp
)

func SetUsername(u string) {
	middleware.SetUsername(sha256.Sum256([]byte(u)))
}

func SetPassword(p string) {
	middleware.SetPassword(sha256.Sum256([]byte(p)))
}

func InitTemplates(templatesFS fs.FS, appTemplates fs.FS) {
	tmpl = template.Must(
		template.
			New("").
			Funcs(view_models.FuncMap()).
			ParseFS(templatesFS, "templates/*.gohtml"))

	stencil.InitAppTemplates(appTemplates, "app_css/app_css.gohtml")
}

var booksListProperties = []string{
	data.TitleProperty,
	data.BookTypeProperty,
	data.AuthorsProperty,
	data.DateCreatedProperty,
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
	rxa, err = kvas.ConnectReduxAssets(data.AbsReduxDir(), fbr, data.ReduxProperties()...)
	if err != nil {
		return err
	}

	app = stencil.NewApp("fedorov", "📇", rxa)

	app.SetNavigation(
		[]string{"Книги", "Поиск"},
		map[string]string{
			"Книги": "stack",
			"Поиск": "search",
		},
		map[string]string{
			"Книги": "/books",
			"Поиск": "/search",
		})

	app.SetTitles(view_models.PropertyTitles, view_models.DigestTitles)

	if err := app.SetListParams("/book?id=", booksListProperties); err != nil {
		return err
	}

	app.SetSearchParams(view_models.SearchProperties)

	return err
}
