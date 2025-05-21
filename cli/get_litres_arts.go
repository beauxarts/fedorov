package cli

import (
	"errors"
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/nod"
	"net/http"
	"net/url"
	"strings"
)

func GetLitResArtsHandler(u *url.URL) error {
	var ids []string
	if idstr := u.Query().Get("id"); idstr != "" {
		ids = strings.Split(idstr, ",")
	}

	var artsTypesStr []string
	if ats := u.Query().Get("arts-type"); ats != "" {
		artsTypesStr = strings.Split(ats, ",")
	}

	artsTypes := make([]litres_integration.ArtsType, 0, len(artsTypesStr))

	if allArtsTypes := u.Query().Has("all-arts-types"); allArtsTypes {
		artsTypes = litres_integration.AllArtsTypes()
	} else {
		for _, ats := range artsTypesStr {
			artsTypes = append(artsTypes, litres_integration.ParseArtsType(ats))
		}
	}

	force := u.Query().Has("force")

	return GetLitResArts(artsTypes, nil, force, ids...)
}

func GetLitResArts(artsTypes []litres_integration.ArtsType, hc *http.Client, force bool, artsIds ...string) error {

	glaa := nod.NewProgress("getting litres arts...")
	defer glaa.Done()

	if len(artsIds) == 0 {
		var err error
		artsIds, err = GetRecentArts(force)
		if err != nil {
			return err
		}
	}

	glaa.TotalInt(len(artsIds))

	if hc == nil {
		var err error
		hc, err = getHttpClient()
		if err != nil {
			return err
		}
	}

	//dc := dolo.NewClient(hc, dolo.Defaults())

	for _, at := range artsTypes {
		if err := getSetArtsType(hc, at, force, artsIds...); err != nil {
			return err
		}
	}

	return nil
}

func getSetArtsType(hc *http.Client, at litres_integration.ArtsType, force bool, ids ...string) error {
	gsat := nod.NewProgress(" %s...", at)
	defer gsat.Done()

	absArtsTypeDir, err := data.AbsArtsTypeDir(at)
	if err != nil {
		return err
	}

	kv, err := kevlar.New(absArtsTypeDir, kevlar.JsonExt)
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

	gsat.TotalInt(len(newIds))

	errs := make(map[string]error)
	for _, id := range newIds {
		if err = getSetData(id, litres_integration.ArtsTypeUrl(at, id), hc, kv); err != nil {
			errs[id] = err
		}
		gsat.Increment()
	}

	if len(errs) > 0 {
		errStrs := make([]string, 0, len(errs))
		for id, err := range errs {
			errStrs = append(errStrs, fmt.Sprintf("%s: %s", id, err.Error()))
		}
		gsat.EndWithResult("errors: " + strings.Join(errStrs, ","))
	}

	return nil
}

func getSetData(id string, u *url.URL, hc *http.Client, kv kevlar.KeyValues) error {

	resp, err := hc.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.New(resp.Status)
	}

	return kv.Set(id, resp.Body)
}
