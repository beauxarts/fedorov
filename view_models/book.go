package view_models

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/stencil_app"
	"github.com/boggydigital/kvas"
)

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
		PropertyOrder:  stencil_app.BooksProperties,
		PropertyTitles: stencil_app.PropertyTitles,
		Sections:       stencil_app.BookSections,
		SectionTitles:  stencil_app.SectionTitles,
	}

	bvm.Title, _ = rxa.GetFirstVal(data.TitleProperty, id)
	bvm.Type, _ = rxa.GetFirstVal(data.BookTypeProperty, id)

	for _, p := range bvm.PropertyOrder {
		// sequence name needs to be formatted in a special way, see below
		if p == data.SequenceNameProperty {
			continue
		}
		bvm.Properties[p] = getPropertyLinks(id, p, rxa)
	}

	//seqNames, _ := rxa.GetAllUnchangedValues(data.SequenceNameProperty, id)
	//seqNumbers, _ := rxa.GetAllUnchangedValues(data.SequenceNumberProperty, id)

	//if len(seqNames) == len(seqNumbers) {
	//	// to format sequence name we need to append sequence relevant number
	//	// and use the regular link href that points to the series overall
	//	bvm.Properties[data.SequenceNameProperty] = make(map[string]string)
	//	for i, name := range seqNames {
	//		nameNumber := fmt.Sprintf("%s %s", name, seqNumbers[i])
	//		bvm.Properties[data.SequenceNameProperty][nameNumber] =
	//			formatPropertyLinkHref(data.SequenceNameProperty, name)
	//	}
	//}

	return bvm
}

func getPropertyLinks(id string, property string, rxa kvas.ReduxAssets) map[string]string {

	propertyLinks := make(map[string]string)

	values, _ := rxa.GetAllUnchangedValues(property, id)

	for _, value := range values {

		linkTitle := value
		propertyLinks[linkTitle] = value
	}

	return propertyLinks
}
