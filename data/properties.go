package data

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
