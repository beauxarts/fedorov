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
	Auth = middleware.BasicHttpAuth
	Log  = nod.RequestLog
)

var port int

func HandleFuncs(p int) {

	port = p

	patternHandlers := map[string]http.Handler{
		// unauth data endpoints
		"GET /books":       Log(http.HandlerFunc(GetBooks)),
		"GET /book":        Log(http.HandlerFunc(GetBook)),
		"GET /list_cover":  Log(http.HandlerFunc(GetListCover)),
		"GET /book_cover":  Log(http.HandlerFunc(GetBookCover)),
		"GET /search":      Log(http.HandlerFunc(GetSearch)),
		"GET /digest":      Log(http.HandlerFunc(GetDigest)),
		"GET /downloads":   Log(http.HandlerFunc(GetDownloads)),
		"GET /description": Log(http.HandlerFunc(GetDescription)),
		// auth data endpoints
		"GET /completed/set":    Auth(Log(http.HandlerFunc(GetCompletedSet)), AdminRole),
		"GET /completed/clear":  Auth(Log(http.HandlerFunc(GetCompletedClear)), AdminRole),
		"GET /local-tags/edit":  Auth(Log(http.HandlerFunc(GetLocalTagsEdit)), AdminRole),
		"GET /local-tags/apply": Auth(Log(http.HandlerFunc(GetLocalTagsApply)), AdminRole),
		// auth media endpoints
		"GET /file": Auth(Log(http.HandlerFunc(GetFile)), AdminRole, SharedRole),
		// start at the books
		"/": http.RedirectHandler("/books", http.StatusPermanentRedirect),
		//robots.txt
		"GET /robots.txt": Log(http.HandlerFunc(GetRobotsTxt)),
	}

	for p, h := range patternHandlers {
		http.HandleFunc(p, h.ServeHTTP)
	}
}
