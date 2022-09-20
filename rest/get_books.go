package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/stencil_app"
	"github.com/boggydigital/nod"
	"golang.org/x/exp/slices"
	"net/http"
	"strconv"
	"time"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {

	// GET /books

	var err error
	if rxa, err = rxa.RefreshReduxAssets(); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	myBooks, ok := rxa.GetAllUnchangedValues(data.MyBooksIdsProperty, data.MyBooksIdsProperty)
	if !ok {
		http.Error(w, nod.ErrorStr("no my books found"), http.StatusInternalServerError)
		return
	}

	if missingDetails, ok := rxa.GetAllUnchangedValues(data.MissingDetailsIdsProperty, data.MissingDetailsIdsProperty); ok {
		filteredBooks := make([]string, 0, len(myBooks))
		for _, id := range myBooks {
			if slices.Contains(missingDetails, id) {
				continue
			}
			filteredBooks = append(filteredBooks, id)
		}
		myBooks = filteredBooks
	}

	booksByType := make(map[string][]string)

	for _, id := range myBooks {
		bt, _ := rxa.GetFirstVal(data.BookTypeProperty, id)
		booksByType[bt] = append(booksByType[bt], id)
	}

	DefaultHeaders(w)

	updated := "recently"
	if scu, ok := rxa.GetFirstVal(data.SyncCompletedProperty, data.SyncCompletedProperty); ok {
		if scui, err := strconv.ParseInt(scu, 10, 64); err == nil {
			updated = time.Unix(scui, 0).Format(time.RFC1123)
		}
	}

	if err := rapp.RenderGroup(
		"Книги",
		stencil_app.BookTypeOrder,
		booksByType,
		stencil_app.BookTypeTitles,
		updated,
		w); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}
