package cli

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/redux"
)

func ReduceLitResArtsDetailsHandler(_ *url.URL) error {

	//TODO: add vangogh style since-hours-ago
	return ReduceLitResArtsDetails()
}

func ReduceLitResArtsDetails() error {

	rlaa := nod.NewProgress("reducing arts details...")
	defer rlaa.Done()

	atr, err := data.NewArtsReader(litres_integration.ArtsTypeDetails)
	if err != nil {
		return err
	}

	rlaa.TotalInt(atr.Len())

	propertyIdValues := make(map[string]map[string][]string)
	for _, p := range data.ArtsDetailsProperties() {
		propertyIdValues[p] = make(map[string][]string)
	}

	for id := range atr.Keys() {

		ad, err := atr.ArtsDetails(id)
		if err != nil {
			return err
		}

		for _, p := range data.ArtsDetailsProperties() {
			propertyIdValues[p][id] = getArtsDetailsPropertyValues(ad, p)
		}

		for p, iv := range getDetailedPropertyValues(ad) {
			for i, v := range iv {
				propertyIdValues[p][i] = v
			}
		}

		rlaa.Increment()
	}

	wra := nod.NewProgress("writing redux values...")
	defer wra.Done()

	reduxDir := data.Pwd.AbsRelDirPath(data.Redux, data.Metadata)

	rdx, err := redux.NewWriter(reduxDir, data.ReduxProperties()...)
	if err != nil {
		return err
	}

	wra.TotalInt(len(propertyIdValues))

	for p, idValues := range propertyIdValues {

		if err := rdx.BatchReplaceValues(p, idValues); err != nil {
			return err
		}

		wra.Increment()
	}

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
		if person.Id > 0 {
			pid := fmtInt(person.Id)
			pkv[data.PersonFullNameProperty][pid] = []string{person.FullName}
			pkv[data.PersonUrlProperty][pid] = []string{person.Url}
		}
	}

	for _, series := range ad.Payload.Data.Series {
		if series.Id > 0 {
			sid := fmtInt(series.Id)
			pkv[data.SeriesNameProperty][sid] = []string{series.Name}
			pkv[data.SeriesUrlProperty][sid] = []string{series.Url}
			pkv[data.SeriesArtsCountProperty][sid] = []string{fmtInt(series.ArtsCount)}
		}
	}

	for _, genre := range ad.Payload.Data.Genres {
		if genre.Id > 0 {
			gid := fmtInt(genre.Id)
			pkv[data.GenreNameProperty][gid] = []string{genre.Name}
			pkv[data.GenreUrlProperty][gid] = []string{genre.Url}
		}
	}

	for _, tag := range ad.Payload.Data.Tags {
		if tag.Id > 0 {
			tid := fmtInt(tag.Id)
			pkv[data.TagNameProperty][tid] = []string{tag.Name}
			pkv[data.TagUrlProperty][tid] = []string{tag.Url}
		}
	}

	if pub := ad.Payload.Data.Publisher; pub != nil {
		if pub.Id > 0 {
			pid := fmtInt(pub.Id)
			pkv[data.PublisherNameProperty][pid] = []string{pub.Name}
			pkv[data.PublisherUrlProperty][pid] = []string{pub.Url}
		}
	}

	for _, rh := range ad.Payload.Data.Rightholders {
		if rh.Id > 0 {
			rid := fmtInt(rh.Id)
			pkv[data.RightholderNameProperty][rid] = []string{rh.Name}
			pkv[data.RightholderUrlProperty][rid] = []string{rh.Url}
		}
	}

	return pkv
}

func getArtsDetailsPropertyValues(ad *litres_integration.ArtsDetails, property string) (values []string) {
	add, val := ad.Payload.Data, ""
	switch property {
	case data.TitleProperty:
		val = add.Title
	case data.SubtitleProperty:
		val = add.Subtitle
	case data.CoverUrlProperty:
		val = add.CoverUrl
	case data.CoverAspectRatioProperty:
		val = fmtFloat(add.CoverRatio)
	case data.ArtTypeProperty:
		val = fmtInt(add.ArtType)
	case data.PriceProperty:
		val = fmtFloat(add.Prices.FinalPrice)
	case data.MinAgeProperty:
		val = fmtInt(add.MinAge)
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
	case data.LivelibRatedAvgProperty:
		val = fmtFloat(add.LivelibRatedAvg)
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
	case data.LitresLabelsProperty:
		if add.Labels.IsNew {
			values = append(values, compton_data.LitresLabelNew)
		}
		if add.Labels.IsBestseller {
			values = append(values, compton_data.LitresLabelBestseller)
		}
		if add.Labels.IsLitresExclusive {
			values = append(values, compton_data.LitresLabelExclusive)
		}
		if add.Labels.IsSalesHit {
			values = append(values, compton_data.LitresLabelSalesHit)
		}
	case data.ArtFourthPresentProperty:
		fmt.Println("")
		// do nothing
	default:
		panic("unknown property " + property)
	}
	if val != "" {
		values = []string{val}
	}
	return values
}
