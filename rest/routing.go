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
		"/books":  Gzip(GetOnly(Log(http.HandlerFunc(GetBooks)))),
		"/book":   Gzip(GetOnly(Log(http.HandlerFunc(GetBook)))),
		"/cover":  GetOnly(Log(http.HandlerFunc(GetCover))),
		"/search": Gzip(GetOnly(Log(http.HandlerFunc(GetSearch)))),
	}

	for p, h := range patternHandlers {
		http.HandleFunc(p, h.ServeHTTP)
	}
}
