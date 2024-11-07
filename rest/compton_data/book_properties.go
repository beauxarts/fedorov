package compton_data

import "github.com/beauxarts/fedorov/data"

var BookProperties = []string{
	data.AuthorsProperty,
	data.CurrentPagesOrSecondsProperty,
	data.SeriesProperty,
	data.ReadersProperty,
	data.IllustratorsProperty,
	data.TranslatorsProperty,
	data.PerformersProperty,
	data.PaintersProperty,
	data.PublishersProperty,
	data.GenresProperty,
	data.TagsProperty,
	data.RightholdersProperty,
	data.MinAgeProperty,
	data.RatedAvgProperty,
	data.LivelibRatedAvgProperty,
	data.ISBNProperty,
	data.DateWrittenAtProperty,
	data.PublicationDateProperty,
	data.TranslatedAtProperty,
	data.RegisteredAtProperty,
	data.AvailableFromProperty,
	data.FirstTimeSaleAtProperty,
	data.LastReleasedAtProperty,
	data.LastUpdatedAtProperty,
}

var BookExternalLinksProperties = []string{
	data.LitresBookLinksProperty,
	data.LitresAuthorLinksProperty,
	data.LitresPublishersLinksProperty,
	data.LitresRightholdersLinksProperty,
	data.LitresSeriesLinksProperty,
	data.LitresGenresLinksProperty,
	data.LitresTagsLinksProperty,
}
