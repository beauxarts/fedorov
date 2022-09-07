package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/match_node"
	"github.com/boggydigital/nod"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/url"
)

const (
	artsPerPage = 42
)

func ReduceMyBooksFreshHandler(_ *url.URL) error {
	return ReduceMyBooksFresh()
}

func ReduceMyBooksFresh() error {

	embia := nod.NewProgress("reducing my books...")
	defer embia.End()

	kv, err := kvas.ConnectLocal(data.AbsMyBooksFreshDir(), kvas.HtmlExt)
	if err != nil {
		return embia.EndWithError(err)
	}

	keys := kv.Keys()
	myBooks := make(map[string][]string, len(keys)*artsPerPage)
	hrefs := make(map[string][]string, len(keys)*artsPerPage)

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
				id, href := idHref(img)
				myBooks[data.MyBooksIdsProperty] = append(myBooks[data.MyBooksIdsProperty], id)
				hrefs[id] = []string{href}
			}
		}
	}

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), nil,
		data.MyBooksIdsProperty,
		data.HrefProperty,
	)
	if err != nil {
		embia.EndWithError(err)
	}

	sra := nod.Begin(" saving redux...")
	defer sra.End()

	if err := rxa.BatchReplaceValues(data.MyBooksIdsProperty, myBooks); err != nil {
		embia.EndWithError(err)
	}

	if err := rxa.BatchReplaceValues(data.HrefProperty, hrefs); err != nil {
		embia.EndWithError(err)
	}

	sra.EndWithResult("done")
	embia.EndWithResult("done")

	return nil
}

func idHref(node *html.Node) (string, string) {
	id, href := "", ""
	for _, attr := range node.Attr {
		if attr.Key == "data-art" {
			id = attr.Val
		}
		if attr.Key == "href" {
			href = attr.Val
		}
		if id != "" && href != "" {
			break
		}
	}
	return id, href
}
