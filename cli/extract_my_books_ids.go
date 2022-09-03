package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/match_node"
	"github.com/boggydigital/nod"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const (
	artsPerPage = 42
)

func ExtractMyBooksIds() error {

	embia := nod.NewProgress("extracting my books ids...")
	defer embia.End()

	kv, err := kvas.ConnectLocal(data.AbsMyBooksFreshDir(), kvas.HtmlExt)
	if err != nil {
		return embia.EndWithError(err)
	}

	keys := kv.Keys()
	myBooksIds := make([]string, 0, len(keys)*artsPerPage)

	for _, key := range keys {

		page, err := kv.Get(key)
		if err != nil {
			return embia.EndWithError(err)
		}

		body, err := html.Parse(page)
		if err != nil {
			return embia.EndWithError(err)
		}

		bcEtc := match_node.NewEtc(atom.Div, "books_container mgrid_wrapper_loader_container")
		if bc := match_node.Match(body, bcEtc); bc != nil {
			imgEtc := match_node.NewEtc(atom.A, "img-a")
			for _, img := range match_node.Matches(bc, imgEtc, -1) {
				if img == nil {
					continue
				}
				for _, attr := range img.Attr {
					if attr.Key == "data-art" {
						myBooksIds = append(myBooksIds, attr.Val)
					}
				}
			}
		}
	}

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), nil, data.MyBooksIdsProperty)
	if err != nil {
		embia.EndWithError(err)
	}

	if err := rxa.ReplaceValues(data.MyBooksIdsProperty, data.MyBooksIdsProperty, myBooksIds...); err != nil {
		embia.EndWithError(err)
	}

	embia.EndWithResult("done")

	return nil
}
