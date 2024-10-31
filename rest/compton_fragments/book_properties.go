package compton_fragments

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/kevlar"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"strings"
)

type formattedProperty struct {
	values  map[string]string
	class   string
	actions map[string]string
}

func BookProperties(r compton.Registrar, id string, rdx kevlar.ReadableRedux) compton.Element {
	grid := compton.GridItems(r).JustifyContent(align.Center)

	for _, property := range compton_data.BookProperties {

		fmtProperty := formatProperty(id, property, rdx)
		if tv := propertyTitleValues(r, property, fmtProperty); tv != nil {
			grid.Append(tv)
		}
	}

	return grid
}

func formatProperty(id, property string, rdx kevlar.ReadableRedux) formattedProperty {

	fmtProperty := formattedProperty{
		actions: make(map[string]string),
		values:  make(map[string]string),
	}

	values, _ := rdx.GetAllValues(property, id)
	//firstValue := ""
	//if len(values) > 0 {
	//	firstValue = values[0]
	//}

	for _, value := range values {
		switch property {
		case data.DateWrittenAtProperty:
			fallthrough
		case data.PublicationDateProperty:
			fallthrough
		case data.TranslatedAtProperty:
			fallthrough
		case data.RegisteredAtProperty:
			fallthrough
		case data.AvailableFromProperty:
			fallthrough
		case data.FirstTimeSaleAtProperty:
			fallthrough
		case data.LastReleasedAtProperty:
			fallthrough
		case data.LastUpdatedAtProperty:
			value, _, _ = strings.Cut(value, "T")
			fmtProperty.values[value] = noHref()
		case data.RatedAvgProperty:
			fallthrough
		case data.PriceProperty:
			fmtProperty.values[value] = noHref()
		//case vangogh_local_data.WishlistedProperty:
		//	if owned {
		//		break
		//	}
		//	title := "No"
		//	if value == vangogh_local_data.TrueValue {
		//		title = "Yes"
		//	}
		//	fmtProperty.values[title] = searchHref(property, value)
		//case vangogh_local_data.IncludesGamesProperty:
		//	fallthrough
		//case vangogh_local_data.IsIncludedByGamesProperty:
		//	fallthrough
		//case vangogh_local_data.RequiresGamesProperty:
		//	fallthrough
		//case vangogh_local_data.IsRequiredByGamesProperty:
		//	refTitle := value
		//	if rtp, ok := rdx.GetLastVal(vangogh_local_data.TitleProperty, value); ok {
		//		refTitle = rtp
		//	}
		//	fmtProperty.values[refTitle] = "/product?id=" + value
		//case vangogh_local_data.GOGOrderDateProperty:
		//	jtd := justTheDate(value)
		//	fmtProperty.values[jtd] = searchHref(property, jtd)
		//case vangogh_local_data.LanguageCodeProperty:
		//	fmtProperty.values[compton_data.FormatLanguage(value)] = searchHref(property, value)
		//case vangogh_local_data.RatingProperty:
		//	fmtProperty.values[fmtGOGRating(value)] = noHref()
		//case vangogh_local_data.TagIdProperty:
		//	tagName := value
		//	if tnp, ok := rdx.GetLastVal(vangogh_local_data.TagNameProperty, value); ok {
		//		tagName = tnp
		//	}
		//	fmtProperty.values[tagName] = searchHref(property, tagName)
		//case vangogh_local_data.PriceProperty:
		//	if !isFree {
		//		if isDiscounted && !owned {
		//			if bpp, ok := rdx.GetLastVal(vangogh_local_data.BasePriceProperty, id); ok {
		//				fmtProperty.values["Base: "+bpp] = noHref()
		//			}
		//			fmtProperty.values["Sale: "+value] = noHref()
		//		} else {
		//			fmtProperty.values[value] = noHref()
		//		}
		//	}
		//case vangogh_local_data.HLTBHoursToCompleteMainProperty:
		//	fallthrough
		//case vangogh_local_data.HLTBHoursToCompletePlusProperty:
		//	fallthrough
		//case vangogh_local_data.HLTBHoursToComplete100Property:
		//	ct := strings.TrimLeft(value, "0") + " hrs"
		//	fmtProperty.values[ct] = noHref()
		//case vangogh_local_data.HLTBReviewScoreProperty:
		//	if value != "0" {
		//		fmtProperty.values[fmtHLTBRating(value)] = noHref()
		//	}
		//case vangogh_local_data.DiscountPercentageProperty:
		//	fmtProperty.values[value] = noHref()
		//case vangogh_local_data.PublishersProperty:
		//	fallthrough
		//case vangogh_local_data.DevelopersProperty:
		//	fmtProperty.values[value] = grdSortedSearchHref(property, value)
		//case vangogh_local_data.EnginesBuildsProperty:
		//	fmtProperty.values[value] = noHref()
		//
		default:
			fmtProperty.values[value] = searchHref(property, value)
		}
	}

	// format actions, class
	switch property {
	//case vangogh_local_data.OwnedProperty:
	//	if res, ok := rdx.GetLastVal(vangogh_local_data.ValidationResultProperty, id); ok {
	//		fmtProperty.class = res
	//	}
	//case vangogh_local_data.WishlistedProperty:
	//	if !owned {
	//		switch firstValue {
	//		case vangogh_local_data.TrueValue:
	//			fmtProperty.actions["Remove"] = "/wishlist/remove?id=" + id
	//		case vangogh_local_data.FalseValue:
	//			fmtProperty.actions["Add"] = "/wishlist/add?id=" + id
	//		}
	//	}
	//case vangogh_local_data.TagIdProperty:
	//	if owned {
	//		fmtProperty.actions["Edit"] = "/tags/edit?id=" + id
	//	}
	//case vangogh_local_data.LocalTagsProperty:
	//	fmtProperty.actions["Edit"] = "/local-tags/edit?id=" + id
	//case vangogh_local_data.SteamReviewScoreDescProperty:
	//	fmtProperty.class = reviewClass(firstValue)
	//case vangogh_local_data.RatingProperty:
	//	fmtProperty.class = reviewClass(fmtGOGRating(firstValue))
	//case vangogh_local_data.HLTBReviewScoreProperty:
	//	fmtProperty.class = reviewClass(fmtHLTBRating(firstValue))
	//case vangogh_local_data.SteamDeckAppCompatibilityCategoryProperty:
	//	fmtProperty.class = firstValue
	//	if firstValue != "" {
	//		fmtProperty.actions["&darr;"] = "#Steam Deck"
	//	}
	}

	return fmtProperty
}

