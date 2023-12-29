package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathology"
	"github.com/boggydigital/stencil/stencil_rest"
	"net/http"
)

func GetRobotsTxt(w http.ResponseWriter, r *http.Request) {

	absInputDir, err := pathology.GetAbsDir(data.Input)
	if err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	stencil_rest.GetRobotsTxt(absInputDir, w, r)
}
