package cli

import (
	"encoding/json"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/nod"
	"net/url"
	"strconv"
)

func ReduceLitResOperationsHandler(_ *url.URL) error {
	return ReduceLitResOperations()
}

func ReduceLitResOperations() error {

	roa := nod.NewProgress("reducing operations...")
	defer roa.EndWithResult("done")

	absLitResOperationsDir, err := data.AbsDataTypeDir(litres_integration.LitResOperations)
	if err != nil {
		return roa.EndWithError(err)
	}

	kv, err := kevlar.NewKeyValues(absLitResOperationsDir, kevlar.JsonExt)
	if err != nil {
		return roa.EndWithError(err)
	}

	keys, err := kv.Keys()
	if err != nil {
		return roa.EndWithError(err)
	}

	totalPages := len(keys)
	artsOperationsOrder := make([]string, 0, totalPages*litres_integration.OperationsLimit)
	artsOperationsEventTimes := make(map[string][]string, totalPages*litres_integration.OperationsLimit)

	for p := 1; p <= totalPages; p++ {

		if ops, err := getArtsOperations(p, kv); err == nil {

			for _, dt := range ops.Payload.Data {
				for _, art := range dt.SpecificData.Arts {
					artId := strconv.Itoa(art.Id)
					artsOperationsOrder = append(artsOperationsOrder, artId)
					artsOperationsEventTimes[artId] = []string{dt.Date}
				}
			}

		} else {
			return roa.EndWithError(err)
		}

		roa.Increment()
	}

	rdx, err := data.NewReduxWriter(
		data.ArtsOperationsOrderProperty,
		data.ArtsOperationsEventTimeProperty)
	if err != nil {
		return roa.EndWithError(err)
	}

	sra := nod.Begin(" saving redux...")
	defer sra.EndWithResult("done")

	if err := rdx.ReplaceValues(data.ArtsOperationsOrderProperty, data.ArtsOperationsOrderProperty, artsOperationsOrder...); err != nil {
		return roa.EndWithError(err)
	}

	if err := rdx.BatchReplaceValues(data.ArtsOperationsEventTimeProperty, artsOperationsEventTimes); err != nil {
		return roa.EndWithError(err)
	}

	return nil
}

func getArtsOperations(p int, kv kevlar.KeyValues) (*litres_integration.Operations, error) {

	page, err := kv.Get(strconv.Itoa(p))
	if err != nil {
		return nil, err
	}
	defer page.Close()

	var ops litres_integration.Operations
	if err := json.NewDecoder(page).Decode(&ops); err != nil {
		return nil, err
	}

	return &ops, nil
}
