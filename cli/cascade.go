package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/nod"
	"net/url"
)

func CascadeHandler(_ *url.URL) error {
	return Cascade()
}

func Cascade() error {

	ca := nod.Begin("cascading reductions...")
	defer ca.End()

	rdx, err := data.NewReduxWriter(data.ReduxProperties()...)
	if err != nil {
		return ca.EndWithError(err)
	}

	if err := cascadePersonsRolesProperties(rdx); err != nil {
		return ca.EndWithError(err)
	}

	if err := cascadeBookCompletedProperty(rdx); err != nil {
		return ca.EndWithError(err)
	}

	ca.EndWithResult("done")

	return nil
}

func cascadePersonsRolesProperties(rdx kevlar.WriteableRedux) error {

	cprpa := nod.NewProgress(" persons roles...")
	defer cprpa.End()

	authors := make(map[string][]string)
	illustrators := make(map[string][]string)
	painters := make(map[string][]string)
	performers := make(map[string][]string)
	publishers := make(map[string][]string)
	readers := make(map[string][]string)
	translators := make(map[string][]string)

	keys := rdx.Keys(data.PersonsIdsProperty)
	cprpa.TotalInt(len(keys))

	for _, id := range keys {

		if persons, ok := rdx.GetAllValues(data.PersonsIdsProperty, id); ok && len(persons) > 0 {
			if roles, sure := rdx.GetAllValues(data.PersonsRolesProperty, id); sure && len(roles) == len(persons) {
				for ii, role := range roles {
					var propertyMap map[string][]string
					switch role {
					case "author":
						propertyMap = authors
					case "illustrator":
						propertyMap = illustrators
					case "painter":
						propertyMap = painters
					case "performer":
						propertyMap = performers
					case "publisher":
						propertyMap = publishers
					case "reader":
						propertyMap = readers
					case "translator":
						propertyMap = translators
					default:
						continue
					}

					personId := persons[ii]

					if fullName, yes := rdx.GetLastVal(data.PersonFullNameProperty, personId); yes && fullName != "" {
						propertyMap[id] = append(propertyMap[id], fullName)
					}
				}
			}
		}
		cprpa.Increment()
	}

	if err := rdx.BatchReplaceValues(data.AuthorsProperty, authors); err != nil {
		return cprpa.EndWithError(err)
	}
	if err := rdx.BatchReplaceValues(data.IllustratorsProperty, illustrators); err != nil {
		return cprpa.EndWithError(err)
	}
	if err := rdx.BatchReplaceValues(data.PaintersProperty, painters); err != nil {
		return cprpa.EndWithError(err)
	}
	if err := rdx.BatchReplaceValues(data.PerformersProperty, performers); err != nil {
		return cprpa.EndWithError(err)
	}
	if err := rdx.BatchReplaceValues(data.PublishersProperty, publishers); err != nil {
		return cprpa.EndWithError(err)
	}
	if err := rdx.BatchReplaceValues(data.ReadersProperty, readers); err != nil {
		return cprpa.EndWithError(err)
	}
	if err := rdx.BatchReplaceValues(data.TranslatorsProperty, translators); err != nil {
		return cprpa.EndWithError(err)
	}

	cprpa.EndWithResult("done")
	return nil
}

func cascadeBookCompletedProperty(rdx kevlar.WriteableRedux) error {

	bca := nod.NewProgress(" " + data.BookCompletedProperty)
	defer bca.End()

	if err := rdx.MustHave(data.TitleProperty, data.BookCompletedProperty); err != nil {
		return bca.EndWithError(err)
	}

	ids := rdx.Keys(data.TitleProperty)
	bca.TotalInt(len(ids))

	completed := make(map[string][]string)

	for _, id := range ids {
		bca.Increment()
		if val, ok := rdx.GetLastVal(data.BookCompletedProperty, id); ok && val != "" {
			completed[id] = []string{"true"}
		}
		completed[id] = []string{"false"}
	}

	if err := rdx.BatchReplaceValues(data.BookCompletedProperty, completed); err != nil {
		return bca.EndWithError(err)
	}

	bca.EndWithResult("done")

	return nil
}
