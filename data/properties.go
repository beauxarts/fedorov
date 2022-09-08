package data

const (
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

	// sorting
	SortProperty       = "sort"
	DescendingProperty = "desc"
	//GenresProperty         = "genres"
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
	}
}
