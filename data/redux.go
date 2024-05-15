package data

import (
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/pathways"
)

func NewReduxReader(assets ...string) (kvas.ReadableRedux, error) {
	reduxDir, err := pathways.GetAbsRelDir(Redux)
	if err != nil {
		return nil, err
	}

	return kvas.NewReduxReader(reduxDir, assets...)
}

func NewReduxWriter(assets ...string) (kvas.WriteableRedux, error) {
	reduxDir, err := pathways.GetAbsRelDir(Redux)
	if err != nil {
		return nil, err
	}

	return kvas.NewReduxWriter(reduxDir, assets...)
}
