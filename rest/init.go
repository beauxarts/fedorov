package rest

import (
	"crypto/sha256"
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/middleware"
	"github.com/boggydigital/pathways"
	"github.com/boggydigital/redux"
)

var (
	rdx redux.Readable
)

func SetUsername(role, u string) {
	middleware.SetUsername(role, sha256.Sum256([]byte(u)))
}

func SetPassword(role, p string) {
	middleware.SetPassword(role, sha256.Sum256([]byte(p)))
}

func Init() error {

	var err error

	reduxDir, err := pathways.GetAbsRelDir(data.Redux)
	if err != nil {
		return err
	}

	if rdx, err = redux.NewReader(reduxDir, data.ReduxProperties()...); err != nil {
		return err
	}

	return err
}
