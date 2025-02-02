package data

import (
	"encoding/json"
	"fmt"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/kevlar"
	"io"
	"iter"
)

type ArtsReader struct {
	artsType  litres_integration.ArtsType
	keyValues kevlar.KeyValues
}

func NewArtsReader(at litres_integration.ArtsType) (*ArtsReader, error) {

	absArtsTypeDir, err := AbsArtsTypeDir(at)
	if err != nil {
		return nil, err
	}

	atr := &ArtsReader{
		artsType: at,
	}

	atr.keyValues, err = kevlar.New(absArtsTypeDir, kevlar.JsonExt)
	if err != nil {
		return nil, err
	}

	return atr, nil
}

func (ar *ArtsReader) Keys() iter.Seq[string] {
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

func (ar *ArtsReader) Cut(id string) error {
	return ar.keyValues.Cut(id)
}

func (ar *ArtsReader) Since(ts int64, mts ...kevlar.MutationType) iter.Seq2[string, kevlar.MutationType] {
	return ar.keyValues.Since(ts, mts...)
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

func (ar *ArtsReader) LogModTime(id string) int64 {
	return ar.keyValues.LogModTime(id)
}

func (ar *ArtsReader) FileModTime(id string) (int64, error) {
	return ar.keyValues.FileModTime(id)
}

func (ar *ArtsReader) Len() int { return ar.keyValues.Len() }
