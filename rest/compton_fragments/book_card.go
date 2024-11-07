package compton_fragments

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/kevlar"
	"strconv"
	"strings"
	"time"
)

func BookCard(r compton.Registrar, id string, hydrated bool, rdx kevlar.ReadableRedux) compton.Element {
	bc := compton.Card(r, id)

	repColor := ""
	if rc, ok := rdx.GetLastVal(data.RepListImageColorProperty, id); ok {
		repColor = rc
		bc.SetAttribute("style", "background-color:color-mix(in display-p3,"+rc+" var(--cma), var(--c-background))")
	}

	posterUrl := "/list_cover?id=" + id
	dhSrc, _ := rdx.GetLastVal(data.DehydratedListImageProperty, id)
	placeholderSrc := dhSrc
	bc.AppendPoster(repColor, placeholderSrc, posterUrl, hydrated)

	bc.WidthPixels(80)
	bc.HeightPixels(120)

	if title, ok := rdx.GetLastVal(data.TitleProperty, id); ok {
		bc.AppendTitle(title)
	}

	artType := litres_integration.ArtTypeText
	if ats, ok := rdx.GetLastVal(data.ArtTypeProperty, id); ok {
		artType = litres_integration.ParseArtType(ats)
	}

	if labels := compton.Labels(r, FormatLabels(id, rdx)...).
		FontSize(size.XSmall).
		ColumnGap(size.XXSmall).
		RowGap(size.XXSmall); labels != nil {
		bc.AppendLabels(labels)
	}

	if authors, ok := rdx.GetAllValues(data.AuthorsProperty, id); ok && len(authors) > 0 {
		bc.AppendProperty(compton_data.PropertyTitles[data.AuthorsProperty], compton.Text(strings.Join(authors, ", ")))
	}

	if readers, ok := rdx.GetAllValues(data.ReadersProperty, id); ok && len(readers) > 0 {
		bc.AppendProperty(compton_data.PropertyTitles[data.ReadersProperty], compton.Text(strings.Join(readers, ", ")))
	} else if illustrators, ok := rdx.GetAllValues(data.IllustratorsProperty, id); ok && len(illustrators) > 0 {
		bc.AppendProperty(compton_data.ShortPropertyTitles[data.IllustratorsProperty], compton.Text(strings.Join(illustrators, ", ")))
	} else if translators, ok := rdx.GetAllValues(data.TranslatorsProperty, id); ok && len(translators) > 0 {
		bc.AppendProperty(compton_data.PropertyTitles[data.TranslatorsProperty], compton.Text(strings.Join(translators, ", ")))
	} else if dwa, ok := rdx.GetLastVal(data.DateWrittenAtProperty, id); ok {
		bc.AppendProperty(compton_data.PropertyTitles[data.DateWrittenAtProperty], compton.Text(fmtYearWrittenAt(dwa)))
	}

	if cpos, ok := rdx.GetLastVal(data.CurrentPagesOrSecondsProperty, id); ok {
		bc.AppendProperty(compton_data.PropertyTitles[data.CurrentPagesOrSecondsProperty],
			compton.Text(fmtCurrentPagesOrSeconds(cpos, artType)))
	}

	return bc
}

func fmtYearWrittenAt(dwa string) string {
	yearWrittenAt := 0
	if dateWrittenAt, err := time.Parse("2006-01-02", dwa); err == nil {
		if dateWrittenAt.Month() == 1 && dateWrittenAt.Day() == 1 {
			yearWrittenAt = dateWrittenAt.Year() - 1
		} else {
			yearWrittenAt = dateWrittenAt.Year()
		}
	}
	return strconv.Itoa(yearWrittenAt)
}
