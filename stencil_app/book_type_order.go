package stencil_app

import (
	"github.com/beauxarts/scrinium/litres_integration"
)

const (
	BookTypeText  = litres_integration.ArtTypeText
	BookTypePDF   = litres_integration.ArtTypePDF
	BookTypeAudio = litres_integration.ArtTypeAudio
)

var BookTypeTitles = map[litres_integration.ArtType]string{
	BookTypeText:  "Текст",
	BookTypePDF:   "PDF",
	BookTypeAudio: "Аудио",
}

var BookTypeOrder = []string{
	BookTypeTitles[BookTypeText],
	BookTypeTitles[BookTypePDF],
	BookTypeTitles[BookTypeAudio],
}
