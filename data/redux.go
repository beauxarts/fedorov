package data

import (
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/pathways"
)

func NewReduxReader(assets ...string) (kevlar.ReadableRedux, error) {
	reduxDir, err := pathways.GetAbsRelDir(Redux)
	if err != nil {
		return nil, err
	}

	return kevlar.NewReduxReader(reduxDir, assets...)
}

func NewReduxWriter(assets ...string) (kevlar.WriteableRedux, error) {
	reduxDir, err := pathways.GetAbsRelDir(Redux)
	if err != nil {
		return nil, err
	}

	return kevlar.NewReduxWriter(reduxDir, assets...)
}
