package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/stencil/stencil_rest"
	"net/http"
)

func GetRobotsTxt(w http.ResponseWriter, r *http.Request) {
	stencil_rest.GetRobotsTxt(data.Pwd(), w, r)
}
