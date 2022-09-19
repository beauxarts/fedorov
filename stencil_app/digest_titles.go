package stencil_app

import "github.com/beauxarts/fedorov/data"

var DigestTitles = map[string]string{
	// book-type
	BookTypePDF:   BookTypeTitles[BookTypePDF],
	BookTypeAudio: BookTypeTitles[BookTypeAudio],
	BookTypeText:  BookTypeTitles[BookTypeText],
	// sort
	data.TitleProperty:          PropertyTitles[data.TitleProperty],
	data.DateCreatedProperty:    PropertyTitles[data.DateCreatedProperty],
	data.DateTranslatedProperty: PropertyTitles[data.DateTranslatedProperty],
	data.DateReleasedProperty:   PropertyTitles[data.DateReleasedProperty],
	// desc
	"true":  "Да",
	"false": "Нет",
}
