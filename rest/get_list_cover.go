package rest

import (
	"github.com/beauxarts/fedorov/data"
	"net/http"
)

func GetListCover(w http.ResponseWriter, r *http.Request) {
	getCover(data.CoverSizesAsc, w, r)
}
