package rest

import (
	"github.com/beauxarts/scrinium/litres_integration"
	"net/http"
)

func GetListCover(w http.ResponseWriter, r *http.Request) {
	sizes := []litres_integration.CoverSize{litres_integration.Size330, litres_integration.Size415, litres_integration.SizeMax}
	getCover(sizes, w, r)
}
