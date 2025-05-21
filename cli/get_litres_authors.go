package cli

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/kevlar_dolo"
	"github.com/boggydigital/nod"
	"maps"
	"net/http"
	"net/url"
	"slices"
	"strings"
)

func GetLitResAuthorsHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	var authorTypesStr []string
	if ats := u.Query().Get("author-type"); ats != "" {
		authorTypesStr = strings.Split(ats, ",")
	}

	authorTypes := make([]litres_integration.AuthorType, 0, len(authorTypesStr))

	if allAuthorTypes := u.Query().Has("all-author-types"); allAuthorTypes {
		authorTypes = litres_integration.AllAuthorTypes()
	} else {
		for _, ats := range authorTypesStr {
			authorTypes = append(authorTypes, litres_integration.ParseAuthorType(ats))
		}
	}

	force := u.Query().Has("force")

	return GetLitResAuthors(authorTypes, nil, force, ids...)
}

func GetLitResAuthors(authorTypes []litres_integration.AuthorType, hc *http.Client, force bool, personsIds ...string) error {
	glaa := nod.NewProgress("getting litres authors...")
	defer glaa.Done()

	if len(personsIds) == 0 {
		var err error
		personsIds, err = getPersonsIds(force)
		if err != nil {
			return err
		}
	}

	glaa.TotalInt(len(personsIds))

	if hc == nil {
		var err error
		hc, err = getHttpClient()
		if err != nil {
			return err
		}
	}

	dc := dolo.NewClient(hc, dolo.Defaults())

	for _, at := range authorTypes {
		if err := getSetAuthorType(dc, at, force, personsIds...); err != nil {
			return err
		}
	}

	return nil
}

func getPersonsIds(force bool, artsIds ...string) ([]string, error) {
	persons := make(map[string]interface{})

	rdx, err := data.NewReduxReader(data.ArtsOperationsOrderProperty, data.PersonsIdsProperty)
	if err != nil {
		return nil, err
	}

	if len(artsIds) == 0 && force {
		if allArtsIds, ok := rdx.GetAllValues(data.ArtsOperationsOrderProperty, data.ArtsOperationsOrderProperty); ok {
			artsIds = allArtsIds
		}
	}

	for _, id := range artsIds {
		if personsIds, sure := rdx.GetAllValues(data.PersonsIdsProperty, id); sure {
			for _, pid := range personsIds {
				persons[pid] = nil
			}
		}
	}

	return slices.Collect(maps.Keys(persons)), nil
}

func getSetAuthorType(dc *dolo.Client, at litres_integration.AuthorType, force bool, ids ...string) error {
	gsat := nod.NewProgress(" %s...", at)
	defer gsat.Done()

	absAuthorTypeDir, err := data.AbsAuthorTypeDir(at)
	if err != nil {
		return err
	}

	kv, err := kevlar.New(absAuthorTypeDir, kevlar.JsonExt)
	if err != nil {
		return err
	}

	newIds := make([]string, 0, len(ids))
	for _, id := range ids {
		if kv.Has(id) && !force {
			continue
		}
		newIds = append(newIds, id)
	}

	indexSetter := kevlar_dolo.NewIndexSetter(kv, newIds...)
	urls := make([]*url.URL, 0, len(newIds))
	for _, id := range newIds {
		urls = append(urls, litres_integration.AuthorUrl(at, id))
	}

	result := "done"

	if errs := dc.GetSet(urls, indexSetter, gsat, force); len(errs) > 0 {
		errIds := make([]string, 0, len(errs))
		for ii := range errs {
			errIds = append(errIds, newIds[ii])
		}
		result = fmt.Sprintf("GetSet error ids: %s", strings.Join(errIds, ","))
	}

	gsat.EndWithResult(result)

	return nil
}
