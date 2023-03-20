package cli

import (
	"github.com/boggydigital/nod"
	"net/http"
	"net/url"
)

func PostCompletionHandler(u *url.URL) error {
	cwu := u.Query().Get("completion-webhook-url")
	return PostCompletion(cwu)
}

func PostCompletion(completionWebhookUrl string) error {

	if completionWebhookUrl == "" {
		return nil
	}

	pca := nod.Begin("posting completion...")
	defer pca.End()

	_, err := http.DefaultClient.Post(completionWebhookUrl, "", nil)
	if err != nil {
		return pca.EndWithError(err)
	}

	pca.EndWithResult("done")

	return nil
}
