package litres_integration

import (
	"errors"
	"github.com/boggydigital/match_node"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/http"
	"net/url"
	"strings"
)

const (
	sessionIdMarker = "getUserDataForSSR"
	sessionIdPfx    = sessionIdMarker + "(\\\\\\\""
	sessionIdSfx    = "\\\\\\\")"
)

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

	if sessionId := matchSessionId(doc); sessionId != "" {
		return sessionId, nil
	}

	return "", errors.New("no session-id found")
}

func matchSessionId(doc *html.Node) string {

	if ndsm := match_node.Match(doc, &nextDataScriptMatcher{}); ndsm != nil && ndsm.FirstChild != nil {
		nextData := ndsm.FirstChild.Data
		if strings.Contains(nextData, sessionIdMarker) {
			if _, sessionId, ok := strings.Cut(nextData, sessionIdPfx); ok {
				if parts := strings.Split(sessionId, sessionIdSfx); len(parts) > 0 {
					return parts[0]
				}
			}
		}
	}

	return ""
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
