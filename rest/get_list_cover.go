package rest

import (
	"github.com/beauxarts/fedorov/litres_integration"
	"net/http"
)

func GetListCover(w http.ResponseWriter, r *http.Request) {
	getCover(litres_integration.CoverSizesAsc, w, r)
}
