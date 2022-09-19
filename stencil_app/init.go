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

func Init(rxa kvas.ReduxAssets) (*stencil.ReduxApp, error) {

	app := stencil.NewApp("fedorov", "gray", rxa)

	app.SetFooter("New York, 🇺🇸", "https://github.com/beauxarts")

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

	app.SetTitles(data.TitleProperty, PropertyTitles, SectionTitles, DigestTitles)

	if err := app.SetLabels(BookLabels); err != nil {
		return app, err
	}

	rf := &rdxFormatter{rxa: rxa}

	app.SetLinkParams("/book", "/cover", rf.fmtTitle, rf.fmtHref)

	if err := app.SetListParams(BooksProperties); err != nil {
		return app, err
	}
	if err := app.SetItemParams(BookProperties, BookSections); err != nil {
		return app, err
	}

	app.SetSearchParams(SearchProperties)

	return app, nil
}

var caser = cases.Title(language.Russian)

type rdxFormatter struct {
	rxa kvas.ReduxAssets
}

func (rf *rdxFormatter) fmtSequenceNameNumber(id, name string) string {
	if err := rf.rxa.IsSupported(data.SequenceNameProperty, data.SequenceNumberProperty); err != nil {
		return name
	}

	names, _ := rf.rxa.GetAllUnchangedValues(data.SequenceNameProperty, id)
	numbers, _ := rf.rxa.GetAllUnchangedValues(data.SequenceNumberProperty, id)

	for ii, sn := range names {
		if sn == name {
			if snum := numbers[ii]; snum != "" {
				return fmt.Sprintf("%s %s", name, snum)
			}
		}
	}

	return name
}

func (rf *rdxFormatter) fmtTitle(id, property, link string) string {
	title := link

	switch property {
	case data.BookCompletedProperty:
		if link == "true" {
			title = "Прочитано"
		} else {
			title = "Не прочитано"
		}
	case data.SequenceNameProperty:
		title = rf.fmtSequenceNameNumber(id, link)
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

func (rf *rdxFormatter) fmtHref(id, property, link string) string {
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
