package cli

import (
	"github.com/boggydigital/nod"
	"net/url"
	"strings"
)

func GetRecentPersonsHandler(u *url.URL) error {

	var artsIds []string
	if idstr := u.Query().Get("arts-id"); idstr != "" {
		artsIds = strings.Split(idstr, ",")
	}

	force := u.Query().Has("force")

	_, err := GetRecentPersons(force, artsIds...)
	return err
}

func GetRecentPersons(force bool, recentArtsIds ...string) ([]string, error) {
	grpa := nod.Begin("getting recent arts persons...")
	defer grpa.End()

	personsIds, err := getPersonsIds(force, recentArtsIds...)
	if err != nil {
		return nil, grpa.EndWithError(err)
	}

	grpa.EndWithResult(strings.Join(personsIds, ","))

	return personsIds, nil
}
