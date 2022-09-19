package data

import "strings"

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

var FormatDescriptors = map[string]string{
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

func LinkFormat(link string) string {
	format := FormatUnknown

	if strings.HasSuffix(link, ".fb2.zip") {
		format = FormatFB2
	} else if strings.HasSuffix(link, ".ios.epub") {
		format = FormatIOSEPUB
	} else if strings.HasSuffix(link, ".epub") {
		format = FormatEPUB
	} else if strings.HasSuffix(link, ".txt.zip") {
		format = FormatTXTZIP
	} else if strings.HasSuffix(link, ".rtf.zip") {
		format = FormatRTF
	} else if strings.HasSuffix(link, ".a4.pdf") {
		format = FormatPDFA4
	} else if strings.HasSuffix(link, ".a6.pdf") {
		format = FormatPDFA6
	} else if strings.HasSuffix(link, ".pdf") {
		format = FormatPDF
	} else if strings.HasSuffix(link, ".mobi.prc") {
		format = FormatMOBI
	} else if strings.HasSuffix(link, ".fb3") {
		format = FormatFB3
	} else if strings.HasSuffix(link, ".txt") {
		format = FormatTXT
	} else if strings.HasSuffix(link, ".mp3.zip") {
		format = FormatMP3
	} else if strings.HasSuffix(link, ".mp3") {
		format = FormatMP3
	} else if strings.HasSuffix(link, ".m4b") {
		format = FormatMP4
	} else if strings.HasSuffix(link, ".zip") {
		format = FormatZIP
	} else if strings.HasSuffix(link, ".azw3") {
		format = FormatAZW3
	}

	return format
}

func FormatDesc(format string) string {
	return FormatDescriptors[format]
}
