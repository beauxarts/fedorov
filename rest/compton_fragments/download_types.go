package compton_fragments

import (
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/size"
	"strconv"
)

var typeColors = map[string]color.Color{
	"mobi.prc":        color.Foreground,
	"azw3":            color.Foreground,
	"epub":            color.Foreground,
	"audio/m4b":       color.Blue,
	"application/pdf": color.Red,
}

func DownloadType(r compton.Registrar, id string, dt *litres_integration.ArtsFilesData) compton.Element {

	downloadStack := compton.FlexItems(r, direction.Column).RowGap(size.Small)
	downloadStack.AddClass("download-type")

	downloadLink := compton.A("/file?id=" + id + "&file=" + dt.TypeFilename())

	downloadLinkStack := compton.FlexItems(r, direction.Column).RowGap(size.Small)
	downloadLink.Append(downloadLinkStack)

	filename := compton.Fspan(r, dt.TypeFilenameSansExt()).FontWeight(font_weight.Bolder)
	downloadLinkStack.Append(filename)

	row := compton.Frow(r).
		FontSize(size.Small)

	row.IconColor(compton.Circle, typeColors[dt.Type()]).
		PropVal("Тип", dt.TypeDescription()).
		PropVal("Формат", dt.Type()).
		PropVal("Размер", fmtBytes(dt.Size))

	artType := litres_integration.ArtTypeText
	cpos := ""
	if dt.Pages != nil {
		cpos = strconv.Itoa(*dt.Pages)
	}
	if dt.Seconds != nil {
		artType = litres_integration.ArtTypeAudio
		cpos = strconv.Itoa(*dt.Seconds)
	}
	row.PropVal("Объем", fmtCurrentPagesOrSeconds(cpos, artType))

	downloadLinkStack.Append(row)

	downloadStack.Append(downloadLink)

	return downloadStack
}
