package rest

import (
	"github.com/boggydigital/middleware"
	"github.com/boggydigital/nod"
	"net/http"
)

var (
	Gzip    = middleware.Gzip
	GetOnly = middleware.GetMethodOnly
	Log     = nod.RequestLog
)

func HandleFuncs() {
	patternHandlers := map[string]http.Handler{
		"/books":     Gzip(GetOnly(Log(http.HandlerFunc(GetBooks)))),
		"/book":      Gzip(GetOnly(Log(http.HandlerFunc(GetBook)))),
		"/cover":     GetOnly(Log(http.HandlerFunc(GetCover))),
		"/search":    Gzip(GetOnly(Log(http.HandlerFunc(GetSearch)))),
		"/downloads": Gzip(GetOnly(Log(http.HandlerFunc(GetDownloads)))),
		"/file":      GetOnly(Log(http.HandlerFunc(GetFile))),
		// root redirect
		"/": http.RedirectHandler("/books", http.StatusPermanentRedirect),
	}

	for p, h := range patternHandlers {
		http.HandleFunc(p, h.ServeHTTP)
	}
}
