package view_models

import "path/filepath"

const (
	// Text books
	FormatFB2     = "FB2"
	FormatEPUB    = "EPUB"
	FormatTXTZIP  = "TXT ZIP"
	FormatRTF     = "RTF"
	FormatPDFA4   = "PDF A4"
	FormatMOBI    = "MOBI"
	FormatAZW3    = "AZW3"
	FormatIOSEPUB = "iOS.EPUB"
	FormatPDFA6   = "PDF A6"
	FormatFB3     = "FB3"
	FormatTXT     = "TXT"
	// PDF books
	FormatPDF = "PDF"
	// Audiobooks
	FormatMP3 = "MP3"
	FormatMP4 = "MP4"
	// Other
	FormatZIP     = "ZIP"
	FormatUnknown = "UNKNOWN"
)

var formatDescriptors = map[string]string{
	FormatFB2:     "Читается всеми российскими электронными книгами и многими программами",
	FormatEPUB:    "Подходит для Apple Books и большинства приложений для чтения",
	FormatTXTZIP:  "Можно открыть на любом компьютере",
	FormatRTF:     "Можно открыть на любом компьютере",
	FormatPDFA4:   "Открывается в программах Adobe Reader, Preview.app",
	FormatMOBI:    "Подходит для электронных книг Amazon Kindle",
	FormatAZW3:    "Подходит для электронных книг Amazon Kindle",
	FormatIOSEPUB: "EPUB, адаптированный для iPhone и iPad",
	FormatPDFA6:   "Оптимизирован под небольшие экраны",
	FormatFB3:     "Развитие формата FB2",
	FormatTXT:     "Можно открыть почти на любом устройстве",
	// PDF books
	FormatPDF: "Скачать в формате PDF",
	// Audiobooks
	FormatMP3: "Скачать книгу в стандартном качестве",
	FormatMP4: "Скачать версию для мобильного телефона",
	// Other
	FormatZIP:     "Zip архив",
	FormatUnknown: "Неизвестный формат",
}

type Downloads struct {
	Id    string
	Files []string
}

func NewDownloads(id string, links []string) *Downloads {
	dvm := &Downloads{
		Id:    id,
		Files: make([]string, 0, len(links)),
	}

	for _, link := range links {
		_, filename := filepath.Split(link)
		dvm.Files = append(dvm.Files, filename)
	}
	return dvm
}
