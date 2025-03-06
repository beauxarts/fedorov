package compton_fragments

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/redux"
)

func RatingAvg(id string, rdx redux.Readable) string {
	if rp, ok := rdx.GetLastVal(data.RatedAvgProperty, id); ok {
		return "Рейтинг: " + rp
	}
	return ""
}
