package compton_fragments

import (
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/direction"
)

func Contents(r compton.Registrar, contents *litres_integration.Contents) compton.Element {

	stack := compton.FlexItems(r, direction.Column)

	list := compton.Ul()

	for _, tocItem := range contents.TocItem {
		list.Append(compton.ListItemText(tocItem.Text))
	}

	stack.Append(list)

	return stack
}
