package rest

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/nod"
	"net/http"
	"os"
	"strconv"
)

func getCover(sizes []litres_integration.CoverSize, w http.ResponseWriter, r *http.Request) {

	idstr := r.URL.Query().Get("id")

	if idstr == "" {
		http.Error(w, nod.ErrorStr("missing required book id"), http.StatusInternalServerError)
		return
	}

	if id, err := strconv.ParseInt(idstr, 10, 64); err == nil {
		for _, size := range sizes {
			coverPath := data.AbsCoverPath(id, size)
			if _, err := os.Stat(coverPath); err == nil {
				w.Header().Set("Cache-Control", "max-age=31536000")
				http.ServeFile(w, r, coverPath)
				return
			}
		}
		_ = nod.Error(fmt.Errorf("no cover for id %s", idstr))
		http.NotFound(w, r)
	} else {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

}
