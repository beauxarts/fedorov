package compton_fragments

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/redux"
	"maps"
	"slices"
	"strings"
)

type formattedProperty struct {
	values  map[string]string
	class   string
	actions map[string]string
}

func Information(r compton.Registrar, id string, rdx redux.Readable) compton.Element {
	grid := compton.FlexItems(r, direction.Row).JustifyContent(align.Start).RowGap(size.Normal)

	artType := litres_integration.ArtTypeText
	if ats, ok := rdx.GetLastVal(data.ArtTypeProperty, id); ok {
		artType = litres_integration.ParseArtType(ats)
	}

	for _, property := range compton_data.BookProperties {

		fmtProperty := formatProperty(id, property, artType, rdx)
		if tv := propertyTitleValues(r, property, fmtProperty); tv != nil {
			grid.Append(tv)
		}
	}

	return grid
}

func formatProperty(id, property string, at litres_integration.ArtType, rdx redux.Readable) formattedProperty {

	fmtProperty := formattedProperty{
		actions: make(map[string]string),
		values:  make(map[string]string),
	}

	var seriesNames []string
	var seriesArtOrder []string
	//var seriesArtCounts []string
	if property == data.SeriesProperty {
		seriesArtOrder, _ = rdx.GetAllValues(data.SeriesArtOrderProperty, id)
		seriesIds, _ := rdx.GetAllValues(data.SeriesIdProperty, id)
		for _, seriesId := range seriesIds {
			seriesName, _ := rdx.GetLastVal(data.SeriesNameProperty, seriesId)
			seriesNames = append(seriesNames, seriesName)
		}
	}

	values, _ := rdx.GetAllValues(property, id)

	for ii, value := range values {
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
		case data.ISBNProperty:
			fallthrough
		case data.LastUpdatedAtProperty:
			value, _, _ = strings.Cut(value, "T")
			fmtProperty.values[value] = noHref()
		case data.LivelibRatedAvgProperty:
			if value != "0.00" && value != "0.0" {
				fmtProperty.values[value] = noHref()
			}
		case data.RatedAvgProperty:
			fallthrough
		case data.PriceProperty:
			fmtProperty.values[value] = noHref()
		case data.CurrentPagesOrSecondsProperty:
			value = fmtCurrentPagesOrSeconds(value, at)
			fmtProperty.values[value] = noHref()
		case data.SeriesProperty:
			if ii < len(seriesNames) {
				value = seriesNames[ii]
				if ii < len(seriesArtOrder) {
					if seriesArtOrder[ii] != "0" {
						value += " " + seriesArtOrder[ii]
					}
				}
				fmtProperty.values[value] = searchHref(data.SeriesProperty, seriesNames[ii])
			}
		default:
			fmtProperty.values[value] = searchHref(property, value)
		}
	}

	// format actions, class
	switch property {
	}

	return fmtProperty
}

func searchHref(property, value string) string {
	return fmt.Sprintf("/search?%s=%s&sort=date-written-at&desc=true", property, value)
}

func propertyTitleValues(r compton.Registrar, property string, fmtProperty formattedProperty) *compton.TitleValuesElement {

	if len(fmtProperty.values) == 0 && len(fmtProperty.actions) == 0 {
		return nil
	}

	tv := compton.TitleValues(r, compton_data.PropertyTitles[property]).
		SetLinksTarget(compton.LinkTargetTop).
		ForegroundColor(color.RepForeground).
		TitleForegroundColor(color.RepGray).
		RowGap(size.XSmall)

	if len(fmtProperty.values) > 0 {

		if len(fmtProperty.values) < 2 {
			tv.AppendLinkValues(fmtProperty.values)
		} else {
			summaryTitle := fmt.Sprintf("%d штук(и)", len(fmtProperty.values))
			//	ForegroundColor(color.Foreground)
			ds := compton.DSSmall(r, summaryTitle, false).
				SummaryMarginBlockEnd(size.Normal).
				DetailsMarginBlockEnd(size.Small)
			row := compton.FlexItems(r, direction.Row).
				JustifyContent(align.Start)
			keys := maps.Keys(fmtProperty.values)
			sortedKeys := slices.Sorted(keys)
			for _, link := range sortedKeys {
				href := fmtProperty.values[link]
				anchor := compton.AText(link, href)
				anchor.SetAttribute("target", "_top")
				row.Append(anchor)
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
