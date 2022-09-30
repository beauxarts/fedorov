package rest

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/nod"
	"net/http"
	"os"
	"strconv"
)

func GetCover(w http.ResponseWriter, r *http.Request) {

	// GET /cover?id

	idstr := r.URL.Query().Get("id")

	if idstr == "" {
		http.Error(w, nod.ErrorStr("missing required book id"), http.StatusInternalServerError)
		return
	}

	if id, err := strconv.ParseInt(idstr, 10, 64); err == nil {
		coverPath := data.AbsCoverPath(id)
		if _, err := os.Stat(coverPath); err == nil {
			w.Header().Set("Cache-Control", "max-age=31536000")
			http.ServeFile(w, r, coverPath)
		} else {
			_ = nod.Error(fmt.Errorf("no cover for id %s", idstr))
			http.NotFound(w, r)
		}
	} else {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

}
