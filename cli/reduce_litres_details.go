package cli

import (
	"errors"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/litres_integration"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"golang.org/x/net/html"
	"net/url"
)

func ReduceLitResDetailsHandler(u *url.URL) error {

	scoreData := true
	if u.Query().Get("score-data") == "false" {
		scoreData = false
	}
	return ReduceLitResBooksDetails(scoreData)
}

func ReduceLitResBooksDetails(scoreData bool) error {

	rmbda := nod.NewProgress("reducing details...")
	defer rmbda.End()

	reduxProps := data.ReduxProperties()

	reductions := make(map[string]map[string][]string, len(reduxProps))
	for _, p := range reduxProps {
		reductions[p] = make(map[string][]string)
	}

	missingDetails := make([]string, 0)

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), nil, reduxProps...)
	if err != nil {
		return rmbda.EndWithError(err)
	}

	kv, err := kvas.ConnectLocal(data.AbsMyBooksDetailsDir(), kvas.HtmlExt)
	if err != nil {
		return rmbda.EndWithError(err)
	}

	ids := kv.Keys()

	rmbda.TotalInt(len(ids))

	dataScore := make(map[string]int)

	for _, id := range ids {

		// don't attempt reducing imported books
		if IsImported(id, rxa) {
			continue
		}

		lrdx, err := ReduceLitResBookDetails(id, kv)
		if err != nil {
			return rmbda.EndWithError(err)
		}

		if isEmpty(lrdx) {
			missingDetails = append(missingDetails, id)
		}

		MapLitresToFedorov(id, lrdx, reductions)

		if scoreData {
			for lp, vals := range lrdx {
				if p, ok := data.LitResPropertyMap[lp]; ok {
					if evs, ok := rxa.GetAllUnchangedValues(p, id); ok {
						dataScore[id] = len(vals) - len(evs)
					}
				}
			}
		}

		rmbda.Increment()
	}

	if scoreData {
		overallDataScore := 0
		for _, score := range dataScore {
			overallDataScore += score
		}

		//data scoring threshold is number of books
		//meaning either big change on small number of books
		//or 1 change on every book in the collection
		if overallDataScore < -len(ids) {
			return rmbda.EndWithError(errors.New("details reduction produced less data than already existed"))
		}
	}

	sra := nod.NewProgress(" saving reductions...")
	defer sra.End()

	sra.TotalInt(len(reductions))

	if err := rxa.ReplaceValues(data.MissingDetailsIdsProperty, data.MissingDetailsIdsProperty, missingDetails...); err != nil {
		rmbda.EndWithError(err)
	}

	for prop, rdx := range reductions {
		if err := rxa.BatchReplaceValues(prop, rdx); err != nil {
			return rmbda.EndWithError(err)
		}
		sra.Increment()
	}

	sra.EndWithResult("done")
	rmbda.EndWithResult("done")

	return nil
}

func ReduceLitResBookDetails(id string, kv kvas.KeyValues) (map[string][]string, error) {
	det, err := kv.Get(id)
	defer det.Close()

	if err != nil {
		return nil, err
	}

	body, err := html.Parse(det)
	if err != nil {
		return nil, err
	}

	return litres_integration.ReduceDetails(body)
}

func MapLitresToFedorov(id string, lrdx map[string][]string, rdx map[string]map[string][]string) {
	for lp, vals := range lrdx {
		if p, ok := data.LitResPropertyMap[lp]; ok {
			if p == litres_integration.KnownIrrelevantProperty {
				continue
			}
			rdx[p][id] = vals
		} else {
			nod.Log("unknown LitRes property %s", lp)
		}
	}
}

func isEmpty(rdx map[string][]string) bool {
	isEmpty := true

	for _, vals := range rdx {
		isEmpty = isEmpty && len(vals) == 0
	}

	return isEmpty
}
