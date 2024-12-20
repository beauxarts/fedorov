package compton_fragments

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/kevlar"
)

func FormatLabels(id string, rdx kevlar.ReadableRedux) []compton.FormattedLabel {
	fmtLabels := make([]compton.FormattedLabel, 0)

	for _, p := range compton_data.LabelProperties {
		fmtLabels = append(fmtLabels, formatLabel(id, p, rdx)...)
	}

	return fmtLabels
}

func formatLabel(id, property string, rdx kevlar.ReadableRedux) []compton.FormattedLabel {

	labels := make([]compton.FormattedLabel, 0)

	val, _ := rdx.GetLastVal(property, id)
	values, _ := rdx.GetAllValues(property, id)

	switch property {
	case data.ArtTypeProperty:
		at := litres_integration.ParseArtType(val)
		label := compton.FormattedLabel{
			Property: property,
			Title:    at.String(),
		}
		labels = append(labels, label)
	case data.LitresLabelsProperty:
		for _, value := range values {
			label := compton.FormattedLabel{
				Property: property,
				Title:    value,
			}
			labels = append(labels, label)
		}
	case data.ArtFourthPresentProperty:
		if val == "true" {
			label := compton.FormattedLabel{
				Property: property,
				Title:    compton_data.PropertyTitles[data.ArtFourthPresentProperty],
			}
			labels = append(labels, label)
		}
	}

	return labels
}
