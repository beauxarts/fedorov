package cli

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/kevlar_dolo"
	"github.com/boggydigital/nod"
	"golang.org/x/exp/maps"
	"net/http"
	"net/url"
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

	sessionId := u.Query().Get("session-id")
	force := u.Query().Has("force")

	return GetLitResAuthors(authorTypes, sessionId, nil, force, ids...)
}

func GetLitResAuthors(authorTypes []litres_integration.AuthorType, sessionId string, hc *http.Client, force bool, ids ...string) error {
	glaa := nod.NewProgress("getting litres authors...")
	defer glaa.End()

	if len(ids) == 0 {

		persons := make(map[string]interface{})

		rdx, err := data.NewReduxReader(data.ArtsHistoryOrderProperty, data.PersonsIdsProperty)
		if err != nil {
			return glaa.EndWithError(err)
		}

		if artsIds, ok := rdx.GetAllValues(data.ArtsHistoryOrderProperty, data.ArtsHistoryOrderProperty); ok {
			for _, id := range artsIds {
				if personsIds, sure := rdx.GetAllValues(data.PersonsIdsProperty, id); sure {
					for _, pid := range personsIds {
						persons[pid] = nil
					}
				}
			}
		}

		ids = maps.Keys(persons)
	}

	glaa.TotalInt(len(ids))

	if hc == nil {
		var err error
		hc, err = getHttpClient()
		if err != nil {
			return glaa.EndWithError(err)
		}
	}

	dc := dolo.NewClient(hc, dolo.Defaults())

	for _, at := range authorTypes {
		if err := getSetAuthorType(dc, at, force, ids...); err != nil {
			return glaa.EndWithError(err)
		}
	}

	glaa.EndWithResult("done")

	return nil
}

func getSetAuthorType(dc *dolo.Client, at litres_integration.AuthorType, force bool, ids ...string) error {
	gsat := nod.NewProgress(" %s...", at)
	defer gsat.End()

	absAuthorTypeDir, err := data.AbsAuthorTypeDir(at)
	if err != nil {
		return gsat.EndWithError(err)
	}

	kv, err := kevlar.NewKeyValues(absAuthorTypeDir, kevlar.JsonExt)
	if err != nil {
		return gsat.EndWithError(err)
	}

	newIds := make([]string, 0, len(ids))
	for _, id := range ids {
		ok, err := kv.Has(id)
		if err != nil {
			return gsat.EndWithError(err)
		}
		if ok && !force {
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
