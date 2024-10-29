package compton_fragments

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/issa"
	"github.com/boggydigital/kevlar"
)

func BookCover(r compton.Registrar, id string, rdx kevlar.ReadableRedux) compton.Element {

	imgSrc := "/book_cover?id=" + id
	var cover compton.Element
	if dehydSrc, sure := rdx.GetLastVal(data.DehydratedItemImageProperty, id); sure {
		hydSrc := issa.HydrateColor(dehydSrc)
		cover = compton.IssaImageHydrated(r, hydSrc, imgSrc)
	} else {
		cover = compton.Img(imgSrc)
	}

	cover.AddClass("book-cover")
	return cover
}
