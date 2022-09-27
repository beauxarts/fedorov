package stencil_app

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/litres_integration"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/stencil"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	appTitle       = "fedorov"
	appAccentColor = "gray"
)

func Init(rxa kvas.ReduxAssets) (*stencil.App, error) {

	app := stencil.NewApp(appTitle, appAccentColor)

	app.SetNavigation(NavItems, NavIcons, NavHrefs)
	app.SetFooter(FooterLocation, FooterRepoUrl)

	app.SetTitles(data.TitleProperty, PropertyTitles, SectionTitles, DigestTitles)
	if err := app.SetLabels(BookLabels, rxa); err != nil {
		return app, err
	}

	app.SetLinkParams(BookPath, CoverPath, fmtTitle, fmtHref, nil)

	if err := app.SetListParams(data.IdProperty, BooksProperties, nil, rxa); err != nil {
		return app, err
	}
	if err := app.SetItemParams(BookProperties, BookSections, rxa); err != nil {
		return app, err
	}
	if err := app.SetSearchParams(SearchScopes, SearchScopeQueries(), SearchProperties); err != nil {
		return app, err
	}

	return app, nil
}

var caser = cases.Title(language.Russian)

func fmtSequenceNameNumber(id, name string, rxa kvas.ReduxAssets) string {
	if err := rxa.IsSupported(
		data.SequenceNameProperty,
		data.SequenceNumberProperty); err != nil {
		return name
	}

	names, _ := rxa.GetAllUnchangedValues(data.SequenceNameProperty, id)
	numbers, _ := rxa.GetAllUnchangedValues(data.SequenceNumberProperty, id)

	for ii, sn := range names {
		if sn == name {
			if snum := numbers[ii]; snum != "" {
				return fmt.Sprintf("%s %s", name, snum)
			}
		}
	}

	return name
}

func fmtTitle(id, property, link string, rxa kvas.ReduxAssets) string {
	title := link

	switch property {
	case data.BookCompletedProperty:
		if link == "true" {
			title = "Прочитано"
		} else {
			title = "Не прочитано"
		}
	case data.SequenceNameProperty:
		title = fmtSequenceNameNumber(id, link, rxa)
	case data.GenresProperty:
		fallthrough
	case data.TagsProperty:
		title = caser.String(title)
	case data.HrefProperty:
		title = "ЛитРес"
	default:
		// do nothing
	}

	return title
}

func fmtHref(id, property, link string, rxa kvas.ReduxAssets) string {
	switch property {
	case data.AuthorsProperty:
		return fmt.Sprintf("/search?%s=%s&sort=date-created&desc=true", property, link)
	case data.TranslatorsProperty:
		return fmt.Sprintf("/search?%s=%s&sort=date-translated&desc=true", property, link)
	case data.HrefProperty:
		return litres_integration.HrefUrl(link).String()
	case data.AgeRatingProperty:
		fallthrough
	case data.ISBNPropertyProperty:
		fallthrough
	case data.TotalPagesProperty:
		fallthrough
	case data.TotalSizeProperty:
		fallthrough
	case data.VolumeProperty:
		fallthrough
	case data.DurationProperty:
		return ""
	default:
		// do nothing
	}
	return fmt.Sprintf("/search?%s=%s", property, link)
}
