package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/nod"
	"golang.org/x/exp/slices"
	"net/http"
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

	DefaultHeaders(w)

	if err := rapp.RenderList("Книги", myBooks, w); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}
