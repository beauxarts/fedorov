package compton_fragments

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/kevlar"
)

var artTypesTitles = map[string]string{
	"0": "Текст",
	"1": "Аудио",
	"4": "PDF",
}

func FormatLabels(id string, rdx kevlar.ReadableRedux) []compton.FormattedLabel {
	fmtLabels := make([]compton.FormattedLabel, 0)

	fmtLabels = append(fmtLabels, formatLabel(id, data.ArtTypeProperty, rdx))

	return fmtLabels
}

func formatLabel(id, property string, rdx kevlar.ReadableRedux) compton.FormattedLabel {

	fmtLabel := compton.FormattedLabel{
		Property: property,
	}

	fmtLabel.Title, _ = rdx.GetLastVal(property, id)

	switch property {
	case data.ArtTypeProperty:
		fmtLabel.Title = artTypesTitles[fmtLabel.Title]
	}

	return fmtLabel
}
