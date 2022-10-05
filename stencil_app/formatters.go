package stencil_app

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/litres_integration"
	"github.com/boggydigital/kvas"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

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

func fmtLabel(_, property, link string, _ kvas.ReduxAssets) string {
	label := link
	switch property {
	case data.BookCompletedProperty:
		if link == "true" {
			label = "Прочитано"
		} else {
			label = "Не прочитано"
		}
	}
	return label
}

func fmtTitle(id, property, link string, rxa kvas.ReduxAssets) string {
	title := link

	switch property {
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
