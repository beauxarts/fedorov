package compton_fragments

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/kevlar"
	"strconv"
	"strings"
	"time"
)

func BookCard(r compton.Registrar, id string, hydrated bool, rdx kevlar.ReadableRedux) compton.Element {
	bc := compton.Card(r, id)

	posterUrl := "/list_cover?id=" + id
	dhSrc, _ := rdx.GetLastVal(data.DehydratedListImageProperty, id)
	placeholderSrc := dhSrc
	bc.AppendPoster(placeholderSrc, posterUrl, hydrated)

	bc.WidthPixels(100)
	bc.HeightPixels(100)

	if title, ok := rdx.GetLastVal(data.TitleProperty, id); ok {
		bc.AppendTitle(title)
	}

	if labels := compton.Labels(r, FormatLabels(id, rdx)...).
		FontSize(size.XSmall).
		ColumnGap(size.XXSmall).
		RowGap(size.XXSmall); labels != nil {
		bc.AppendLabels(labels)
	}

	if authors, ok := rdx.GetAllValues(data.AuthorsProperty, id); ok && len(authors) > 0 {
		bc.AppendProperty("Автор", compton.Text(strings.Join(authors, ", ")))
	}

	if dwa, ok := rdx.GetLastVal(data.DateWrittenAtProperty, id); ok {
		yearWrittenAt := 0
		if dateWrittenAt, err := time.Parse("2006-01-02", dwa); err == nil {
			if dateWrittenAt.Month() == 1 && dateWrittenAt.Day() == 1 {
				yearWrittenAt = dateWrittenAt.Year() - 1
			} else {
				yearWrittenAt = dateWrittenAt.Year()
			}
		}
		bc.AppendProperty("Написано", compton.Text(strconv.Itoa(yearWrittenAt)))

	}

	return bc
}
