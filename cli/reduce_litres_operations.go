package cli

import (
	"encoding/json"
	"net/url"
	"sort"
	"strconv"
	"time"

	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/redux"
)

func ReduceLitResOperationsHandler(_ *url.URL) error {
	return ReduceLitResOperations()
}

func ReduceLitResOperations() error {

	roa := nod.NewProgress("reducing operations...")
	defer roa.Done()

	absLitResOperationsDir := data.AbsDataTypeDir(litres_integration.LitResOperations)

	kv, err := kevlar.New(absLitResOperationsDir, kevlar.JsonExt)
	if err != nil {
		return err
	}

	totalPages := kv.Len()
	artsOperationsDates := make(map[string][]string, totalPages*litres_integration.OperationsLimit)
	artsFourthPresent := make(map[string][]string, totalPages*litres_integration.OperationsLimit)

	for p := 1; p <= totalPages; p++ {

		if ops, err := getArtsOperations(p, kv); err == nil {

			for _, dt := range ops.Payload.Data {
				fourthPresent := false
				if p := dt.SpecificData.Product; p != nil && *p == "4th_present" {
					fourthPresent = true
				}
				for _, art := range dt.SpecificData.Arts {
					artId := strconv.Itoa(art.Id)
					artsOperationsDates[artId] = []string{dt.Date}
					artsFourthPresent[artId] = []string{strconv.FormatBool(fourthPresent)}
				}
			}
		} else {
			return err
		}

		roa.Increment()
	}

	reduxDir := data.Pwd.AbsRelDirPath(data.Redux, data.Metadata)

	rdx, err := redux.NewWriter(reduxDir,
		data.ArtsOperationsOrderProperty,
		data.ArtsOperationsEventTimeProperty,
		data.ArtFourthPresentProperty,
		data.FreeArtsProperty)
	if err != nil {
		return err
	}

	latestArtsIds, err := getLatestOperationsFreeArts(artsOperationsDates, rdx)
	if err != nil {
		return err
	}

	sra := nod.Begin(" saving redux...")
	defer sra.Done()

	if err = rdx.ReplaceValues(data.ArtsOperationsOrderProperty, data.ArtsOperationsOrderProperty, latestArtsIds...); err != nil {
		return err
	}

	if err = rdx.BatchReplaceValues(data.ArtsOperationsEventTimeProperty, artsOperationsDates); err != nil {
		return err
	}

	if err = rdx.BatchReplaceValues(data.ArtFourthPresentProperty, artsFourthPresent); err != nil {
		return err
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

func getLatestOperationsFreeArts(artsOperationsDates map[string][]string, rdx redux.Readable) ([]string, error) {

	artsOperationsTimes := make(map[string]time.Time)

	for artId, dates := range artsOperationsDates {
		for _, date := range dates {
			if dt, err := time.Parse("2006-01-02T15:04:04", date); err != nil {
				return nil, err
			} else {
				artsOperationsTimes[artId] = dt
			}
		}
	}

	for artId := range rdx.Keys(data.FreeArtsProperty) {
		if date, ok := rdx.GetLastVal(data.FreeArtsProperty, artId); ok && date != "" {
			if dt, err := time.Parse(time.RFC3339, date); err != nil {
				return nil, err
			} else {
				artsOperationsTimes[artId] = dt
			}
		}
	}

	latestArtsIds := make([]string, 0, len(artsOperationsTimes))

	for artId := range artsOperationsTimes {
		latestArtsIds = append(latestArtsIds, artId)
	}

	sort.SliceStable(latestArtsIds, func(i, j int) bool {
		return artsOperationsTimes[latestArtsIds[i]].After(artsOperationsTimes[latestArtsIds[j]])
	})

	return latestArtsIds, nil
}
