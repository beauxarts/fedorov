package data

import (
	"github.com/beauxarts/litres_integration"
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
	SequenceNameProperty      = "sequence-name"
	SequenceNumberProperty    = "sequence-number"
	DateReleasedProperty      = "date-released"
	DateTranslatedProperty    = "date-translated"
	DateCreatedProperty       = "date-created"
	AgeRatingProperty         = "age-rating"
	VolumeProperty            = "volume"
	DurationProperty          = "duration"
	ISBNPropertyProperty      = "isbn"
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
	// aggregate
	AnyTextProperty = "any-text"
	// sorting
	SortProperty       = "sort"
	DescendingProperty = "desc"
	// sync events
	SyncCompletedProperty = "sync-completed"
	// imported
	ImportedProperty = "imported"
)

func ReduxProperties() []string {
	return []string{
		MyBooksIdsProperty,
		HrefProperty,
		TitleProperty,
		DescriptionProperty,
		DownloadLinksProperty,
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
		ISBNPropertyProperty,
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
	"ISBN:":                                    ISBNPropertyProperty,
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
