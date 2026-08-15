package litres_integration

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/boggydigital/camino"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const getUserDataForSsr = "getUserDataForSSR"

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

	var sessionId string
	if sessionId, err = matchSessionId(doc); err == nil && sessionId != "" {
		return sessionId, nil
	} else if err != nil {
		return "", err
	} else {
		return "", errors.New("sessionId is empty")
	}
}

func matchSessionId(doc *html.Node) (string, error) {

	if ndsm := camino.Match(doc, &getUserDataForSsrScriptMatcher{}); ndsm != nil && ndsm.FirstChild != nil {
		if _, firstPass, ok := strings.Cut(ndsm.FirstChild.Data, getUserDataForSsr); ok {
			if secondPass, _, sure := strings.Cut(firstPass, ":"); sure {
				sessionId := strings.Trim(secondPass, "(\\\")")
				return sessionId, nil
			}
		}
	}

	return "", errors.New("next data buildId not found")
}

type getUserDataForSsrScriptMatcher struct {
}

func (sm *getUserDataForSsrScriptMatcher) Match(node *html.Node) bool {
	if node.DataAtom == atom.Script &&
		node.FirstChild != nil &&
		len(node.FirstChild.Data) > 0 {

		return strings.Contains(node.FirstChild.Data, getUserDataForSsr)

	}
	return false
}
