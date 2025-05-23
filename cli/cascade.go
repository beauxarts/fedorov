package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathways"
	"github.com/boggydigital/redux"
	"net/url"
)

func CascadeHandler(_ *url.URL) error {
	return Cascade()
}

func Cascade() error {

	ca := nod.Begin("cascading reductions...")
	defer ca.Done()

	reduxDir, err := pathways.GetAbsRelDir(data.Redux)
	if err != nil {
		return err
	}

	rdx, err := redux.NewWriter(reduxDir, data.ReduxProperties()...)
	if err != nil {
		return err
	}

	if err = cascadePersonsRolesProperties(rdx); err != nil {
		return err
	}

	if err = cascadeIdNameProperties(rdx); err != nil {
		return err
	}

	return nil
}

func cascadePersonsRolesProperties(rdx redux.Writeable) error {

	cprpa := nod.NewProgress(" persons roles...")
	defer cprpa.Done()

	authors := make(map[string][]string)
	illustrators := make(map[string][]string)
	painters := make(map[string][]string)
	performers := make(map[string][]string)
	readers := make(map[string][]string)
	translators := make(map[string][]string)

	cprpa.TotalInt(rdx.Len(data.PersonsIdsProperty))

	for id := range rdx.Keys(data.PersonsIdsProperty) {

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
		return err
	}
	if err := rdx.BatchReplaceValues(data.IllustratorsProperty, illustrators); err != nil {
		return err
	}
	if err := rdx.BatchReplaceValues(data.PaintersProperty, painters); err != nil {
		return err
	}
	if err := rdx.BatchReplaceValues(data.PerformersProperty, performers); err != nil {
		return err
	}
	if err := rdx.BatchReplaceValues(data.ReadersProperty, readers); err != nil {
		return err
	}
	if err := rdx.BatchReplaceValues(data.TranslatorsProperty, translators); err != nil {
		return err
	}

	return nil
}

func cascadeIdNameProperties(rdx redux.Writeable) error {

	cinpa := nod.NewProgress(" id, name properties...")
	defer cinpa.Done()

	idNameProperties := map[string]string{
		data.GenresIdsProperty:       data.GenreNameProperty,
		data.TagsIdsProperty:         data.TagNameProperty,
		data.PublisherIdProperty:     data.PublisherNameProperty,
		data.RightholdersIdsProperty: data.RightholderNameProperty,
		data.SeriesIdProperty:        data.SeriesNameProperty,
	}
	outputProperties := map[string]string{
		data.GenresIdsProperty:       data.GenresProperty,
		data.TagsIdsProperty:         data.TagsProperty,
		data.PublisherIdProperty:     data.PublishersProperty,
		data.RightholdersIdsProperty: data.RightholdersProperty,
		data.SeriesIdProperty:        data.SeriesProperty,
	}

	cinpa.TotalInt(len(idNameProperties))

	for idp, np := range idNameProperties {
		propertyValues := cascadeIdNameProperty(idp, np, rdx)
		if err := rdx.BatchReplaceValues(outputProperties[idp], propertyValues); err != nil {
			return err
		}
		cinpa.Increment()
	}

	return nil
}

func cascadeIdNameProperty(idProperty, nameProperty string, rdx redux.Writeable) map[string][]string {

	values := make(map[string][]string)

	for id := range rdx.Keys(idProperty) {

		if nameIds, ok := rdx.GetAllValues(idProperty, id); ok {
			for _, nameId := range nameIds {
				if name, sure := rdx.GetLastVal(nameProperty, nameId); sure {
					values[id] = append(values[id], name)
				}
			}
		}
	}

	return values
}
