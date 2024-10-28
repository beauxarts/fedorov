package compton_fragments

import (
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/boggydigital/compton"
)

func AppNavLinks(r compton.Registrar, current string) compton.Element {
	targets := compton.TextLinks(
		compton_data.AppNavLinks,
		current,
		compton_data.AppNavOrder...)
	compton.SetIcons(targets, compton_data.AppNavIcons)

	return compton.NavLinksTargets(r, targets...)
}
