package stencil_app

import "github.com/beauxarts/fedorov/data"

var BookProperties = []string{
	data.BookTypeProperty,

	data.BookCompletedProperty,
	data.LocalTagsProperty,

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
	data.CopyrightHoldersProperty,
	data.PublishersProperty,
	data.GenresProperty,
	data.TagsProperty,
	data.SequenceNameProperty,
	data.DateReleasedProperty,
	data.DateTranslatedProperty,
	data.DateCreatedProperty,
	data.AgeRatingProperty,
	data.DurationProperty,
	data.VolumeProperty,
	data.ISBNProperty,
	data.TotalSizeProperty,
	data.TotalPagesProperty,
	data.HrefProperty,
}

var BookLabels = []string{
	data.BookTypeProperty,
	data.BookCompletedProperty,
	data.LocalTagsProperty,
	data.ImportedProperty,
}

var BookHiddenProperties = []string{
	data.BookTypeProperty,
}
