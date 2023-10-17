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
	Auth     = middleware.BasicHttpAuth
	BrGzip   = middleware.BrGzip
	GetOnly  = middleware.GetMethodOnly
	PostOnly = middleware.PostMethodOnly
	Static   = middleware.Static
	Log      = nod.RequestLog
)

var port int

func HandleFuncs(p int) {

	port = p

	patternHandlers := map[string]http.Handler{
		// unauth data endpoints
		"/books":       BrGzip(GetOnly(Static(Log(http.HandlerFunc(GetBooks))))),
		"/book":        BrGzip(GetOnly(Static(Log(http.HandlerFunc(GetBook))))),
		"/list_cover":  GetOnly(Log(http.HandlerFunc(GetListCover))),
		"/book_cover":  GetOnly(Log(http.HandlerFunc(GetBookCover))),
		"/search":      BrGzip(GetOnly(Static(Log(http.HandlerFunc(GetSearch))))),
		"/digest":      BrGzip(GetOnly(Log(http.HandlerFunc(GetDigest)))),
		"/downloads":   BrGzip(GetOnly(Log(http.HandlerFunc(GetDownloads)))),
		"/description": BrGzip(GetOnly(Log(http.HandlerFunc(GetDescription)))),
		// auth data endpoints
		"/completed/set":    Auth(BrGzip(GetOnly(Log(http.HandlerFunc(GetCompletedSet)))), AdminRole),
		"/completed/clear":  Auth(BrGzip(GetOnly(Log(http.HandlerFunc(GetCompletedClear)))), AdminRole),
		"/local-tags/edit":  Auth(BrGzip(GetOnly(Log(http.HandlerFunc(GetLocalTagsEdit)))), AdminRole),
		"/local-tags/apply": Auth(BrGzip(GetOnly(Log(http.HandlerFunc(GetLocalTagsApply)))), AdminRole),
		// auth media endpoints
		"/file": Auth(GetOnly(Log(http.HandlerFunc(GetFile))), AdminRole, SharedRole),
		// prerender
		"/prerender": PostOnly(Log(http.HandlerFunc(PostPrerender))),
		// start at the books
		"/": http.RedirectHandler("/books", http.StatusPermanentRedirect),
		//robots.txt
		"/robots.txt": BrGzip(GetOnly(Log(http.HandlerFunc(GetRobotsTxt)))),
	}

	for p, h := range patternHandlers {
		http.HandleFunc(p, h.ServeHTTP)
	}
}
