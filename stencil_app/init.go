package stencil_app

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/stencil"
)

const (
	appTitle        = "fedorov"
	appFavIconEmoji = "📇"
)

func Init(rdx kevlar.ReadableRedux) (*stencil.AppConfiguration, error) {

	app := stencil.NewAppConfig(appTitle, appFavIconEmoji)

	app.SetNavigation(NavItems, NavIcons, NavHrefs)
	app.SetFooter(FooterLocation, FooterRepoUrl)

	if err := app.SetCommonConfiguration(
		BookLabels,
		nil,
		nil,
		"", //data.TitleProperty,
		PropertyTitles,
		SectionTitles,
		rdx); err != nil {
		return app, err
	}

	if err := app.SetListConfiguration(
		BooksProperties,
		BooksLabels,
		BookPath,
		data.IdProperty,
		ListCoverPath,
		rdx); err != nil {
		return app, err
	}

	app.SetDehydratedImagesConfiguration(
		data.DehydratedListImageProperty,
		data.DehydratedItemImageProperty)

	if err := app.SetItemConfiguration(
		BookProperties,
		nil,
		BookHiddenProperties,
		BookSections,
		data.IdProperty,
		BookCoverPath,
		rdx); err != nil {
		return app, err
	}

	app.SetFormatterConfiguration(
		fmtLabel, fmtTitle, fmtHref, nil, fmtAction, fmtActionHref)

	if err := app.SetSearchConfiguration(
		nil,
		DigestProperties,
		SearchScopes,
		SearchScopeQueries()); err != nil {
		return app, err
	}

	return app, nil
}
