package data

const (
	//reduced from the list
	MyBooksIdsProperty = "my-books-ids"
	HrefProperty       = "href"
	// reduced from detail page
	TitleProperty            = "title"
	AuthorsProperty          = "authors"
	CoauthorsProperty        = "coauthors"
	DownloadLinksProperty    = "download-links"
	SequenceNameProperty     = "sequence-name"
	SequenceNumberProperty   = "sequence-number"
	DateReleasedProperty     = "date-released"
	DateTranslatedProperty   = "date-translated"
	DateCreatedProperty      = "date-created"
	AgeRatingProperty        = "age-rating"
	VolumeProperty           = "volume"
	DurationProperty         = "duration"
	ISBNPropertyProperty     = "isbn"
	TranslatorsProperty      = "translators"
	ReadersProperty          = "readers"
	IllustratorsProperty     = "illustrators"
	CopyrightHoldersProperty = "copyright-holders"
	ComposersProperty        = "composers"
	AdapterProperty          = "adapter"
	PerformersProperty       = "performers"
	DirectorsProperty        = "directors"
	SoundDirectorsProperty   = "sound-directors"
	PublishersProperty       = "publishers"
	TotalSizeProperty        = "total-size"
	TotalPagesProperty       = "total-pages"

	//GenresProperty         = "genres"
)

func ReduxProperties() []string {
	return []string{
		MyBooksIdsProperty,
		HrefProperty,
		TitleProperty,
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
	}
}
