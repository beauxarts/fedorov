package compton_fragments

import (
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/font_weight"
)

func Contents(r compton.Registrar, contents *litres_integration.Contents) compton.Element {

	stack := compton.FlexItems(r, direction.Column)

	if contents == nil {
		stack.Append(compton.Fspan(r, "Для данной книги оглавление не найдено").
			ForegroundColor(color.Gray).
			TextAlign(align.Center))
		return stack
	}

	list := compton.Ul()

	for _, tocItem := range contents.TocItem {
		li := compton.Li()
		li.AddClass("deep" + tocItem.Deep)
		tocText := compton.Fspan(r, tocItem.Text)
		if tocItem.Deep == "0" {
			tocText.FontWeight(font_weight.Bolder)
		}
		li.Append(tocText)
		list.Append(li)
	}

	stack.Append(list)

	return stack
}
