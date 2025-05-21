package cli

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/nod"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func GetLitResOperationsHandler(u *url.URL) error {
	sessionId := u.Query().Get("session-id")

	return GetLitResOperations(sessionId, nil)
}

func GetLitResOperations(sessionId string, hc *http.Client) error {
	goa := nod.Begin("fetching LitRes operations...")
	defer goa.Done()

	absLitResOperationsDir, err := data.AbsDataTypeDir(litres_integration.LitResOperations)
	if err != nil {
		return err
	}

	kv, err := kevlar.New(absLitResOperationsDir, kevlar.JsonExt)
	if err != nil {
		return err
	}

	if hc == nil {
		hc, err = getHttpClient()
		if err != nil {
			return err
		}
	}

	page := 1
	hasMore := true

	for hasMore {
		if hasMore, err = getOperationsPage(page, sessionId, hc, kv); err != nil {
			return err
		}
		page++
	}

	return nil
}

func getOperationsPage(page int, sessionId string, hc *http.Client, kv kevlar.KeyValues) (bool, error) {

	req, err := http.NewRequest(http.MethodGet,
		litres_integration.OperationsUrl(page).String(), nil)
	if err != nil {
		return false, err
	}

	addHeaders(req, sessionId)

	resp, err := hc.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, errors.New(resp.Status)
	}

	var bts []byte
	buf := bytes.NewBuffer(bts)

	tr := io.TeeReader(resp.Body, buf)

	if err := kv.Set(strconv.Itoa(page), tr); err != nil {
		return false, err
	}

	var ops litres_integration.Operations

	if err := json.NewDecoder(buf).Decode(&ops); err != nil {
		return false, err
	}

	return ops.Payload.Pagination.NextPage != nil, nil
}
