package compton_fragments

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/kevlar"
	"strconv"
	"strings"
)

func Similar(r compton.Registrar, id string, artsSimilar *litres_integration.ArtsSimilar, rdx kevlar.ReadableRedux) compton.Element {

	stack := compton.FlexItems(r, direction.Column)

	if artsSimilar == nil || len(artsSimilar.Payload.Data) == 0 {
		stack.Append(compton.Fspan(r, "Для данной книги пока нет сходных книг").
			ForegroundColor(color.Gray).
			TextAlign(align.Center))
		return stack
	}

	for ii, art := range artsSimilar.Payload.Data {

		linkHref := "https://litres.ru" + art.Url
		linkColor := color.Cyan

		artId := strconv.Itoa(art.Id)
		if rdx.HasKey(data.TitleProperty, artId) {
			linkHref = "/book?id=" + artId
			linkColor = color.Foreground
		}

		link := compton.A(linkHref)
		link.SetAttribute("target", "_top")

		linkStack := compton.FlexItems(r, direction.Column).RowGap(size.Unset)

		linkTitleRow := compton.FlexItems(r, direction.Row).
			ColumnGap(size.Small).
			RowGap(size.XSmall).
			JustifyContent(align.Center).
			AlignItems(align.Center)

		linkTitle := compton.Fspan(r, art.Title).
			TextAlign(align.Center).
			ForegroundColor(linkColor).
			FontWeight(font_weight.Bolder)

		linkLabels := compton.Labels(r, FormatLabels(id, rdx)...).
			FontSize(size.XSmall).
			RowGap(size.XSmall).
			ColumnGap(size.XSmall)

		linkTitleRow.Append(linkTitle, linkLabels)

		linkStack.Append(linkTitleRow)

		frow := compton.Frow(r)

		//frow.Elements(linkLabels)

		authors := make([]string, 0)
		for _, person := range art.Persons {
			if person.Role == "author" {
				authors = append(authors, person.FullName)
			}
		}
		if len(authors) > 0 {
			frow.PropVal("Автор", strings.Join(authors, ", "))
		}
		frow.PropVal("Рейтинг", fmt.Sprintf("%.1f", art.Rating.RatedAvg))

		if art.DateWrittenAt != "" {
			frow.PropVal("Написано", fmtYearWrittenAt(art.DateWrittenAt))
		}

		linkStack.Append(compton.FICenter(r, frow))

		link.Append(linkStack)
		stack.Append(link)

		if ii != len(artsSimilar.Payload.Data)-1 {
			stack.Append(compton.Hr())
		}
	}

	return stack
}
