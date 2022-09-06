package view_models

const (
	// Text books
	FormatFB2     = "FB2"
	FormatEPUB    = "EPUB"
	FormatTXTZIP  = "TXT ZIP"
	FormatRTF     = "RTF"
	FormatPDFA4   = "PDF A4"
	FormatMOBI    = "MOBI"
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
	FormatFB2:     "Подходит для смартфонов, планшетов на Android, электронных книг (кроме Kindle) и многих программ",
	FormatEPUB:    "Подходит для устройств на iOS (iPhone, iPad, iMac) и большинства приложений для чтения",
	FormatTXTZIP:  "Можно открыть на любом компьютере",
	FormatRTF:     "Можно открыть на любом компьютере",
	FormatPDFA4:   "Открывается в программе Adobe Reader",
	FormatMOBI:    "Подходит для электронных книг Kindle и Android-приложений",
	FormatIOSEPUB: "Идеально подойдет для iPhone и iPad",
	FormatPDFA6:   "Оптимизирован и подойдет для смартфонов",
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
	Links        []string
	Availability map[string]bool
}

func NewDownloads(links []string, availability map[string]bool) *Downloads {
	return &Downloads{
		Links:        links,
		Availability: availability,
	}
}
