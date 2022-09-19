package stencil_app

import "github.com/beauxarts/fedorov/data"

var DigestTitles = map[string]string{
	// book-type
	"pdf":   "PDF",
	"аудио": "Аудио",
	"текст": "Текст",
	// sort
	data.TitleProperty:          PropertyTitles[data.TitleProperty],
	data.DateCreatedProperty:    PropertyTitles[data.DateCreatedProperty],
	data.DateTranslatedProperty: PropertyTitles[data.DateTranslatedProperty],
	data.DateReleasedProperty:   PropertyTitles[data.DateReleasedProperty],
	// desc
	"true":  "Да",
	"false": "Нет",
}
