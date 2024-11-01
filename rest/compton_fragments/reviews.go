package compton_fragments

import (
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/direction"
	"strconv"
	"strings"
)

func Reviews(r compton.Registrar, id string, artsReviews *litres_integration.ArtsReviews) compton.Element {

	stack := compton.FlexItems(r, direction.Column)

	for ii, review := range artsReviews.Payload.Data {

		metadataFrow := compton.Frow(r)
		if review.ItemRating != nil {
			ir := strconv.Itoa(*review.ItemRating)
			metadataFrow.PropVal("Оценка", ir)
		}
		metadataFrow.PropVal("Автор", review.UserDisplayName)
		cat, _, _ := strings.Cut(review.CreatedAt, "T")
		metadataFrow.PropVal("Написано", cat)

		stack.Append(metadataFrow)

		stack.Append(compton.DivText(review.Text))

		likesFrow := compton.Frow(r)

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
