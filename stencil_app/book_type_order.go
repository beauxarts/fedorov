package stencil_app

const (
	BookTypeText  = "текст"
	BookTypePDF   = "pdf"
	BookTypeAudio = "аудио"
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
