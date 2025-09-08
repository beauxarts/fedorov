package rest

import (
	"net/http"

	"github.com/boggydigital/middleware"
	"github.com/boggydigital/nod"
)

const (
	SharedRole = "shared"
	AdminRole  = "admin"
)

var (
	Auth = middleware.BasicHttpAuth
	Log  = nod.RequestLog
)

func HandleFuncs() {

	patternHandlers := map[string]http.Handler{
		"GET /manifest.json": Log(http.HandlerFunc(GetManifest)),
		"GET /icon.png":      Log(http.HandlerFunc(GetIcon)),

		"GET /latest": Log(http.HandlerFunc(GetLatest)),
		"GET /search": Log(http.HandlerFunc(GetSearch)),
		"GET /book":   Log(http.HandlerFunc(GetBook)),

		// book page sections
		"GET /information":    Log(http.HandlerFunc(GetInformation)),
		"GET /external-links": Log(http.HandlerFunc(GetExternalLinks)),
		"GET /files":          Log(http.HandlerFunc(GetFiles)),
		"GET /videos":         Log(http.HandlerFunc(GetVideos)),
		"GET /annotation":     Log(http.HandlerFunc(GetAnnotation)),
		"GET /similar":        Log(http.HandlerFunc(GetSimilar)),
		"GET /reviews":        Log(http.HandlerFunc(GetReviews)),
		"GET /contents":       Log(http.HandlerFunc(GetContents)),

		"GET /list_cover": Log(http.HandlerFunc(GetListCover)),
		"GET /book_cover": Log(http.HandlerFunc(GetBookCover)),

		"GET /file": Auth(Log(http.HandlerFunc(GetFile)), AdminRole, SharedRole),

		"/": http.RedirectHandler("/latest", http.StatusPermanentRedirect),
	}

	for p, h := range patternHandlers {
		http.HandleFunc(p, h.ServeHTTP)
	}
}
