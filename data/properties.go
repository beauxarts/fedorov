package data

const (
	//reduced from the list
	MyBooksIdsProperty = "my-books-ids"
	HrefProperty       = "href"
	// reduced from detail page
	TitleProperty          = "title"
	AuthorsProperty        = "authors"
	TranslatorsProperty    = "translators"
	ReadersProperty        = "readers"
	GenresProperty         = "genres"
	PublisherProperty      = "publisher"
	ISBNPropertyProperty   = "isbn"
	YearProperty           = "year"
	SequenceNameProperty   = "sequence-name"
	SequenceNumberProperty = "sequence-number"
)

func ReduxProperties() []string {
	return []string{
		MyBooksIdsProperty,
		HrefProperty,
		TitleProperty,
		AuthorsProperty,
		TranslatorsProperty,
		ReadersProperty,
		GenresProperty,
		PublisherProperty,
		ISBNPropertyProperty,
		YearProperty,
		SequenceNameProperty,
		SequenceNumberProperty,
	}
}
