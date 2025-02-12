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

	sharedUsername := u.Query().Get("shared-username")
	sharedPassword := u.Query().Get("shared-password")
	adminUsername := u.Query().Get("admin-username")
	adminPassword := u.Query().Get("admin-password")

	if sharedUsername != "" && sharedPassword != "" {
		rest.SetUsername(rest.SharedRole, sharedUsername)
		rest.SetPassword(rest.SharedRole, sharedPassword)
	}

	if adminUsername != "" && adminPassword != "" {
		rest.SetUsername(rest.AdminRole, adminUsername)
		rest.SetPassword(rest.AdminRole, adminPassword)
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
	defer sa.Done()

	if err := rest.Init(); err != nil {
		return err
	}

	rest.HandleFuncs(port)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
