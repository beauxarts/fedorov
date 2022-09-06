package cli

import (
	"fmt"
	"github.com/beauxarts/fedorov/rest"
	"github.com/boggydigital/nod"
	"net/http"
)

func Serve(port int, stderr bool) error {

	if stderr {
		nod.EnableStdErrLogger()
		nod.DisableOutput(nod.StdOut)
	}

	sa := nod.Begin("serving at port %d...", port)
	defer sa.End()

	if err := rest.Init(); err != nil {
		return err
	}

	rest.HandleFuncs()

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}