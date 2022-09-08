package view_models

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/litres_integration"
	"github.com/boggydigital/kvas"
)

var detailsPropertyOrder = []string{
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
	data.SequenceNameProperty,
	//data.SequenceNumberProperty,
	data.DateReleasedProperty,
	data.DateTranslatedProperty,
	data.DateCreatedProperty,
	data.AgeRatingProperty,
	data.DurationProperty,
	data.VolumeProperty,
	data.ISBNPropertyProperty,
	data.TotalSizeProperty,
	data.TotalPagesProperty,
	data.HrefProperty,
}

type Book struct {
	Id string
	// Title
	Title string
	Type  string
	// Text properties
	Properties      map[string]map[string]string
	PropertyOrder   []string
	PropertyTitles  map[string]string
	PropertyClasses map[string]string
	// Sections
	Sections      []string
	SectionTitles map[string]string
}

func NewBook(id string, rxa kvas.ReduxAssets) *Book {

	bvm := &Book{
		Id:             id,
		Properties:     make(map[string]map[string]string),
		PropertyOrder:  detailsPropertyOrder,
		PropertyTitles: propertyTitles,
		Sections:       []string{DescriptionSection, DownloadsSection},
		SectionTitles:  sectionTitles,
	}

	bvm.Title, _ = rxa.GetFirstVal(data.TitleProperty, id)
	bvm.Type, _ = rxa.GetFirstVal(data.BookTypeProperty, id)

	for _, p := range detailsPropertyOrder {
		// sequence name needs to be formatted in a special way, see below
		if p == data.SequenceNameProperty {
			continue
		}
		bvm.Properties[p] = getPropertyLinks(id, p, rxa)
	}

	seqNames, _ := rxa.GetAllUnchangedValues(data.SequenceNameProperty, id)
	seqNumbers, _ := rxa.GetAllUnchangedValues(data.SequenceNumberProperty, id)

	if len(seqNames) == len(seqNumbers) {
		// to format sequence name we need to append sequence relevant number
		// and use the regular link href that points to the series overall
		bvm.Properties[data.SequenceNameProperty] = make(map[string]string)
		for i, name := range seqNames {
			nameNumber := fmt.Sprintf("%s %s", name, seqNumbers[i])
			bvm.Properties[data.SequenceNameProperty][nameNumber] =
				formatPropertyLinkHref(data.SequenceNameProperty, name)
		}
	}

	return bvm
}

func getPropertyLinks(id string, property string, rxa kvas.ReduxAssets) map[string]string {

	propertyLinks := make(map[string]string)

	values, _ := rxa.GetAllUnchangedValues(property, id)

	for _, value := range values {

		linkTitle := formatPropertyLinkTitle(property, value)
		propertyLinks[linkTitle] = formatPropertyLinkHref(property, value)
	}

	return propertyLinks
}

func formatPropertyLinkTitle(property, link string) string {
	title := link

	switch property {
	case data.HrefProperty:
		title = "ЛитРес"
	default:
		// do nothing
	}

	return title
}

func formatPropertyLinkHref(property, link string) string {
	switch property {
	case data.HrefProperty:
		return litres_integration.HrefUrl(link).String()
	case data.AgeRatingProperty:
		fallthrough
	case data.ISBNPropertyProperty:
		fallthrough
	case data.TotalPagesProperty:
		fallthrough
	case data.TotalSizeProperty:
		fallthrough
	case data.VolumeProperty:
		fallthrough
	case data.DurationProperty:
		return ""
	default:
		// do nothing
	}
	return fmt.Sprintf("/search?%s=%s", property, link)
}
