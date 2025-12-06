package compton_fragments

import (
	"strings"

	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/issa"
	"github.com/boggydigital/redux"
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

	var repColor = issa.NeutralRepColor
	posterUrl := "/list_cover?id=" + id

	if rc, ok := rdx.GetLastVal(data.RepListImageColorProperty, id); ok {
		repColor = rc
	}

	var placeholderSrc string
	if dhSrc, ok := rdx.GetLastVal(data.DehydratedListImageProperty, id); ok {
		placeholderSrc = dhSrc
	}

	bc.SetAttribute("style", "--c-rep:"+repColor)
	poster := bc.AppendPoster(repColor, placeholderSrc, posterUrl, hydrated)

	poster.WidthPixels(80)
	poster.HeightPixels(120)

	if title, ok := rdx.GetLastVal(data.TitleProperty, id); ok {
		bc.AppendTitle(title)
	} else {
		bc.AppendTitle("[БЕЗ НАЗВАНИЯ]")
	}

	bc.AppendBadges(compton.Badges(r, FormatBadges(id, rdx)...))

	properties, values := SummarizeBookProperties(id, rdx)
	for _, p := range properties {
		bc.AppendProperty(compton_data.ShortPropertyTitles[p], compton.Text(strings.Join(values[p], ", ")))
		//pp.SetAttribute("style", "view-transition-name:"+p+id)
	}

	return bc
}
