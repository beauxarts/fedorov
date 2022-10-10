package rest

import (
	"github.com/boggydigital/middleware"
	"github.com/boggydigital/nod"
	"net/http"
)

var (
	Auth    = middleware.BasicHttpAuth
	Gzip    = middleware.Gzip
	GetOnly = middleware.GetMethodOnly
	Log     = nod.RequestLog
)

func HandleFuncs() {
	patternHandlers := map[string]http.Handler{
		// unauth data endpoints
		"/books":       Gzip(GetOnly(Log(http.HandlerFunc(GetBooks)))),
		"/book":        Gzip(GetOnly(Log(http.HandlerFunc(GetBook)))),
		"/cover":       GetOnly(Log(http.HandlerFunc(GetCover))),
		"/search":      Gzip(GetOnly(Log(http.HandlerFunc(GetSearch)))),
		"/downloads":   Gzip(GetOnly(Log(http.HandlerFunc(GetDownloads)))),
		"/description": Gzip(GetOnly(Log(http.HandlerFunc(GetDescription)))),
		// auth data endpoints
		"/completed/set":    Auth(Gzip(GetOnly(Log(http.HandlerFunc(GetCompletedSet))))),
		"/completed/clear":  Auth(Gzip(GetOnly(Log(http.HandlerFunc(GetCompletedClear))))),
		"/local-tags/edit":  Auth(Gzip(GetOnly(Log(http.HandlerFunc(GetLocalTagsEdit))))),
		"/local-tags/apply": Auth(Gzip(GetOnly(Log(http.HandlerFunc(GetLocalTagsApply))))),
		// auth media endpoints
		"/file": Auth(GetOnly(Log(http.HandlerFunc(GetFile)))),
		// start at the books
		"/": http.RedirectHandler("/books", http.StatusPermanentRedirect),
	}

	for p, h := range patternHandlers {
		http.HandleFunc(p, h.ServeHTTP)
	}
}
