package compton_data

const (
	InformationSection   = "information"
	AnnotationSection    = "annotation"
	ExternalLinksSection = "external-links"
	SimilarSection       = "similar"
	ReviewsSection       = "reviews"
	QuotesSection        = "quotes"
	VideosSection        = "videos"
	ContentsSection      = "contents"
	FilesSection         = "files"
)

var SectionTitles = map[string]string{
	InformationSection:   "Информация",
	AnnotationSection:    "Аннотация",
	ExternalLinksSection: "Ссылки",
	ReviewsSection:       "Отзывы",
	QuotesSection:        "Цитаты",
	SimilarSection:       "Сходные книги",
	VideosSection:        "Видео",
	ContentsSection:      "Оглавление",
	FilesSection:         "Файлы",
}