func searchHref(property, value string) string {
	return fmt.Sprintf("/search?%s=%s", property, value)
}

func propertyTitleValues(r compton.Registrar, property string, fmtProperty formattedProperty) *compton.TitleValuesElement {

	if len(fmtProperty.values) == 0 && len(fmtProperty.actions) == 0 {
		return nil
	}

	tv := compton.TitleValues(r, compton_data.PropertyTitles[property]).
		ForegroundColor(color.Gray).TitleForegroundColor(color.Foreground)

	if len(fmtProperty.values) > 0 {

		if len(fmtProperty.values) < 4 {
			tv.AppendLinkValues(fmtProperty.values)
		} else {
			summaryTitle := fmt.Sprintf("%d values", len(fmtProperty.values))
			summaryElement := compton.Fspan(r, summaryTitle).
				ForegroundColor(color.Foreground)
			ds := compton.DSSmall(r, summaryElement, false).
				SummaryMarginBlockEnd(size.Normal).
				DetailsMarginBlockEnd(size.Small)
			row := compton.FlexItems(r, direction.Row).
				JustifyContent(align.Start)
			keys := maps.Keys(fmtProperty.values)
			slices.Sort(keys)
			for _, link := range keys {
				href := fmtProperty.values[link]
				row.Append(compton.AText(link, href))
			}
			ds.Append(row)
			tv.AppendValues(ds)
		}

		if fmtProperty.class != "" {
			tv.AddClass(fmtProperty.class)
		}
	}

	if len(fmtProperty.actions) > 0 {
		for ac, acHref := range fmtProperty.actions {
			actionLink := compton.A(acHref)
			actionLink.Append(compton.Fspan(r, ac).ForegroundColor(color.Blue))
			tv.AppendValues(actionLink)
		}
	}

	return tv
}

func noHref() string {
	return ""
}
