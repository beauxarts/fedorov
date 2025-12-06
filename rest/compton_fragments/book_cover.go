package compton_fragments

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/issa"
	"github.com/boggydigital/redux"
	"strconv"
)

func BookCover(r compton.Registrar, id string, rdx redux.Readable) compton.Element {

	imgSrc := "/book_cover?id=" + id
	var cover compton.Element
	if dehydSrc, sure := rdx.GetLastVal(data.DehydratedItemImageProperty, id); sure {
		hydSrc := issa.HydrateColor(dehydSrc)
		repColor, _ := rdx.GetLastVal(data.RepItemImageColorProperty, id)

		issaImg := compton.IssaImageHydrated(r, repColor, hydSrc, imgSrc)
		if ar, ok := rdx.GetLastVal(data.CoverAspectRatioProperty, id); ok {
			if arf, err := strconv.ParseFloat(ar, 64); err == nil {
				issaImg.AspectRatio(arf)
			}
		}

		cover = issaImg
	} else {
		cover = compton.Img(imgSrc)
	}

	cover.AddClass("book-cover")
	//cover.SetAttribute("style", "view-transition-name:product-image-"+id)

	return cover
}
