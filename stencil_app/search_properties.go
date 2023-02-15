package stencil_app

import "github.com/beauxarts/fedorov/data"

var SearchProperties = []string{
	data.AnyTextProperty,
	data.TitleProperty,
	data.AuthorsProperty,
	data.BookTypeProperty,
	data.BookCompletedProperty,
	data.CoauthorsProperty,
	data.TranslatorsProperty,
	data.ReadersProperty,
	data.IllustratorsProperty,
	data.ComposersProperty,
	data.AdapterProperty,
	data.PerformersProperty,
	data.DirectorsProperty,
	data.SoundDirectorsProperty,
	data.GenresProperty,
	data.TagsProperty,
	data.DescriptionProperty,
	data.CopyrightHoldersProperty,
	data.PublishersProperty,
	data.SequenceNameProperty,
	data.DateReleasedProperty,
	data.DateTranslatedProperty,
	data.DateCreatedProperty,
	data.ISBNProperty,
	data.SortProperty,
	data.DescendingProperty,
	data.ImportedProperty,
}

var SearchHighlightProperties = []string{
	data.AnyTextProperty,
	data.TitleProperty,
	data.AuthorsProperty,
}
