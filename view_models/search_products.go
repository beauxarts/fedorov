package view_models

import "github.com/boggydigital/kvas"

type searchProducts struct {
	SearchProperties []string
	Query            map[string]string
	Digests          map[string][]string
	DigestsTitles    map[string]string
	Shelf            *Shelf
}

func NewSearchProducts(ids []string, rxa kvas.ReduxAssets) *searchProducts {
	return &searchProducts{
		SearchProperties: SearchProperties,
		Query:            make(map[string]string),
		//DigestsTitles:    digestTitles,
		Shelf: NewShelf(ids, rxa),
	}
}
