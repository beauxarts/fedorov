package compton_pages

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/kevlar"
)

func Annotation(id string, rdx kevlar.ReadableRedux) compton.PageElement {

	s := compton_fragments.ProductSection(compton_data.AnnotationSection)

	if annotation, ok := rdx.GetLastVal(data.HTMLAnnotationProperty, id); ok {

		s.Append(compton.PreText(annotation))

	}

	return s
}
