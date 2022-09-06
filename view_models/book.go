package view_models

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/litres_integration"
	"github.com/boggydigital/kvas"
)

var detailsPropertyOrder = []string{
	data.AuthorsProperty,
	data.CoauthorsProperty,
	data.TranslatorsProperty,
	data.ReadersProperty,
	data.IllustratorsProperty,
	data.ComposersProperty,
	data.AdapterProperty,
	data.PerformersProperty,
	data.DirectorsProperty,
	data.SoundDirectorsProperty,
	data.CopyrightHoldersProperty,
	data.PublishersProperty,
	data.HrefProperty,
	data.SequenceNameProperty,
	data.SequenceNumberProperty,
	data.DateReleasedProperty,
	data.DateTranslatedProperty,
	data.DateCreatedProperty,
	data.AgeRatingProperty,
	data.DurationProperty,
	data.VolumeProperty,
	data.ISBNPropertyProperty,
	data.TotalSizeProperty,
	data.TotalPagesProperty,
}

var propertyTitles = map[string]string{
	data.TitleProperty:            "Название",
	data.AuthorsProperty:          "Автор(ы)",
	data.CoauthorsProperty:        "Cоавтор(ы)",
	data.TranslatorsProperty:      "Переводчик(и)",
	data.ReadersProperty:          "Чтец(ы)",
	data.IllustratorsProperty:     "Иллюстратор(ы)",
	data.ComposersProperty:        "Композитор(ы)",
	data.AdapterProperty:          "Адаптация",
	data.PerformersProperty:       "Исполнители",
	data.DirectorsProperty:        "Режиссер(ы)",
	data.SoundDirectorsProperty:   "Звукорежиссер(ы)",
	data.CopyrightHoldersProperty: "Правообладатели",
	data.PublishersProperty:       "Издатели",
	data.HrefProperty:             "ЛитРес",
	data.SequenceNameProperty:     "Серия",
	data.SequenceNumberProperty:   "Номер в серии",
	data.DateReleasedProperty:     "Опубликовано",
	data.DateTranslatedProperty:   "Переведено",
	data.DateCreatedProperty:      "Написано",
	data.AgeRatingProperty:        "Возрастной рейтинг",
	data.DurationProperty:         "Длительность",
	data.VolumeProperty:           "Объем",
	data.ISBNPropertyProperty:     "ISBN",
	data.TotalSizeProperty:        "Общий размер",
	data.TotalPagesProperty:       "Всего страниц",
}

type Book struct {
	Id    string
	Title string

	Properties      map[string]map[string]string
	PropertyOrder   []string
	PropertyTitles  map[string]string
	PropertyClasses map[string]string
}

func NewBook(id string, rxa kvas.ReduxAssets) *Book {

	bvm := &Book{
		Id:             id,
		Properties:     make(map[string]map[string]string),
		PropertyOrder:  detailsPropertyOrder,
		PropertyTitles: propertyTitles,
	}

	bvm.Title, _ = rxa.GetFirstVal(data.TitleProperty, id)

	for _, p := range detailsPropertyOrder {
		bvm.Properties[p] = getPropertyLinks(id, p, rxa)
	}

	return bvm
}

func getPropertyLinks(id string, property string, rxa kvas.ReduxAssets) map[string]string {

	propertyLinks := make(map[string]string)

	values, _ := rxa.GetAllUnchangedValues(property, id)

	for _, value := range values {

		linkTitle := formatPropertyLinkTitle(property, value)
		propertyLinks[linkTitle] = formatPropertyLinkHref(property, value)
	}

	return propertyLinks
}

func formatPropertyLinkTitle(property, link string) string {
	title := link

	switch property {
	case data.HrefProperty:
		title = "Ссылка"
	default:
		// do nothing
	}

	return title
}

func formatPropertyLinkHref(property, link string) string {
	switch property {
	case data.HrefProperty:
		return litres_integration.HrefUrl(link).String()
	default:
		// do nothing
	}
	return fmt.Sprintf("/search?%s=%s", property, link)
}
