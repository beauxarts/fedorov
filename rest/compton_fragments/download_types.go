package compton_fragments

import (
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/font_weight"
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

	downloadStack := compton.FlexItems(r, direction.Column)

	downloadLink := compton.A("/file?id=" + id + "&file=" + dt.TypeFilename())

	filename := compton.Fspan(r, dt.TypeFilenameSansExt()).FontWeight(font_weight.Bolder)
	downloadLink.Append(filename)

	row := compton.Frow(r)

	row.IconColor(typeColors[dt.Type()])
	row.PropVal("Тип", dt.TypeDescription())
	row.PropVal("Формат", dt.Type())
	row.PropVal("Размер", fmtBytes(dt.Size))

	if dt.Pages != nil {
		row.PropVal("Объем (страниц)", strconv.Itoa(*dt.Pages))
	}

	if dt.Seconds != nil {
		row.PropVal("Длительность", fmtSeconds(*dt.Seconds))
	}

	downloadLink.Append(row)

	downloadStack.Append(downloadLink)

	return downloadStack
}
