package cli

import (
	_ "image/jpeg"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/boggydigital/issa"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/redux"
)

func DehydrateHandler(u *url.URL) error {
	var artsIds []string
	if idstr := u.Query().Get("arts-id"); idstr != "" {
		artsIds = strings.Split(idstr, ",")
	}

	force := u.Query().Has("force")

	return Dehydrate(force, artsIds...)
}

func Dehydrate(force bool, artsIds ...string) error {

	di := nod.Begin("dehydrating images...")
	defer di.Done()

	properties := data.DehydratedProperties()
	properties = append(properties, data.ArtsOperationsOrderProperty)

	reduxDir := data.Pwd.AbsRelDirPath(data.Redux, data.Metadata)

	rdx, err := redux.NewWriter(reduxDir, properties...)
	if err != nil {
		return err
	}

	if force {
		artsIds, err = GetRecentArts(force)
		if err != nil {
			return err
		}
	}

	if err := dehydrateImages(
		rdx,
		data.DehydratedItemImageProperty,
		data.DehydratedItemImageModifiedProperty,
		data.RepItemImageColorProperty,
		litres_integration.CoverSizesDesc,
		force,
		artsIds...); err != nil {
		return err
	}

	if err := dehydrateImages(
		rdx,
		data.DehydratedListImageProperty,
		data.DehydratedListImageModifiedProperty,
		data.RepListImageColorProperty,
		litres_integration.CoverSizesAsc,
		force,
		artsIds...); err != nil {
		return err
	}

	return nil
}

func dehydrateImages(
	rdx redux.Writeable,
	imageProperty, modifiedProperty, repColorProperty string,
	sizes []litres_integration.CoverSize,
	force bool,
	ids ...string) error {

	di := nod.NewProgress(" dehydrating %s...", imageProperty)
	defer di.Done()

	di.TotalInt(len(ids))

	dehydratedImages := make(map[string][]string)
	dehydratedImageModified := make(map[string][]string)
	repColors := make(map[string][]string)

	for _, idStr := range ids {

		if _, ok := rdx.GetLastVal(imageProperty, idStr); ok && !force {
			continue
		}

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			nod.Log(err.Error())
			continue
		}

		for _, size := range sizes {

			acp := data.AbsCoverImagePath(id, size)
			if dhi, rc, err := issa.DehydrateImageRepColor(acp); err == nil {
				dehydratedImages[idStr] = []string{dhi}
				dehydratedImageModified[idStr] = []string{strconv.FormatInt(time.Now().Unix(), 10)}
				repColors[idStr] = []string{rc}
				// stop dehydrating at the best quality available
				break
			} else {
				nod.Log(err.Error())
			}
		}
		di.Increment()
	}

	if err := rdx.BatchReplaceValues(imageProperty, dehydratedImages); err != nil {
		return err
	}

	if err := rdx.BatchReplaceValues(modifiedProperty, dehydratedImageModified); err != nil {
		return err
	}

	if err := rdx.BatchReplaceValues(repColorProperty, repColors); err != nil {
		return err
	}

	return nil
}
