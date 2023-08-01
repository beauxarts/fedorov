package rest

import (
	"github.com/beauxarts/fedorov/data"
	"net/http"
)

func GetBookCover(w http.ResponseWriter, r *http.Request) {
	getCover(data.CoverSizesDesc, w, r)
}
