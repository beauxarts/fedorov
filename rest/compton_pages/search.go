package compton_pages

import (
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/beauxarts/fedorov/rest/compton_styles"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/kevlar"
)

func Search(ids []string, from, to, total int, rdx kevlar.ReadableRedux) compton.PageElement {

	p := compton.Page(compton_data.AppNavSearch)
	p.RegisterStyles(compton_styles.Styles, "book-labels.css")

	stack := compton.FlexItems(p, direction.Column)
	p.Append(stack)

	appNav := compton_fragments.AppNavLinks(p, compton_data.AppNavSearch)
	stack.Append(appNav)

	return p
}
