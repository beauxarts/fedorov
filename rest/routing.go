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
		"/books": middleware.Gzip(GetOnly(Log(http.HandlerFunc(GetBooks)))),
	}

	for p, h := range patternHandlers {
		http.HandleFunc(p, h.ServeHTTP)
	}
}
