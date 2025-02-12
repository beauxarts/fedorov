package cli

import (
	"github.com/boggydigital/nod"
	"net/url"
	"strings"
)

func GetRecentSeriesHandler(u *url.URL) error {

	var artsIds []string
	if idstr := u.Query().Get("arts-id"); idstr != "" {
		artsIds = strings.Split(idstr, ",")
	}

	force := u.Query().Has("force")

	_, err := GetRecentSeries(force, artsIds...)
	return err
}

func GetRecentSeries(force bool, recentArtsIds ...string) ([]string, error) {
	grsa := nod.Begin("getting recent arts series...")
	defer grsa.Done()

	seriesIds, err := getSeriesIds(force, recentArtsIds...)
	if err != nil {
		return nil, err
	}

	grsa.EndWithResult(strings.Join(seriesIds, ","))

	return seriesIds, nil
}
