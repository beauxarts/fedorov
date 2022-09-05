package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/match_node"
	"github.com/boggydigital/nod"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func ReduceMyBooksDetails() error {

	rmbda := nod.NewProgress("reducing my books details...")
	defer rmbda.End()

	reduxProps := data.ReduxProperties()

	reductions := make(map[string]map[string][]string, len(reduxProps))
	for _, p := range reduxProps {
		reductions[p] = make(map[string][]string)
	}

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), nil, reduxProps...)
	if err != nil {
		return rmbda.EndWithError(err)
	}

	kv, err := kvas.ConnectLocal(data.AbsMyBooksDetailsDir(), kvas.HtmlExt)
	if err != nil {
		return rmbda.EndWithError(err)
	}

	ids := kv.Keys()

	rmbda.TotalInt(len(ids))

	for _, id := range ids {

		det, err := kv.Get(id)
		if err != nil {
			det.Close()
			return rmbda.EndWithError(err)
		}

		body, err := html.Parse(det)
		if err != nil {
			det.Close()
			return rmbda.EndWithError(err)
		}

		rdx := reduceDetails(body)
		for p, vals := range rdx {
			reductions[p][id] = vals
		}

		det.Close()
		rmbda.Increment()
	}

	sra := nod.NewProgress(" saving reductions...")
	defer sra.End()

	sra.TotalInt(len(reductions))

	for prop, rdx := range reductions {
		if err := rxa.BatchReplaceValues(prop, rdx); err != nil {
			return rmbda.EndWithError(err)
		}
		sra.Increment()
	}

	sra.EndWithResult("done")
	rmbda.EndWithResult("done")

	return nil
}

func reduceDetails(body *html.Node) map[string][]string {

	rdx := make(map[string][]string)

	bookNameEtc := match_node.NewEtc(atom.Div, "biblio_book_name")
	if bbn := match_node.Match(body, bookNameEtc); bbn != nil {
		for n := bbn.FirstChild; n != nil; n = n.NextSibling {
			if n.DataAtom == atom.H1 {
				rdx[data.TitleProperty] = []string{n.FirstChild.Data}
			}
		}
	}

	authorsEtc := match_node.NewEtc(atom.Div, "biblio_book_author")
	if an := match_node.Match(body, authorsEtc); an != nil {
		authorEtc := match_node.NewEtc(atom.A, "biblio_book_author__link")
		authorLinks := match_node.Matches(an, authorEtc, -1)
		authors := make([]string, 0, len(authorLinks))
		for _, al := range authorLinks {
			if span := al.FirstChild; span != nil {
				authors = append(authors, span.FirstChild.Data)
			}
		}
		rdx[data.AuthorsProperty] = authors
	}

	downloadsEtc := match_node.NewEtc(atom.Div, "book_download")
	if bdn := match_node.Match(body, downloadsEtc); bdn != nil {
		downloadLinkEtc := match_node.NewEtc(atom.A, "biblio_book_download_file__link")
		downloadLinks := match_node.Matches(bdn, downloadLinkEtc, -1)
		links := make([]string, 0, len(downloadLinks))
		for _, dl := range downloadLinks {
			links = append(links, getAttribute(dl, "href"))
		}
		rdx[data.DownloadLinksProperty] = links
	}

	if len(rdx[data.DownloadLinksProperty]) == 0 {
		// check for PDF links
		newButtonEtc := match_node.NewEtc(atom.Div, "bb_newbutton")
		for _, nb := range match_node.Matches(body, newButtonEtc, -1) {
			if getAttribute(nb, "data-type") == "pdf" {
				rdx[data.DownloadLinksProperty] = []string{getAttribute(nb.FirstChild, "href")}
			}
		}
	}

	// check for additional materials
	additionalMaterialsEtc := match_node.NewEtc(atom.Div, "bb_newbutton bb_newbutton_add-materials")
	if amn := match_node.Match(body, additionalMaterialsEtc); amn != nil {
		linkEtc := match_node.NewEtc(atom.A, "bb_newbutton_inner_link")
		if link := match_node.Match(amn, linkEtc); link != nil {
			rdx[data.DownloadLinksProperty] = append(rdx[data.DownloadLinksProperty], getAttribute(link, "href"))
		}
	}

	return rdx
}

func getAttribute(node *html.Node, attrName string) string {
	for _, attr := range node.Attr {
		if attr.Key == attrName {
			return attr.Val
		}
	}
	return ""
}
