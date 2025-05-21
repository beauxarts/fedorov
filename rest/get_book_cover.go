package rest

import (
	"github.com/beauxarts/fedorov/litres_integration"
	"net/http"
)

func GetBookCover(w http.ResponseWriter, r *http.Request) {
	getCover(litres_integration.CoverSizesDesc, w, r)
}
