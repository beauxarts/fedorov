package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pasu"
	"net/url"
	"strconv"
)

func ReduceLitResArtsDetailsHandler(u *url.URL) error {

	//TODO: add vangogh style since-hours-ago
	return ReduceLitResArtsDetails()
}

func ReduceLitResArtsDetails() error {

	rlaa := nod.NewProgress("reducing arts details...")
	defer rlaa.End()

	atr, err := data.NewArtsReader(litres_integration.ArtsTypeDetails)
	if err != nil {
		return rlaa.EndWithError(err)
	}

	ids := atr.Keys()

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

		for p, iv := range getDetailedPropertyValues(ad) {
			for i, v := range iv {
				propertyIdValues[p][i] = v
			}
		}

		rlaa.Increment()
	}

	rlaa.EndWithResult("done")

	wra := nod.NewProgress("writing redux values...")
	defer wra.End()

	reduxDir, err := pasu.GetAbsRelDir(data.Redux)
	if err != nil {
		return wra.EndWithError(err)
	}

	rdx, err := kvas.NewReduxWriter(reduxDir, data.ReduxProperties()...)
	if err != nil {
		return wra.EndWithError(err)
	}

	wra.TotalInt(len(propertyIdValues))

	for p, idValues := range propertyIdValues {

		if err := rdx.BatchReplaceValues(p, idValues); err != nil {
			return wra.EndWithError(err)
		}

		wra.Increment()
	}

	wra.EndWithResult("done")

	return nil
}

func fmtFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}

func fmtInt(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func getDetailedPropertyValues(ad *litres_integration.ArtsDetails) (pkv map[string]map[string][]string) {

	properties := []string{
		data.PersonFullNameProperty,
		data.PersonUrlProperty,
		data.SeriesNameProperty,
		data.SeriesArtsCountProperty,
		data.SeriesUrlProperty,
		data.GenreNameProperty,
		data.GenreUrlProperty,
		data.TagNameProperty,
		data.TagUrlProperty,
		data.PublisherNameProperty,
		data.PublisherUrlProperty,
		data.RightholderNameProperty,
		data.RightholderUrlProperty,
	}
	pkv = make(map[string]map[string][]string)
	for _, p := range properties {
		pkv[p] = make(map[string][]string)
	}

	for _, person := range ad.Payload.Data.Persons {
		pid := fmtInt(person.Id)
		pkv[data.PersonFullNameProperty][pid] = []string{person.FullName}
		pkv[data.PersonUrlProperty][pid] = []string{person.Url}
	}

	for _, series := range ad.Payload.Data.Series {
		sid := fmtInt(series.Id)
		pkv[data.SeriesNameProperty][sid] = []string{series.Name}
		pkv[data.SeriesUrlProperty][sid] = []string{series.Url}
		pkv[data.SeriesArtsCountProperty][sid] = []string{fmtInt(series.ArtsCount)}
	}

	for _, genre := range ad.Payload.Data.Genres {
		gid := fmtInt(genre.Id)
		pkv[data.GenreNameProperty][gid] = []string{genre.Name}
		pkv[data.GenreUrlProperty][gid] = []string{genre.Url}
	}

	for _, tag := range ad.Payload.Data.Tags {
		tid := fmtInt(tag.Id)
		pkv[data.GenreNameProperty][tid] = []string{tag.Name}
		pkv[data.GenreUrlProperty][tid] = []string{tag.Url}
	}

	if pub := ad.Payload.Data.Publisher; pub != nil {
		pid := fmtInt(pub.Id)
		pkv[data.PublisherNameProperty][pid] = []string{pub.Name}
		pkv[data.GenreUrlProperty][pid] = []string{pub.Url}
	}

	for _, rh := range ad.Payload.Data.Rightholders {
		rid := fmtInt(rh.Id)
		pkv[data.RightholderNameProperty][rid] = []string{rh.Name}
		pkv[data.RightholderUrlProperty][rid] = []string{rh.Url}
	}

	return pkv
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
