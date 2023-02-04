package rest

import (
	"github.com/boggydigital/middleware"
	"github.com/boggydigital/nod"
	"net/http"
)

const (
	SharedRole = "shared"
	AdminRole  = "admin"
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
		"/list_cover":  GetOnly(Log(http.HandlerFunc(GetListCover))),
		"/book_cover":  GetOnly(Log(http.HandlerFunc(GetBookCover))),
		"/search":      Gzip(GetOnly(Log(http.HandlerFunc(GetSearch)))),
		"/digest":      Gzip(GetOnly(Log(http.HandlerFunc(GetDigest)))),
		"/downloads":   Gzip(GetOnly(Log(http.HandlerFunc(GetDownloads)))),
		"/description": Gzip(GetOnly(Log(http.HandlerFunc(GetDescription)))),
		// auth data endpoints
		"/completed/set":    Auth(Gzip(GetOnly(Log(http.HandlerFunc(GetCompletedSet)))), AdminRole),
		"/completed/clear":  Auth(Gzip(GetOnly(Log(http.HandlerFunc(GetCompletedClear)))), AdminRole),
		"/local-tags/edit":  Auth(Gzip(GetOnly(Log(http.HandlerFunc(GetLocalTagsEdit)))), AdminRole),
		"/local-tags/apply": Auth(Gzip(GetOnly(Log(http.HandlerFunc(GetLocalTagsApply)))), AdminRole),
		// auth media endpoints
		"/file": Auth(GetOnly(Log(http.HandlerFunc(GetFile))), AdminRole, SharedRole),
		// start at the books
		"/": http.RedirectHandler("/books", http.StatusPermanentRedirect),
		//robots.txt
		"/robots.txt": Gzip(GetOnly(Log(http.HandlerFunc(GetRobotsTxt)))),
	}

	for p, h := range patternHandlers {
		http.HandleFunc(p, h.ServeHTTP)
	}
}
