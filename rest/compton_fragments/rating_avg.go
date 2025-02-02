package compton_fragments

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/redux"
)

func RatingAvg(r compton.Registrar, id string, rdx redux.Readable) compton.Element {
	if rp, ok := rdx.GetLastVal(data.RatedAvgProperty, id); ok {
		return compton.Fspan(r, "Рейтинг: "+rp).
			BackgroundColor(color.Background).
			ForegroundColor(color.Foreground).
			FontSize(size.Small).
			FontWeight(font_weight.Normal).
			PaddingInline(size.XSmall).
			PaddingBlock(size.XXSmall).
			BorderRadius(size.XXSmall)
	}
	return nil
}
