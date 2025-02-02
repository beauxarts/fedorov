package compton_fragments

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/redux"
	"strings"
)

func SummarizeBookProperties(id string, rdx redux.Readable) ([]string, map[string][]string) {
	properties := make([]string, 0)
	values := make(map[string][]string)

	properties = append(properties, data.AuthorsProperty)
	values[data.AuthorsProperty], _ = rdx.GetAllValues(data.AuthorsProperty, id)

	artType := litres_integration.ArtTypeText
	if ats, ok := rdx.GetLastVal(data.ArtTypeProperty, id); ok {
		artType = litres_integration.ParseArtType(ats)
	}

	if readers, ok := rdx.GetAllValues(data.ReadersProperty, id); ok && len(readers) > 0 {
		properties = append(properties, data.ReadersProperty)
		values[data.ReadersProperty] = readers
	} else if illustrators, ok := rdx.GetAllValues(data.IllustratorsProperty, id); ok && len(illustrators) > 0 {
		properties = append(properties, data.IllustratorsProperty)
		values[data.IllustratorsProperty] = illustrators
	} else if translators, ok := rdx.GetAllValues(data.TranslatorsProperty, id); ok && len(translators) > 0 {
		properties = append(properties, data.TranslatorsProperty)
		values[data.TranslatorsProperty] = translators
	} else if dwa, ok := rdx.GetLastVal(data.DateWrittenAtProperty, id); ok {
		properties = append(properties, data.DateWrittenAtProperty)
		values[data.DateWrittenAtProperty] = []string{fmtYearWrittenAt(dwa)}
	}

	if cpos, ok := rdx.GetLastVal(data.CurrentPagesOrSecondsProperty, id); ok {
		properties = append(properties, data.CurrentPagesOrSecondsProperty)
		values[data.CurrentPagesOrSecondsProperty] = []string{fmtCurrentPagesOrSeconds(cpos, artType)}
	}

	return properties, values
}

func BookCard(r compton.Registrar, id string, hydrated bool, rdx redux.Readable) compton.Element {
	bc := compton.Card(r, id)

	repColor := ""
	if rc, ok := rdx.GetLastVal(data.RepListImageColorProperty, id); ok {
		repColor = rc
		compton.SetTint(bc, repColor)
	}

	posterUrl := "/list_cover?id=" + id
	dhSrc, _ := rdx.GetLastVal(data.DehydratedListImageProperty, id)
	placeholderSrc := dhSrc
	bc.AppendPoster(repColor, placeholderSrc, posterUrl, hydrated)

	bc.WidthPixels(80)
	bc.HeightPixels(120)

	if title, ok := rdx.GetLastVal(data.TitleProperty, id); ok {
		bc.AppendTitle(title)
	} else {
		bc.AppendTitle("[БЕЗ НАЗВАНИЯ]")
	}

	if labels := compton.Labels(r, FormatLabels(id, rdx)...).
		FontSize(size.XSmall).
		ColumnGap(size.XXSmall).
		RowGap(size.XXSmall); labels != nil {
		bc.AppendLabels(labels)
	}

	properties, values := SummarizeBookProperties(id, rdx)
	for _, p := range properties {
		bc.AppendProperty(compton_data.PropertyTitles[p], compton.Text(strings.Join(values[p], ", ")))
	}

	return bc
}
