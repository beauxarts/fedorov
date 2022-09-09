package view_models

import "github.com/beauxarts/fedorov/data"

var SearchProperties = []string{
	data.AnyTextProperty,
	data.TitleProperty,
	data.BookTypeProperty,
	data.AuthorsProperty,
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
	data.ISBNPropertyProperty,
	data.SortProperty,
	data.DescendingProperty,
}
