package cli

import "net/http"

func addHeaders(req *http.Request, sessionId string) {
	req.Header.Set("app-id", "115")
	req.Header.Set("Session-Id", sessionId)
}
