package view_models

import "github.com/beauxarts/fedorov/data"

var DigestTitles = map[string]string{
	// book-type
	"pdf":   "PDF",
	"аудио": "Аудио",
	"текст": "Текст",
	// sort
	data.TitleProperty:          propertyTitles[data.TitleProperty],
	data.DateCreatedProperty:    propertyTitles[data.DateCreatedProperty],
	data.DateTranslatedProperty: propertyTitles[data.DateTranslatedProperty],
	data.DateReleasedProperty:   propertyTitles[data.DateReleasedProperty],
	// desc
	"true":  "Да",
	"false": "Нет",
}
