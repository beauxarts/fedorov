package rest

import (
	"github.com/beauxarts/litres_integration"
	"net/http"
)

func GetBookCover(w http.ResponseWriter, r *http.Request) {
	sizes := []litres_integration.CoverSize{litres_integration.SizeMax, litres_integration.Size415, litres_integration.Size330}
	getCover(sizes, w, r)
}
