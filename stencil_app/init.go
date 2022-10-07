package stencil_app

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/stencil"
)

const (
	appTitle       = "fedorov"
	appAccentColor = "gray"
)

func Init(rxa kvas.ReduxAssets) (*stencil.AppConfiguration, error) {

	app := stencil.NewAppConfig(appTitle, appAccentColor)

	app.SetNavigation(NavItems, NavIcons, NavHrefs)
	app.SetFooter(FooterLocation, FooterRepoUrl)

	if err := app.SetCommonConfiguration(
		BookLabels,
		nil,
		nil,
		data.TitleProperty,
		PropertyTitles,
		SectionTitles,
		DigestTitles,
		rxa); err != nil {
		return app, err
	}

	if err := app.SetListConfiguration(
		BooksProperties,
		BooksLabels,
		BookPath,
		data.IdProperty,
		CoverPath,
		rxa); err != nil {
		return app, err
	}

	if err := app.SetItemConfiguration(
		BookProperties,
		nil,
		BookLabels,
		BookSections,
		data.IdProperty,
		CoverPath,
		rxa); err != nil {
		return app, err
	}

	app.SetFormatterConfiguration(
		fmtLabel, fmtTitle, fmtHref, nil, nil, nil)

	if err := app.SetSearchConfiguration(
		SearchProperties,
		SearchScopes,
		SearchScopeQueries()); err != nil {
		return app, err
	}

	return app, nil
}
