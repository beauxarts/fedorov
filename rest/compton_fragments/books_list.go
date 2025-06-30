package compton_fragments

import (
	"github.com/beauxarts/fedorov/rest/compton_styles"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/redux"
)

const dehydratedCount = 10

func BooksList(r compton.Registrar, ids []string, from, to int, rdx redux.Readable) compton.Element {

	r.RegisterStyles(compton_styles.Styles, "books-list.css")

	productCards := compton.FlexItems(r, direction.Row).JustifyContent(align.Center).Width(size.FullWidth)
	productCards.AddClass("books-list")

	for ii := from; ii < to; ii++ {
		id := ids[ii]
		productLink := compton.A("/book?id=" + id)

		if productCard := BookCard(r, id, ii-from < dehydratedCount, rdx); productCard != nil {
			productLink.Append(productCard)
			productCards.Append(productLink)
		}
	}

	return productCards
}
