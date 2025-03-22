package compton_fragments

import (
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/redux"
	"strconv"
)

const dehydratedCount = 10

func BooksList(r compton.Registrar, ids []string, from, to int, rdx redux.Readable) compton.Element {
	productCards := compton.GridItems(r).JustifyContent(align.Center)

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
