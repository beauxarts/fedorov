package data

import (
	"github.com/boggydigital/pathways"
	"github.com/boggydigital/redux"
)

func NewReduxReader(assets ...string) (redux.Readable, error) {
	reduxDir, err := pathways.GetAbsRelDir(Redux)
	if err != nil {
		return nil, err
	}

	return redux.NewReader(reduxDir, assets...)
}

func NewReduxWriter(assets ...string) (redux.Writeable, error) {
	reduxDir, err := pathways.GetAbsRelDir(Redux)
	if err != nil {
		return nil, err
	}

	return redux.NewWriter(reduxDir, assets...)
}
