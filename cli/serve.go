package cli

import (
	"fmt"
	"github.com/beauxarts/fedorov/rest"
	"github.com/boggydigital/nod"
	"net/http"
	"net/url"
	"strconv"
)

func ServeHandler(u *url.URL) error {

	username := u.Query().Get("username")
	password := u.Query().Get("password")

	if username != "" && password != "" {
		rest.SetUsername(username)
		rest.SetPassword(password)
	}

	portstr := u.Query().Get("port")
	if port, err := strconv.ParseInt(portstr, 10, 32); err == nil {
		stderr := u.Query().Has("stderr")
		return Serve(int(port), stderr)
	} else {
		return err
	}
}

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
