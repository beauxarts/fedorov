package stencil_app

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
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

	names, _ := rxa.GetAllValues(data.SequenceNameProperty, id)
	numbers, _ := rxa.GetAllValues(data.SequenceNumberProperty, id)

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
			return "Прочитано"
		} else {
			return ""
		}
	case data.ImportedProperty:
		if link == "true" {
			return "Импорт"
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
	case data.BookCompletedProperty:
		if link == "true" {
			return "Да"
		} else {
			return "Нет"
		}
	case data.ImportedProperty:
		if link == "true" {
			return "Да"
		} else {
			return "Нет"
		}
	default:
		// do nothing
	}

	return title
}

func fmtHref(_, property, link string, _ kvas.ReduxAssets) string {
	switch property {
	case data.SequenceNameProperty:
		fallthrough
	case data.AuthorsProperty:
		return fmt.Sprintf("/search?%s=%s&sort=date-created&desc=true", property, link)
	case data.TranslatorsProperty:
		return fmt.Sprintf("/search?%s=%s&sort=date-translated&desc=true", property, link)
	case data.HrefProperty:
		return litres_integration.HrefUrl(link).String()
	case data.AgeRatingProperty:
		fallthrough
	case data.ISBNProperty:
		fallthrough
	case data.TotalPagesProperty:
		fallthrough
	case data.TotalSizeProperty:
		fallthrough
	case data.VolumeProperty:
		fallthrough
	case data.DurationProperty:
		return ""
	case data.DehydratedListImageProperty:
		fallthrough
	case data.DehydratedItemImageProperty:
		return link
	default:
		// do nothing
	}
	return fmt.Sprintf("/search?%s=%s", property, link)
}

func fmtAction(id, property, link string, rxa kvas.ReduxAssets) string {
	switch property {
	case data.BookCompletedProperty:
		if link == "true" {
			return "Очистить"
		} else {
			return "Отметить"
		}
	case data.LocalTagsProperty:
		return "Изменить"
	}
	return ""
}

func fmtActionHref(id, property, link string, _ kvas.ReduxAssets) string {
	switch property {
	case data.BookCompletedProperty:
		switch link {
		case "Очистить":
			return "/completed/clear?id=" + id
		case "Отметить":
			return "/completed/set?id=" + id
		}
	case data.LocalTagsProperty:
		return "/local-tags/edit?id=" + id
	}
	return ""
}
