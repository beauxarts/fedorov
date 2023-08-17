package cli

import (
	"github.com/boggydigital/nod"
	"net/http"
	"net/url"
)

func PostCompletionHandler(u *url.URL) error {
	wu := u.Query().Get("webhook-url")
	return PostCompletion(wu)
}

func PostCompletion(webhookUrl string) error {

	pca := nod.Begin("posting completion...")
	defer pca.End()

	if webhookUrl == "" {
		pca.EndWithResult("webhook-url is empty")
		return nil
	}

	_, err := http.DefaultClient.Post(webhookUrl, "", nil)
	if err != nil {
		return pca.EndWithError(err)
	}

	pca.EndWithResult("done")

	return nil
}
