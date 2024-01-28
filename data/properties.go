package data

const (
	IdProperty = "id"
	// arts history
	ArtsHistoryOrderProperty     = "arts-history-order"
	ArtsHistoryEventTimeProperty = "arts-history-event-time"
	TitleProperty                = "title"

	// legacy reduced from detail page
	//AuthorsProperty           = "authors"
	//CoauthorsProperty         = "coauthors"
	//DescriptionProperty       = "description"
	//DownloadLinksProperty     = "download-links"
	//DownloadTitlesProperty    = "download-titles"
	//SequenceNameProperty      = "sequence-name"
	//SequenceNumberProperty    = "sequence-number"
	//DateReleasedProperty      = "date-released"
	//DateTranslatedProperty    = "date-translated"
	//DateCreatedProperty       = "date-created"
	//AgeRatingProperty         = "age-rating"
	//VolumeProperty            = "volume"
	//DurationProperty          = "duration"
	//ISBNProperty              = "isbn"
	//TranslatorsProperty       = "translators"
	//ReadersProperty           = "readers"
	//IllustratorsProperty      = "illustrators"
	//CopyrightHoldersProperty  = "copyright-holders"
	//ComposersProperty         = "composers"
	//AdapterProperty           = "adapter"
	//PerformersProperty        = "performers"
	//DirectorsProperty         = "directors"
	//SoundDirectorsProperty    = "sound-directors"
	//PublishersProperty        = "publishers"
	//TotalSizeProperty         = "total-size"
	//TotalPagesProperty        = "total-pages"
	//MissingDetailsIdsProperty = "missing-details-ids"
	//BookTypeProperty          = "book-type"
	//GenresProperty            = "genres"
	//TagsProperty              = "tags"
	//ImageProperty             = "image"
	//LanguageProperty          = "language"

	// local properties
	LocalTagsProperty     = "local-tags"
	BookCompletedProperty = "book-completed"
	// sorting
	SortProperty       = "sort"
	DescendingProperty = "desc"
	// sync events
	SyncCompletedProperty = "sync-completed"
	// imported
	ImportedProperty   = "imported"
	DataSourceProperty = "data-source"
	// dehydrated images
	DehydratedListImageProperty         = "dehydrated-list-image"
	DehydratedListImageModifiedProperty = "dehydrated-list-image-modified"
	DehydratedItemImageProperty         = "dehydrated-item-image"
	DehydratedItemImageModifiedProperty = "dehydrated-item-image-modified"
)

func ReduxProperties() []string {
	return []string{
		ArtsHistoryOrderProperty,
		ArtsHistoryEventTimeProperty,
		//
		SyncCompletedProperty,
		ImportedProperty,
		DataSourceProperty,
		DehydratedListImageProperty,
		DehydratedItemImageProperty,
	}
}

func ImportedProperties() []string {
	return []string{
		ImportedProperty,
		DataSourceProperty,
	}
}

//var LiveLibPropertyMap = map[string]string{
//	livelib_integration.TitleProperty:          TitleProperty,
//	livelib_integration.AuthorsProperty:        AuthorsProperty,
//	livelib_integration.TranslatorsProperty:    TranslatorsProperty,
//	livelib_integration.DescriptionProperty:    DescriptionProperty,
//	livelib_integration.ISBNProperty:           ISBNProperty,
//	livelib_integration.SequenceNameProperty:   SequenceNameProperty,
//	livelib_integration.SequenceNumberProperty: SequenceNumberProperty,
//	livelib_integration.GenresProperty:         GenresProperty,
//	livelib_integration.TagsProperty:           TagsProperty,
//	livelib_integration.AgeRatingProperty:      AgeRatingProperty,
//	livelib_integration.PublishersProperty:     PublishersProperty,
//	livelib_integration.YearPublishedProperty:  DateCreatedProperty,
//	livelib_integration.LanguageProperty:       LanguageProperty,
//	//livelib_integration.ImageProperty
//	//livelib_integration.EditionSeriesProperty
//}
