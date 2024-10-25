package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/nod"
	"net/url"
	"strings"
	"time"
)

const recentDays = 1

func GetRecentArtsHandler(u *url.URL) error {
	force := u.Query().Has("force")

	_, err := GetRecentArts(force)
	return err
}

func GetRecentArts(force bool) ([]string, error) {
	graa := nod.Begin("getting recent arts...")
	defer graa.End()

	rdx, err := data.NewReduxReader(data.ArtsHistoryEventTimeProperty)
	if err != nil {
		return nil, graa.EndWithError(err)
	}

	ids := make([]string, 0)

	if force {
		ids = rdx.Keys(data.ArtsHistoryEventTimeProperty)
	} else {

		earliestDate := time.Now().AddDate(0, 0, -recentDays)

		for _, id := range rdx.Keys(data.ArtsHistoryEventTimeProperty) {
			if ets, ok := rdx.GetLastVal(data.ArtsHistoryEventTimeProperty, id); ok {
				if et, err := time.Parse("2006-01-02T15:04:05Z", ets); err == nil {
					if et.After(earliestDate) {
						ids = append(ids, id)
					}
					continue
				}
				return nil, graa.EndWithError(err)
			}
		}
	}

	if len(ids) == 0 {
		graa.EndWithResult("found nothing")
	} else {
		graa.EndWithResult(strings.Join(ids, ","))
	}

	return ids, nil
}
