package data

const (
	IdProperty = "id"

	// arts history
	ArtsHistoryOrderProperty     = "arts-history-order"
	ArtsHistoryEventTimeProperty = "arts-history-event-time"

	// arts details properties
	CoverUrlProperty               = "cover-url"
	TitleProperty                  = "title"
	ArtTypeProperty                = "art-type"
	PriceProperty                  = "price"
	MinAgeProperty                 = "min-age"
	SymbolsCountProperty           = "symbols-count"
	LastUpdatedAtProperty          = "last-updated-at"
	LastReleasedAtProperty         = "last-released-at"
	AvailableFromProperty          = "available-from"
	PersonsIdsProperty             = "persons-ids"
	PersonsRolesProperty           = "persons-roles"
	PersonFullNameProperty         = "person-full-name"
	PersonUrlProperty              = "person-url"
	RatedAvgProperty               = "rated-avg"
	RatedTotalCountProperty        = "rated-total-count"
	LinkedArtsIdsProperty          = "linked-arts-ids"
	SeriesIdProperty               = "series-ids"
	SeriesArtOrderProperty         = "series-art-order"
	SeriesArtsCountProperty        = "series-arts-count"
	SeriesNameProperty             = "series-name"
	SeriesUrlProperty              = "series-url"
	DateWrittenAtProperty          = "date-written-at"
	AlternativeVersionsProperty    = "alternative-versions"
	HTMLAnnotationProperty         = "html-annotation"
	HTMLAnnotationLitResProperty   = "html-annotation-litres"
	FirstTimeSaleAtProperty        = "first-time-sale-at"
	LiveLibRatedAvgProperty        = "livelib-rated-avg"
	LiveLibRatedTotalCountProperty = "livelib-rated-total-count"
	GenresIdsProperty              = "genres-ids"
	GenreNameProperty              = "genre-name"
	GenreUrlProperty               = "genre-url"
	TagsIdsProperty                = "tags-ids"
	TagNameProperty                = "tag-name"
	TagUrlProperty                 = "tag-url"
	ISBNProperty                   = "isbn"
	PublicationDateProperty        = "publication-date"
	YouTubeVideosProperty          = "youtube-videos"
	ContentsUrlProperty            = "contents-url"
	RegisteredAtProperty           = "registered-at"
	TranslatedAtProperty           = "translated-at"
	CurrentPagesOrSecondsProperty  = "current-pages-or-seconds"
	PublisherIdProperty            = "publisher-id"
	PublisherNameProperty          = "publisher-name"
	PublisherUrlProperty           = "publisher-url"
	RightholdersIdsProperty        = "rightholders-ids"
	RightholderNameProperty        = "rightholder-name"
	RightholderUrlProperty         = "rightholder-url"

	// arts files properties

	//legacy reduced from detail page
	LegacyAuthorsProperty           = "authors"
	LegacyCoauthorsProperty         = "coauthors"
	LegacyDescriptionProperty       = "description"
	LegacyDownloadLinksProperty     = "download-links"
	LegacyDownloadTitlesProperty    = "download-titles"
	LegacySequenceNameProperty      = "sequence-name"
	LegacySequenceNumberProperty    = "sequence-number"
	LegacyDateReleasedProperty      = "date-released"
	LegacyDateTranslatedProperty    = "date-translated"
	LegacyDateCreatedProperty       = "date-created"
	LegacyAgeRatingProperty         = "age-rating"
	LegacyVolumeProperty            = "volume"
	LegacyDurationProperty          = "duration"
	LegacyISBNProperty              = "isbn"
	LegacyTranslatorsProperty       = "translators"
	LegacyReadersProperty           = "readers"
	LegacyIllustratorsProperty      = "illustrators"
	LegacyCopyrightHoldersProperty  = "copyright-holders"
	LegacyComposersProperty         = "composers"
	LegacyAdapterProperty           = "adapter"
	LegacyPerformersProperty        = "performers"
	LegacyDirectorsProperty         = "directors"
	LegacySoundDirectorsProperty    = "sound-directors"
	LegacyPublishersProperty        = "publishers"
	LegacyTotalSizeProperty         = "total-size"
	LegacyTotalPagesProperty        = "total-pages"
	LegacyMissingDetailsIdsProperty = "missing-details-ids"
	LegacyBookTypeProperty          = "book-type"
	LegacyGenresProperty            = "genres"
	LegacyTagsProperty              = "tags"
	LegacyImageProperty             = "image"
	LegacyLanguageProperty          = "language"

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

func ArtsDetailsReduxProperties() []string {
	return []string{
		CoverUrlProperty,
		TitleProperty,
		ArtTypeProperty,
		PriceProperty,
		MinAgeProperty,
		SymbolsCountProperty,
		LastUpdatedAtProperty,
		LastReleasedAtProperty,
		AvailableFromProperty,
		PersonsIdsProperty,
		PersonsRolesProperty,
		PersonFullNameProperty,
		PersonUrlProperty,
		RatedAvgProperty,
		RatedTotalCountProperty,
		LinkedArtsIdsProperty,
		SeriesIdProperty,
		SeriesArtOrderProperty,
		SeriesArtsCountProperty,
		SeriesNameProperty,
		SeriesUrlProperty,
		DateWrittenAtProperty,
		AlternativeVersionsProperty,
		HTMLAnnotationProperty,
		HTMLAnnotationLitResProperty,
		FirstTimeSaleAtProperty,
		LiveLibRatedAvgProperty,
		LiveLibRatedTotalCountProperty,
		GenresIdsProperty,
		GenreNameProperty,
		GenreUrlProperty,
		TagsIdsProperty,
		TagNameProperty,
		TagUrlProperty,
		ISBNProperty,
		PublicationDateProperty,
		YouTubeVideosProperty,
		ContentsUrlProperty,
		RegisteredAtProperty,
		TranslatedAtProperty,
		CurrentPagesOrSecondsProperty,
		PublisherIdProperty,
		PublisherNameProperty,
		PublisherUrlProperty,
		RightholdersIdsProperty,
		RightholderNameProperty,
		RightholderUrlProperty,
	}
}

func LegacyReduxProperties() []string {
	return []string{
		LegacyAuthorsProperty,
		LegacyCoauthorsProperty,
		LegacyDescriptionProperty,
		LegacyDownloadLinksProperty,
		LegacyDownloadTitlesProperty,
		LegacySequenceNameProperty,
		LegacySequenceNumberProperty,
		LegacyDateReleasedProperty,
		LegacyDateTranslatedProperty,
		LegacyDateCreatedProperty,
		LegacyAgeRatingProperty,
		LegacyVolumeProperty,
		LegacyDurationProperty,
		LegacyISBNProperty,
		LegacyTranslatorsProperty,
		LegacyReadersProperty,
		LegacyIllustratorsProperty,
		LegacyCopyrightHoldersProperty,
		LegacyComposersProperty,
		LegacyAdapterProperty,
		LegacyPerformersProperty,
		LegacyDirectorsProperty,
		LegacySoundDirectorsProperty,
		LegacyPublishersProperty,
		LegacyTotalSizeProperty,
		LegacyTotalPagesProperty,
		LegacyMissingDetailsIdsProperty,
		LegacyBookTypeProperty,
		LegacyGenresProperty,
		LegacyTagsProperty,
		LegacyImageProperty,
		LegacyLanguageProperty,
	}
}

func ReduxProperties() []string {
	return append(ArtsDetailsReduxProperties(),
		[]string{
			ArtsHistoryOrderProperty,
			ArtsHistoryEventTimeProperty,
			//
			SyncCompletedProperty,
			ImportedProperty,
			DataSourceProperty,
			DehydratedListImageProperty,
			DehydratedItemImageProperty,
		}...)
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
