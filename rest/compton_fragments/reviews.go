package compton_fragments

import (
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/size"
	"strconv"
	"strings"
)

func Reviews(r compton.Registrar, artsReviews *litres_integration.ArtsReviews) compton.Element {

	stack := compton.FlexItems(r, direction.Column)

	if len(artsReviews.Payload.Data) == 0 {
		stack.Append(compton.Fspan(r, "Для данной книги пока нет отзывов").
			ForegroundColor(color.RepGray).
			TextAlign(align.Center))
	}

	for ii, review := range artsReviews.Payload.Data {

		metadataFrow := compton.Frow(r).
			FontSize(size.XSmall)
		if review.ItemRating != nil {
			ir := strconv.Itoa(*review.ItemRating)
			metadataFrow.PropVal("Оценка", ir)
		}
		metadataFrow.PropVal("Автор", review.UserDisplayName)
		cat, _, _ := strings.Cut(review.CreatedAt, "T")
		metadataFrow.PropVal("Написано", cat)

		stack.Append(metadataFrow)

		stack.Append(compton.PreText(review.Text))

		likesFrow := compton.Frow(r).
			FontSize(size.XSmall)

		if review.LikesCount > 0 {
			likesFrow.PropVal("Оценили", strconv.Itoa(review.LikesCount))
		}
		if review.DislikesCount > 0 {
			likesFrow.PropVal("Не оценили", strconv.Itoa(review.DislikesCount))
		}
		stack.Append(likesFrow)

		if ii != len(artsReviews.Payload.Data)-1 {
			stack.Append(compton.Hr())
		}
	}

	return stack
}
