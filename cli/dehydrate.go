package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/issa"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/nod"
	"image"
	"image/color"
	_ "image/jpeg"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func DehydrateHandler(u *url.URL) error {
	idSet := make(map[string]bool)
	if idstr := u.Query().Get("id"); idstr != "" {
		for _, id := range strings.Split(idstr, ",") {
			idSet[id] = true
		}
	}

	all := u.Query().Has("all")
	overwrite := u.Query().Has("overwrite")

	return Dehydrate(idSet, all, overwrite)
}

func Dehydrate(idSet map[string]bool, all, overwrite bool) error {

	di := nod.Begin("dehydrating images...")
	defer di.End()

	rdx, err := data.NewReduxWriter(
		data.DehydratedListImageProperty,
		data.DehydratedListImageModifiedProperty,
		data.DehydratedItemImageProperty,
		data.DehydratedItemImageModifiedProperty,
		data.ArtsHistoryOrderProperty)
	if err != nil {
		return di.EndWithError(err)
	}

	if all {
		if all, ok := rdx.GetAllValues(data.ArtsHistoryOrderProperty, data.ArtsHistoryOrderProperty); ok {
			for _, id := range all {
				idSet[id] = true
			}
		}
	}

	if err := dehydrateImages(
		idSet,
		rdx,
		data.DehydratedItemImageProperty,
		data.DehydratedItemImageModifiedProperty,
		data.CoverSizesDesc,
		overwrite); err != nil {
		return di.EndWithError(err)
	}

	if err := dehydrateImages(
		idSet,
		rdx,
		data.DehydratedListImageProperty,
		data.DehydratedListImageModifiedProperty,
		data.CoverSizesAsc,
		overwrite); err != nil {
		return di.EndWithError(err)
	}

	return nil
}

func dehydrateImages(
	idSet map[string]bool,
	rdx kevlar.WriteableRedux,
	imageProperty, modifiedProperty string,
	sizes []litres_integration.CoverSize,
	overwrite bool) error {

	di := nod.NewProgress(" dehydrating %s...", imageProperty)
	defer di.End()

	di.TotalInt(len(idSet))

	plt := issa.ColorPalette()

	dehydratedImages := make(map[string][]string)
	dehydratedImageModified := make(map[string][]string)

	for idStr := range idSet {

		if _, ok := rdx.GetLastVal(imageProperty, idStr); ok && !overwrite {
			continue
		}

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			nod.Log(err.Error())
			continue
		}

		for _, size := range sizes {

			acp, err := data.AbsCoverImagePath(id, size)
			if err != nil {
				return di.EndWithError(err)
			}
			if dhi, err := dehydrateImage(acp, plt); err == nil {
				dehydratedImages[idStr] = []string{dhi}
				dehydratedImageModified[idStr] = []string{strconv.FormatInt(time.Now().Unix(), 10)}
				// stop dehydrating at the best quality available
				break
			} else {
				nod.Log(err.Error())
			}
		}
		di.Increment()
	}

	if err := rdx.BatchReplaceValues(imageProperty, dehydratedImages); err != nil {
		return di.EndWithError(err)
	}

	if err := rdx.BatchReplaceValues(modifiedProperty, dehydratedImageModified); err != nil {
		return di.EndWithError(err)
	}

	return nil
}

func dehydrateImage(absImagePath string, plt color.Palette) (string, error) {
	dhi := ""

	fi, err := os.Open(absImagePath)
	if err != nil {
		return dhi, err
	}
	defer fi.Close()

	img, _, err := image.Decode(fi)
	if err != nil {
		return dhi, err
	}

	gif := issa.GIFImage(img, plt, issa.DefaultSampling)

	return issa.DehydrateColor(gif)
}
