package compton_fragments

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/kevlar"
)

func SetTint(id string, p compton.Element, rdx kevlar.ReadableRedux) {
	if repColor, ok := rdx.GetLastVal(data.RepItemImageColorProperty, id); ok {
		p.SetAttribute("style", "background-color:color-mix(in display-p3,"+repColor+" var(--cma),var(--c-background))")
	}
}
