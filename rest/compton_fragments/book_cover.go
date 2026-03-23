package compton_fragments

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/redux"
)

func BookCover(id string, rdx redux.Readable) compton.Element {

	imgSrc := "/book_cover?id=" + id

	imgEager := compton.ImageEager(imgSrc)
	if ar, ok := rdx.GetLastVal(data.CoverAspectRatioProperty, id); ok {
		imgEager.SetAttribute("style", "aspect-ratio:"+ar)
	}

	imgEager.AddClass("book-cover")

	return imgEager
}
