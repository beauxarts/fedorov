package compton_fragments

import (
	"github.com/beauxarts/fedorov/rest/compton_styles"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/redux"
	"strconv"
)

const dehydratedCount = 10

func BooksList(r compton.Registrar, ids []string, from, to int, rdx redux.Readable) compton.Element {

	r.RegisterStyles(compton_styles.Styles, "books-list.css")

	productCards := compton.FlexItems(r, direction.Row).JustifyContent(align.Center).Width(size.FullWidth)
	productCards.AddClass("books-list")

	if (to - from) < 10 {
		productCards.AddClass("items-" + strconv.Itoa(to-from))
	}

	for ii := from; ii < to; ii++ {
		id := ids[ii]
		productLink := compton.A("/book?id=" + id)

		productCard := BookCard(r, id, ii-from < dehydratedCount, rdx)
		productLink.Append(productCard)
		productCards.Append(productLink)
	}

	return productCards
}
