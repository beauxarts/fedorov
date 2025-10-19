package compton_fragments

import (
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/boggydigital/compton"
)

func AppNavLinks(r compton.Registrar, current string) compton.Element {

	appNavLinks := compton.NavLinks(r)
	appNavLinks.SetAttribute("style", "view-transition-name:primary-nav")

	appNavLinks.AppendLink(r, &compton.NavTarget{
		Href:     "/latest",
		Title:    compton_data.AppNavLatest,
		Symbol:   compton.Bookmark,
		Selected: current == compton_data.AppNavLatest,
	})

	appNavLinks.AppendLink(r, &compton.NavTarget{
		Href:     "/search",
		Title:    compton_data.AppNavSearch,
		Symbol:   compton.Search,
		Selected: current == compton_data.AppNavSearch,
	})

	return appNavLinks
}
