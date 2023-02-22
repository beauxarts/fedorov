package stencil_app

import "github.com/beauxarts/fedorov/data"

const (
	BookTypeText  = data.BookTypeText
	BookTypePDF   = data.BookTypePDF
	BookTypeAudio = data.BookTypeAudio
)

var BookTypeOrder = []string{
	BookTypeText,
	BookTypePDF,
	BookTypeAudio,
}

var BookTypeTitles = map[string]string{
	BookTypeText:  "Текст",
	BookTypePDF:   "PDF",
	BookTypeAudio: "Аудио",
}
