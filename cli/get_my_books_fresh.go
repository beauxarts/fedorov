package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/litres_integration"
	"github.com/boggydigital/coost"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/match_node"
	"github.com/boggydigital/nod"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/http"
	"net/url"
	"strconv"
)

const maxSupportedPages = 1000

func GetMyBooksFreshHandler(u *url.URL) error {
	hc, err := coost.NewHttpClientFromFile(data.AbsCookiesFilename(), litres_integration.LitResHost)
	if err != nil {
		return err
	}

	return GetMyBooksFresh(hc)
}

func GetMyBooksFresh(hc *http.Client) error {

	gmba := nod.NewProgress("fetching my books fresh...")
	defer gmba.End()

	kv, err := kvas.ConnectLocal(data.AbsMyBooksFreshDir(), kvas.HtmlExt)
	if err != nil {
		return gmba.EndWithError(err)
	}

	// get the first page and extract total pages

	page := 1

	resp, err := hc.Get(litres_integration.MyBooksFreshUrl(page).String())
	if err != nil {
		return gmba.EndWithError(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		if err := kv.Set(strconv.Itoa(page), resp.Body); err != nil {
			return gmba.EndWithError(err)
		}
	}

	totalPages, err := getTotalPages(kv)
	gmba.TotalInt(totalPages)

	// increment to account for the first page downloaded
	gmba.Increment()

	for page = 2; page <= totalPages; page++ {

		resp, err := hc.Get(litres_integration.MyBooksFreshUrl(page).String())
		if err != nil {
			return gmba.EndWithError(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			gmba.Increment()
			continue
		}

		if err := kv.Set(strconv.Itoa(page), resp.Body); err != nil {
			return gmba.EndWithError(err)
		}

		// sleep for 5 seconds to throttle server requests
		//time.Sleep(time.Second * 5)

		gmba.Increment()
	}

	return nil
}

func getTotalPages(kv kvas.KeyValues) (int, error) {
	if fp, err := kv.Get("1"); err == nil {
		body, err := html.Parse(fp)
		if err != nil {
			return -1, err
		}

		bcEtc := match_node.NewEtc(atom.Div, "books_container mgrid_wrapper_loader_container")
		if bc := match_node.Match(body, bcEtc); bc != nil {
			for _, attr := range bc.Attr {
				if attr.Key == "data-pages" {
					i, err := strconv.ParseInt(attr.Val, 10, 32)
					return int(i), err
				}
			}
		}
		return -1, nil
	} else {
		return -1, err
	}
}
