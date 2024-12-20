package cli

import (
	"encoding/json"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/nod"
	"net/url"
	"strconv"
	"time"
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

		if order, eventTimes, err := artsOperationsOrderEventTimes(p, kv); err == nil {

			artsOperationsOrder = append(artsOperationsOrder, order...)
			for artId, ets := range eventTimes {
				artsOperationsEventTimes[artId] = ets
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

func artsOperationsOrderEventTimes(p int, kv kevlar.KeyValues) ([]string, map[string][]string, error) {

	page, err := kv.Get(strconv.Itoa(p))
	if err != nil {
		return nil, nil, err
	}
	defer page.Close()

	var ops litres_integration.Operations
	if err := json.NewDecoder(page).Decode(&ops); err != nil {
		return nil, nil, err
	}

	order := make([]string, 0, len(ops.Payload.Data)*3)
	eventTimes := make(map[string][]string)

	for _, dt := range ops.Payload.Data {
		et, err := time.Parse("2006-01-02T15:04:05", dt.Date)
		if err != nil {
			return nil, nil, err
		}
		for _, art := range dt.SpecificData.Arts {
			artId := strconv.Itoa(art.Id)
			order = append(order, artId)
			eventTimes[artId] = []string{et.Format(time.RFC3339)}
		}
	}

	return order, eventTimes, nil
}
