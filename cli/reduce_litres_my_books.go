package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/match_node"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pasu"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/url"
	"sort"
	"strconv"
)

const (
	artsPerPage = 42
)

func ReduceLitResMyBooksHandler(_ *url.URL) error {
	return ReduceLitResMyBooks()
}

func ReduceLitResMyBooks() error {

	embia := nod.NewProgress("reducing my books...")
	defer embia.End()

	absLitResMyBooksFreshDir, err := data.AbsDataTypeDir(litres_integration.LitResMyBooksFresh)
	if err != nil {
		return embia.EndWithError(err)
	}

	kv, err := kvas.ConnectLocal(absLitResMyBooksFreshDir, kvas.HtmlExt)
	if err != nil {
		return embia.EndWithError(err)
	}

	keys := kv.Keys()
	myBooks := make(map[string][]string, len(keys)*artsPerPage)
	hrefs := make(map[string][]string, len(keys)*artsPerPage)

	// sort my-books keys (page numbers) before iterating through them
	// to enforce last bought - shown at the top order
	iks := make([]int64, 0, len(keys))
	for _, k := range keys {
		if ik, err := strconv.ParseInt(k, 10, 64); err == nil {
			iks = append(iks, ik)
		}
	}
	sort.Slice(iks, func(i, j int) bool { return iks[i] < iks[j] })

	for _, ik := range iks {

		page, err := kv.Get(strconv.FormatInt(ik, 10))
		if err != nil {
			return embia.EndWithError(err)
		}

		body, err := html.Parse(page)
		if err != nil {
			return embia.EndWithError(err)
		}

		bcEtc := match_node.NewEtc(atom.Div, "books_container mgrid_wrapper_loader_container", true)
		if bc := match_node.Match(body, bcEtc); bc != nil {
			imgEtc := match_node.NewEtc(atom.A, "img-a", true)
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

	absReduxDir, err := pasu.GetAbsRelDir(data.Redux)
	if err != nil {
		return embia.EndWithError(err)
	}

	rdx, err := kvas.NewReduxWriter(absReduxDir,
		data.MyBooksIdsProperty,
		data.HrefProperty,
		data.ImportedProperty)
	if err != nil {
		return embia.EndWithError(err)
	}

	sra := nod.Begin(" saving redux...")
	defer sra.End()

	// add previously imported book to my books before saving that set
	myBooks[data.MyBooksIdsProperty] = append(myBooks[data.MyBooksIdsProperty], rdx.Keys(data.ImportedProperty)...)

	if err := rdx.BatchReplaceValues(data.MyBooksIdsProperty, myBooks); err != nil {
		return embia.EndWithError(err)
	}

	if err := rdx.BatchReplaceValues(data.HrefProperty, hrefs); err != nil {
		return embia.EndWithError(err)
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
