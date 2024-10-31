package compton_pages

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
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

func Files(id string) compton.PageElement {
	s := compton_fragments.ProductSection(compton_data.FilesSection)

	filesReader, err := data.NewArtsReader(litres_integration.ArtsTypeFiles)
	if err != nil {
		return s.Error(err)
	}

	artFiles, err := filesReader.ArtsFiles(id)
	if err != nil {
		return s.Error(err)
	}

	pageStack := compton.FlexItems(s, direction.Column)
	s.Append(pageStack)

	dlTypes := artFiles.PreferredDownloadsTypes()
	for ii, dt := range dlTypes {
		if dte := downloadType(s, id, dt); dte != nil {
			pageStack.Append(dte)
		}
		if ii != len(dlTypes)-1 {
			pageStack.Append(compton.Hr())
		}
	}

	return s
}

func downloadType(r compton.Registrar, id string, dt *litres_integration.ArtsFilesData) compton.Element {

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

func fmtBytes(b int) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func fmtSeconds(ts int) string {
	if ts == 0 {
		return "unknown"
	}

	hours := ts / (60 * 60)
	minutes := (ts / 60) % 60
	seconds := ts % 60

	if hours == 0 {
		return fmt.Sprintf("%2d:%2d", minutes, seconds)
	} else {
		return fmt.Sprintf("%d:%2d:%2d", hours, minutes, seconds)
	}
}
