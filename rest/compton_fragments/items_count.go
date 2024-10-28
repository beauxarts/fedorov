package compton_fragments

import (
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/size"
	"strconv"
	"strings"
)

const (
	singleItem          = "1 книга"
	manyItemsSinglePage = "{total} книг"
	manyItemsManyPages  = "{from}-{to} из {total} книг"
)

func ItemsCount(r compton.Registrar, from, to, total int) compton.Element {
	title := ""
	switch total {
	case 1:
		title = singleItem
	case to:
		title = strings.Replace(manyItemsSinglePage, "{total}", strconv.Itoa(total), 1)
	default:
		title = strings.Replace(manyItemsManyPages, "{from}", strconv.Itoa(from+1), 1)
		title = strings.Replace(title, "{to}", strconv.Itoa(to), 1)
		title = strings.Replace(title, "{total}", strconv.Itoa(total), 1)
	}

	itemsCount := compton.Fspan(r, title).
		ForegroundColor(color.Gray).
		FontSize(size.XSmall).
		FontWeight(font_weight.Normal)

	return itemsCount
}
