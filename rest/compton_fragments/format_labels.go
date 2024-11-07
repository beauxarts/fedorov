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
		fmtLabels = append(fmtLabels, formatLabel(id, p, rdx))
	}

	return fmtLabels
}

func formatLabel(id, property string, rdx kevlar.ReadableRedux) compton.FormattedLabel {

	fmtLabel := compton.FormattedLabel{
		Property: property,
	}

	val, _ := rdx.GetLastVal(property, id)

	switch property {
	case data.ArtTypeProperty:
		at := litres_integration.ParseArtType(val)
		fmtLabel.Title = at.String()
	}

	return fmtLabel
}
