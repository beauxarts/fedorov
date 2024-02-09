package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/nod"
	"net/url"
	"strconv"
)

func ReduceLitResArtsDetailsHandler(u *url.URL) error {

	//TODO: add vangogh style since-hours-ago
	return ReduceLitResArtsDetails(0)
}

func ReduceLitResArtsDetails(since int64) error {

	rlaa := nod.NewProgress("reducing arts details...")
	defer rlaa.End()

	atr, err := data.NewArtsReader(litres_integration.ArtsTypeDetails)
	if err != nil {
		return rlaa.EndWithError(err)
	}

	var ids []string
	if since > 0 {
		ids = atr.ModifiedAfter(since, false)
	} else {
		ids = atr.Keys()
	}

	rlaa.TotalInt(len(ids))

	propertyIdValues := make(map[string]map[string][]string)
	for _, p := range data.ArtsDetailsReduxProperties() {
		propertyIdValues[p] = make(map[string][]string)
	}

	for _, id := range ids {

		ad, err := atr.ArtsDetails(id)
		if err != nil {
			return rlaa.EndWithError(err)
		}

		for _, p := range data.ArtsDetailsReduxProperties() {
			propertyIdValues[p][id] = getArtsDetailsPropertyValues(ad, p)
		}

		rlaa.Increment()
	}

	rlaa.EndWithResult("done")

	return nil
}

func fmtFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}

func fmtInt(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func getArtsDetailsPropertyValues(ad *litres_integration.ArtsDetails, property string) (values []string) {
	add, val := ad.Payload.Data, ""
	switch property {
	case data.TitleProperty:
		val = add.Title
	case data.CoverUrlProperty:
		val = add.CoverUrl
	case data.ArtTypeProperty:
		val = fmtInt(add.ArtType)
	case data.PriceProperty:
		val = fmtFloat(add.Prices.FinalPrice)
	case data.MinAgeProperty:
		val = fmtInt(add.MinAge)
	case data.SymbolsCountProperty:
		val = fmtInt(add.SymbolsCount)
	case data.LastUpdatedAtProperty:
		val = add.LastUpdatedAt
	case data.LastReleasedAtProperty:
		val = add.LastReleasedAt
	case data.AvailableFromProperty:
		val = add.AvailableFrom
	case data.PersonsIdsProperty:
		values = make([]string, 0, len(add.Persons))
		for _, person := range add.Persons {
			values = append(values, fmtInt(person.Id))
		}
	case data.PersonsRolesProperty:
		values = make([]string, 0, len(add.Persons))
		for _, person := range add.Persons {
			values = append(values, person.Role)
		}
	case data.PersonFullNameProperty:
		// do nothing
	case data.PersonUrlProperty:
		// do nothing
	case data.RatedAvgProperty:
		val = fmtFloat(add.Rating.RatedAvg)
	case data.RatedTotalCountProperty:
		val = fmtInt(add.Rating.RatedTotalCount)
	case data.LinkedArtsIdsProperty:
		values = make([]string, 0, len(add.LinkedArts))
		for _, linkedArt := range add.LinkedArts {
			values = append(values, fmtInt(linkedArt.Id))
		}
	case data.SeriesIdProperty:
		values = make([]string, 0, len(add.Series))
		for _, series := range add.Series {
			values = append(values, fmtInt(series.Id))
		}
	case data.SeriesArtOrderProperty:
		values = make([]string, 0, len(add.Series))
		for _, series := range add.Series {
			values = append(values, fmtInt(series.ArtOrder))
		}
	case data.SeriesArtsCountProperty:
		// do nothing
	case data.SeriesNameProperty:
		// do nothing
	case data.SeriesUrlProperty:
		// do nothing
	case data.DateWrittenAtProperty:
		val = add.DateWrittenAt
	case data.AlternativeVersionsProperty:
		val = fmtInt(add.AlternativeVersion.Id)
	case data.HTMLAnnotationProperty:
		val = add.HTMLAnnotation
	case data.HTMLAnnotationLitResProperty:
		val = add.HTMLAnnotationLitres
	case data.FirstTimeSaleAtProperty:
		val = add.FirstTimeSaleAt
	case data.LiveLibRatedAvgProperty:
		val = fmtFloat(add.LivelibRatedAvg)
	case data.LiveLibRatedTotalCountProperty:
		val = fmtInt(add.LivelibRatedCount)
	case data.GenresIdsProperty:
		values = make([]string, 0, len(add.Genres))
		for _, genre := range add.Genres {
			values = append(values, fmtInt(genre.Id))
		}
	case data.GenreNameProperty:
		// do nothing
	case data.GenreUrlProperty:
		// do nothing
	case data.TagsIdsProperty:
		values = make([]string, 0, len(add.Tags))
		for _, tag := range add.Tags {
			values = append(values, fmtInt(tag.Id))
		}
	case data.TagNameProperty:
		// do nothing
	case data.TagUrlProperty:
		// do nothing
	case data.ISBNProperty:
		val = add.ISBN
	case data.PublicationDateProperty:
		val = add.PublicationDate
	case data.YouTubeVideosProperty:
		values = make([]string, 0, len(add.YoutubeVideos))
		for _, ytVideo := range add.YoutubeVideos {
			values = append(values, ytVideo.Url)
		}
	case data.ContentsUrlProperty:
		val = add.ContentsUrl
	case data.RegisteredAtProperty:
		val = add.AdditionalInfo.RegisteredAt
	case data.TranslatedAtProperty:
		val = add.AdditionalInfo.TranslatedAt
	case data.CurrentPagesOrSecondsProperty:
		val = fmtInt(add.AdditionalInfo.CurrentPagesOrSeconds)
	case data.PublisherIdProperty:
		if pub := add.Publisher; pub != nil {
			val = fmtInt(add.Publisher.Id)
		}
	case data.PublisherNameProperty:
		// do nothing
	case data.PublisherUrlProperty:
	// do nothing
	case data.RightholdersIdsProperty:
		values = make([]string, 0, len(add.Rightholders))
		for _, rightholder := range add.Rightholders {
			values = append(values, fmtInt(rightholder.Id))
		}
	case data.RightholderNameProperty:
		// do nothing
	case data.RightholderUrlProperty:
		// do nothing
	default:
		panic("unknown property " + property)
	}
	if val != "" {
		values = []string{val}
	}
	return values
}
