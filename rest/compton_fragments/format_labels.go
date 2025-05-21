package compton_fragments

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/redux"
)

var artTypeColors = map[litres_integration.ArtType]color.Color{
	litres_integration.ArtTypeText:  color.Foreground,
	litres_integration.ArtTypeAudio: color.Blue,
	litres_integration.ArtTypePDF:   color.Red,
}

var artTypeBadges = map[litres_integration.ArtType]string{
	litres_integration.ArtTypeText:  "ТЕКСТ",
	litres_integration.ArtTypeAudio: "АУДИО",
	litres_integration.ArtTypePDF:   "PDF",
}

func FormatBadges(id string, rdx redux.Readable) []compton.FormattedBadge {
	fmtBadges := make([]compton.FormattedBadge, 0)

	for _, p := range compton_data.LabelProperties {
		fmtBadges = append(fmtBadges, formatBadge(id, p, rdx)...)
	}

	return fmtBadges
}

func formatBadge(id, property string, rdx redux.Readable) []compton.FormattedBadge {

	badges := make([]compton.FormattedBadge, 0)

	val, _ := rdx.GetLastVal(property, id)
	values, _ := rdx.GetAllValues(property, id)

	switch property {
	case data.ArtTypeProperty:
		at := litres_integration.ParseArtType(val)
		badge := compton.FormattedBadge{
			Title: artTypeBadges[at],
			Color: artTypeColors[at],
			Class: property,
		}
		badges = append(badges, badge)
	case data.LitresLabelsProperty:
		for _, value := range values {
			label := compton.FormattedBadge{
				Title: value,
				Class: property,
			}
			badges = append(badges, label)
		}
	case data.ArtFourthPresentProperty:
		if val == "true" {
			label := compton.FormattedBadge{
				Title: compton_data.PropertyTitles[data.ArtFourthPresentProperty],
				Color: color.Purple,
				Class: property,
			}
			badges = append(badges, label)
		}
	}

	return badges
}
