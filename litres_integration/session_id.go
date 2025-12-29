package litres_integration

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/boggydigital/match_node"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type nextDataPageProps struct {
	Props struct {
		PageProps struct {
			InitialState string `json:"initialState"`
		} `json:"pageProps"`
	} `json:"props"`
}

type initialStateBrowser struct {
	Browser struct {
		Screen struct {
			Width  interface{} `json:"width"`
			Height interface{} `json:"height"`
		} `json:"screen"`
		UrlParams struct {
		} `json:"urlParams"`
		Pda        bool   `json:"pda"`
		ClientHost string `json:"clientHost"`
		Headers    struct {
			ClientHost      string `json:"client-host"`
			UiCurrency      string `json:"ui-currency"`
			UiLanguageCode  string `json:"Ui-Language-Code"`
			SafemodeEnabled string `json:"Safemode-Enabled"`
			XRequestId      string `json:"x-request-id"`
			Basket          string `json:"Basket"`
			Wishlist        string `json:"Wishlist"`
			SessionId       string `json:"Session-Id"`
		} `json:"headers"`
		IsRobot bool `json:"isRobot"`
	} `json:"browser"`
}

func GetSessionId(httpClient *http.Client) (string, error) {
	rootUrl := &url.URL{
		Scheme: httpsScheme,
		Host:   LitResHost,
	}

	resp, err := httpClient.Get(rootUrl.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	if sessionId, err := matchSessionId(doc); err == nil && sessionId != "" {
		return sessionId, nil
	} else if err != nil {
		return "", err
	} else {
		return "", errors.New("sessionId is empty")
	}
}

func matchSessionId(doc *html.Node) (string, error) {

	if ndsm := match_node.Match(doc, &nextDataScriptMatcher{}); ndsm != nil && ndsm.FirstChild != nil {
		nextData := ndsm.FirstChild.Data

		var ndPageProps nextDataPageProps
		if err := json.NewDecoder(strings.NewReader(nextData)).Decode(&ndPageProps); err != nil {
			return "", err
		}

		initialState := ndPageProps.Props.PageProps.InitialState

		var isBrowser initialStateBrowser
		if err := json.NewDecoder(strings.NewReader(initialState)).Decode(&isBrowser); err != nil {
			return "", err
		}

		return isBrowser.Browser.Headers.SessionId, nil
	}

	return "", errors.New("next data buildId not found")
}

type nextDataScriptMatcher struct {
}

func (ndsm *nextDataScriptMatcher) Match(node *html.Node) bool {
	if node.DataAtom == atom.Script &&
		len(node.Attr) > 0 {

		for _, attr := range node.Attr {
			if attr.Key == "id" &&
				attr.Val == "__NEXT_DATA__" {
				return true
			}
		}

	}
	return false
}
