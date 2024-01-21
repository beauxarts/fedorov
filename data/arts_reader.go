package data

import (
	"encoding/json"
	"fmt"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"io"
)

type ArtsReader struct {
	artsType  litres_integration.ArtsType
	keyValues kvas.KeyValues
}

func NewArtsReader(at litres_integration.ArtsType) (*ArtsReader, error) {

	absArtsTypeDir, err := AbsArtsTypeDir(at)
	if err != nil {
		return nil, err
	}

	atr := &ArtsReader{
		artsType: at,
	}

	atr.keyValues, err = kvas.ConnectLocal(absArtsTypeDir, kvas.JsonExt)
	if err != nil {
		return nil, err
	}

	return atr, nil
}

func (ar *ArtsReader) Keys() []string {
	return ar.keyValues.Keys()
}

func (ar *ArtsReader) Has(id string) bool { return ar.keyValues.Has(id) }

func (ar *ArtsReader) Get(id string) (io.ReadCloser, error) { return ar.keyValues.Get(id) }

func (ar *ArtsReader) GetFromStorage(id string) (io.ReadCloser, error) {
	return ar.keyValues.Get(id)
}

func (ar *ArtsReader) Set(id string, data io.Reader) error {
	return ar.keyValues.Set(id, data)
}

func (ar *ArtsReader) Cut(id string) (bool, error) {
	return ar.keyValues.Cut(id)
}

func (ar *ArtsReader) CreatedAfter(timestamp int64) []string {
	return ar.keyValues.CreatedAfter(timestamp)
}

func (ar *ArtsReader) ModifiedAfter(timestamp int64, excludeCreated bool) []string {
	return ar.keyValues.ModifiedAfter(timestamp, excludeCreated)
}

func (ar *ArtsReader) IsModifiedAfter(id string, timestamp int64) bool {
	return ar.keyValues.IsModifiedAfter(id, timestamp)
}

func (ar *ArtsReader) readValue(id string, val interface{}) error {
	spReadCloser, err := ar.keyValues.Get(id)
	if err != nil {
		return err
	}

	if spReadCloser == nil {
		return nil
	}

	defer spReadCloser.Close()

	if err := json.NewDecoder(spReadCloser).Decode(val); err != nil {
		return err
	}

	return nil
}

func (ar *ArtsReader) ArtsDetails(id string) (artsDetails *litres_integration.ArtsDetails, err error) {
	err = ar.readValue(id, &artsDetails)
	return artsDetails, err
}

func (ar *ArtsReader) ArtsSimilar(id string) (artsSimilar *litres_integration.ArtsSimilar, err error) {
	err = ar.readValue(id, &artsSimilar)
	return artsSimilar, err
}

func (ar *ArtsReader) ArtsQuotes(id string) (artsQuotes *litres_integration.ArtsQuotes, err error) {
	err = ar.readValue(id, &artsQuotes)
	return artsQuotes, err
}

func (ar *ArtsReader) ArtsFiles(id string) (artsFiles *litres_integration.ArtsFiles, err error) {
	err = ar.readValue(id, &artsFiles)
	return artsFiles, err
}

func (ar *ArtsReader) ArtsReviews(id string) (artsReviews *litres_integration.ArtsReviews, err error) {
	err = ar.readValue(id, &artsReviews)
	return artsReviews, err
}

func (ar *ArtsReader) ReadValue(id string) (interface{}, error) {
	switch ar.artsType {
	case litres_integration.ArtsTypeDetails:
		return ar.ArtsDetails(id)
	case litres_integration.ArtsTypeSimilar:
		return ar.ArtsSimilar(id)
	case litres_integration.ArtsTypeQuotes:
		return ar.ArtsQuotes(id)
	case litres_integration.ArtsTypeFiles:
		return ar.ArtsFiles(id)
	case litres_integration.ArtsTypeReviews:
		return ar.ArtsReviews(id)
	default:
		return nil, fmt.Errorf("cannot create %s value", ar.artsType)
	}
}

func (ar *ArtsReader) ArtsType() litres_integration.ArtsType {
	return ar.artsType
}

func (ar *ArtsReader) IndexCurrentModTime() (int64, error) {
	return ar.keyValues.IndexCurrentModTime()
}

func (ar *ArtsReader) CurrentModTime(id string) (int64, error) {
	return ar.keyValues.CurrentModTime(id)
}

func (ar *ArtsReader) IndexRefresh() error {
	return ar.keyValues.IndexRefresh()
}

func (ar *ArtsReader) VetIndexOnly(fix bool, tpw nod.TotalProgressWriter) ([]string, error) {
	return ar.keyValues.VetIndexOnly(fix, tpw)
}

func (ar *ArtsReader) VetIndexMissing(fix bool, tpw nod.TotalProgressWriter) ([]string, error) {
	return ar.keyValues.VetIndexMissing(fix, tpw)
}
