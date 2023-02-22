package data

import (
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/beauxarts/scrinium/livelib_integration"
)

const (
	IdProperty = "id"
	//reduced from the list
	MyBooksIdsProperty = "my-books-ids"
	HrefProperty       = "href"
	// reduced from detail page
	TitleProperty             = "title"
	AuthorsProperty           = "authors"
	CoauthorsProperty         = "coauthors"
	DescriptionProperty       = "description"
	DownloadLinksProperty     = "download-links"
	DownloadTitlesProperty    = "download-titles"
	SequenceNameProperty      = "sequence-name"
	SequenceNumberProperty    = "sequence-number"
	DateReleasedProperty      = "date-released"
	DateTranslatedProperty    = "date-translated"
	DateCreatedProperty       = "date-created"
	AgeRatingProperty         = "age-rating"
	VolumeProperty            = "volume"
	DurationProperty          = "duration"
	ISBNProperty              = "isbn"
	TranslatorsProperty       = "translators"
	ReadersProperty           = "readers"
	IllustratorsProperty      = "illustrators"
	CopyrightHoldersProperty  = "copyright-holders"
	ComposersProperty         = "composers"
	AdapterProperty           = "adapter"
	PerformersProperty        = "performers"
	DirectorsProperty         = "directors"
	SoundDirectorsProperty    = "sound-directors"
	PublishersProperty        = "publishers"
	TotalSizeProperty         = "total-size"
	TotalPagesProperty        = "total-pages"
	MissingDetailsIdsProperty = "missing-details-ids"
	BookTypeProperty          = "book-type"
	GenresProperty            = "genres"
	TagsProperty              = "tags"
	LocalTagsProperty         = "local-tags"
	BookCompletedProperty     = "book-completed"
	ImageProperty             = "image"
	LanguageProperty          = "language"
	// aggregate
	AnyTextProperty = "any-text"
	// sorting
	SortProperty       = "sort"
	DescendingProperty = "desc"
	// sync events
	SyncCompletedProperty = "sync-completed"
	// imported
	ImportedProperty   = "imported"
	DataSourceProperty = "data-source"
	// my-books-order
	MyBooksOrderProperty = "my-books-order"
)

func ReduxProperties() []string {
	return []string{
		MyBooksIdsProperty,
		HrefProperty,
		TitleProperty,
		DescriptionProperty,
		DownloadLinksProperty,
		DownloadTitlesProperty,
		AuthorsProperty,
		CoauthorsProperty,
		SequenceNameProperty,
		SequenceNumberProperty,
		DateReleasedProperty,
		DateTranslatedProperty,
		DateCreatedProperty,
		AgeRatingProperty,
		DurationProperty,
		VolumeProperty,
		ISBNProperty,
		TranslatorsProperty,
		ReadersProperty,
		IllustratorsProperty,
		CopyrightHoldersProperty,
		ComposersProperty,
		AdapterProperty,
		PerformersProperty,
		DirectorsProperty,
		SoundDirectorsProperty,
		PublishersProperty,
		TotalSizeProperty,
		TotalPagesProperty,
		MissingDetailsIdsProperty,
		BookTypeProperty,
		BookCompletedProperty,
		GenresProperty,
		TagsProperty,
		LocalTagsProperty,
		SyncCompletedProperty,
		ImportedProperty,
		DataSourceProperty,
		MyBooksOrderProperty,
		LanguageProperty,
	}
}

func AnyTextProperties() []string {
	return []string{
		TitleProperty,
		AuthorsProperty,
		CoauthorsProperty,
		DescriptionProperty,
		SequenceNameProperty,
		TranslatorsProperty,
		ReadersProperty,
		IllustratorsProperty,
		CopyrightHoldersProperty,
		ComposersProperty,
		AdapterProperty,
		PerformersProperty,
		DirectorsProperty,
		SoundDirectorsProperty,
		PublishersProperty,
		GenresProperty,
		TagsProperty,
	}
}

var LitResPropertyMap = map[string]string{
	litres_integration.TitleProperty:          TitleProperty,
	litres_integration.TypeProperty:           BookTypeProperty,
	litres_integration.AuthorsProperty:        AuthorsProperty,
	litres_integration.DownloadLinksProperty:  DownloadLinksProperty,
	litres_integration.DescriptionProperty:    DescriptionProperty,
	litres_integration.SequenceNameProperty:   SequenceNameProperty,
	litres_integration.SequenceNumberProperty: SequenceNumberProperty,
	litres_integration.GenresProperty:         GenresProperty,
	litres_integration.TagsProperty:           TagsProperty,
	"Соавтор:":                                CoauthorsProperty,
	"Возрастное ограничение:":                 AgeRatingProperty,
	"Объем:":                                   VolumeProperty,
	"Длительность:":                            DurationProperty,
	"Дата выхода на ЛитРес:":                   DateReleasedProperty,
	"Дата перевода:":                           DateTranslatedProperty,
	"Дата написания:":                          DateCreatedProperty,
	"ISBN:":                                    ISBNProperty,
	"Переводчики:":                             TranslatorsProperty,
	"Переводчик:":                              TranslatorsProperty,
	"Чтецы:":                                   ReadersProperty,
	"Чтец:":                                    ReadersProperty,
	"Художники:":                               IllustratorsProperty,
	"Художник:":                                IllustratorsProperty,
	"Правообладатели:":                         CopyrightHoldersProperty,
	"Правообладатель:":                         CopyrightHoldersProperty,
	"Композиторы:":                             ComposersProperty,
	"Композитор:":                              ComposersProperty,
	"Адаптация:":                               AdapterProperty,
	"Исполнители:":                             PerformersProperty,
	"Режиссер:":                                DirectorsProperty,
	"Звукорежиссер:":                           SoundDirectorsProperty,
	"Издатели:":                                PublishersProperty,
	"Издатель:":                                PublishersProperty,
	"Общий размер:":                            TotalSizeProperty,
	"Общее кол-во страниц:":                    TotalPagesProperty,
	"Оглавление":                               litres_integration.KnownIrrelevantProperty,
	"Размер страницы:":                         litres_integration.KnownIrrelevantProperty,
	litres_integration.KnownIrrelevantProperty: litres_integration.KnownIrrelevantProperty,
}

var LiveLibPropertyMap = map[string]string{
	livelib_integration.TitleProperty:          TitleProperty,
	livelib_integration.AuthorsProperty:        AuthorsProperty,
	livelib_integration.TranslatorsProperty:    TranslatorsProperty,
	livelib_integration.DescriptionProperty:    DescriptionProperty,
	livelib_integration.ISBNProperty:           ISBNProperty,
	livelib_integration.SequenceNameProperty:   SequenceNameProperty,
	livelib_integration.SequenceNumberProperty: SequenceNumberProperty,
	livelib_integration.GenresProperty:         GenresProperty,
	livelib_integration.TagsProperty:           TagsProperty,
	livelib_integration.AgeRatingProperty:      AgeRatingProperty,
	livelib_integration.PublishersProperty:     PublishersProperty,
	livelib_integration.YearPublishedProperty:  DateCreatedProperty,
	livelib_integration.LanguageProperty:       LanguageProperty,
	//livelib_integration.ImageProperty
	//livelib_integration.EditionSeriesProperty
}
