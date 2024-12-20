package data

const (
	IdProperty = "id"

	// operations
	ArtsOperationsOrderProperty     = "arts-operations-order"
	ArtsOperationsEventTimeProperty = "arts-operations-event-time"

	// arts details properties
	CoverUrlProperty              = "cover-url"
	CoverAspectRatioProperty      = "cover-aspect-ratio"
	TitleProperty                 = "title"
	SubtitleProperty              = "subtitle"
	ArtTypeProperty               = "art-type"
	ArtFourthPresentProperty      = "art-fourth-present"
	PriceProperty                 = "price"
	MinAgeProperty                = "min-age"
	LastUpdatedAtProperty         = "last-updated-at"
	LastReleasedAtProperty        = "last-released-at"
	AvailableFromProperty         = "available-from"
	PersonsIdsProperty            = "persons-ids"
	PersonsRolesProperty          = "persons-roles"
	PersonFullNameProperty        = "person-full-name"
	PersonUrlProperty             = "person-url"
	RatedAvgProperty              = "rated-avg"
	RatedTotalCountProperty       = "rated-total-count"
	LivelibRatedAvgProperty       = "livelib-rated-avg"
	LitresLabelsProperty          = "litres-labels"
	LinkedArtsIdsProperty         = "linked-arts-ids"
	SeriesIdProperty              = "series-ids"
	SeriesArtOrderProperty        = "series-art-order"
	SeriesArtsCountProperty       = "series-arts-count"
	SeriesNameProperty            = "series-name"
	SeriesUrlProperty             = "series-url"
	SeriesProperty                = "series"
	DateWrittenAtProperty         = "date-written-at"
	AlternativeVersionsProperty   = "alternative-versions"
	HTMLAnnotationProperty        = "html-annotation"
	HTMLAnnotationLitResProperty  = "html-annotation-litres"
	FirstTimeSaleAtProperty       = "first-time-sale-at"
	GenresIdsProperty             = "genres-ids"
	GenreNameProperty             = "genre-name"
	GenreUrlProperty              = "genre-url"
	GenresProperty                = "genres"
	TagsIdsProperty               = "tags-ids"
	TagNameProperty               = "tag-name"
	TagUrlProperty                = "tag-url"
	TagsProperty                  = "tags"
	ISBNProperty                  = "isbn"
	PublicationDateProperty       = "publication-date"
	YouTubeVideosProperty         = "youtube-videos"
	VideoTitleProperty            = "video-title"
	VideoDurationProperty         = "video-duration"
	VideoErrorProperty            = "video-error"
	ContentsUrlProperty           = "contents-url"
	RegisteredAtProperty          = "registered-at"
	TranslatedAtProperty          = "translated-at"
	CurrentPagesOrSecondsProperty = "current-pages-or-seconds"
	PublisherIdProperty           = "publisher-id"
	PublisherNameProperty         = "publisher-name"
	PublisherUrlProperty          = "publisher-url"
	PublishersProperty            = "publishers"
	RightholdersIdsProperty       = "rightholders-ids"
	RightholderNameProperty       = "rightholder-name"
	RightholderUrlProperty        = "rightholder-url"
	RightholdersProperty          = "rightholders"

	// persons roles
	AuthorsProperty      = "authors"
	IllustratorsProperty = "illustrators"
	PaintersProperty     = "painters"
	PerformersProperty   = "performers"
	ReadersProperty      = "readers"
	TranslatorsProperty  = "translators"

	// arts files properties

	// local properties
	//LocalTagsProperty     = "local-tags"
	//BookCompletedProperty = "book-completed"

	// sorting
	SortProperty       = "sort"
	DescendingProperty = "desc"
	// sync events
	SyncCompletedProperty = "sync-completed"
	// dehydrated images
	DehydratedListImageProperty         = "dehydrated-list-image"
	DehydratedListImageModifiedProperty = "dehydrated-list-image-modified"
	DehydratedItemImageProperty         = "dehydrated-item-image"
	DehydratedItemImageModifiedProperty = "dehydrated-item-image-modified"
	RepItemImageColorProperty           = "rep-item-image-color"
	RepListImageColorProperty           = "rep-list-image-color"

	// Litres links
	LitresBookLinksProperty         = "litres-book-links"
	LitresAuthorLinksProperty       = "litres-author-links"
	LitresSeriesLinksProperty       = "litres-series-links"
	LitresPublishersLinksProperty   = "litres-publishers-links"
	LitresRightholdersLinksProperty = "litres-rightholders-links"
	LitresGenresLinksProperty       = "litres-genres-links"
	LitresTagsLinksProperty         = "litres-tags-links"
)

func ArtsDetailsProperties() []string {
	return []string{
		CoverUrlProperty,
		CoverAspectRatioProperty,
		TitleProperty,
		SubtitleProperty,
		ArtTypeProperty,
		ArtFourthPresentProperty,
		PriceProperty,
		MinAgeProperty,
		LastUpdatedAtProperty,
		LastReleasedAtProperty,
		AvailableFromProperty,
		PersonsIdsProperty,
		PersonsRolesProperty,
		PersonFullNameProperty,
		PersonUrlProperty,
		RatedAvgProperty,
		LivelibRatedAvgProperty,
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
		GenresIdsProperty,
		GenreNameProperty,
		GenreUrlProperty,
		TagsIdsProperty,
		TagNameProperty,
		TagUrlProperty,
		YouTubeVideosProperty,
		ISBNProperty,
		PublicationDateProperty,
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

func PersonsRolesProperties() []string {
	return []string{
		AuthorsProperty,
		IllustratorsProperty,
		PaintersProperty,
		PerformersProperty,
		ReadersProperty,
		TranslatorsProperty,
	}
}

func IdNameProperties() []string {
	return []string{
		GenresProperty,
		TagsProperty,
		PublishersProperty,
		RightholdersProperty,
		SeriesProperty,
	}
}

func VideoProperties() []string {
	return []string{
		VideoTitleProperty,
		VideoErrorProperty,
		VideoDurationProperty,
	}
}

func DehydratedProperties() []string {
	return []string{
		DehydratedItemImageProperty,
		DehydratedItemImageModifiedProperty,
		RepItemImageColorProperty,
		DehydratedListImageProperty,
		DehydratedListImageModifiedProperty,
		RepListImageColorProperty,
	}
}

func ReduxProperties() []string {
	properties := ArtsDetailsProperties()
	properties = append(properties, PersonsRolesProperties()...)
	properties = append(properties, IdNameProperties()...)
	properties = append(properties, VideoProperties()...)
	properties = append(properties, DehydratedProperties()...)
	properties = append(properties, []string{
		ArtsOperationsOrderProperty,
		ArtsOperationsEventTimeProperty,
		//
		SyncCompletedProperty,
	}...)
	return properties
}
