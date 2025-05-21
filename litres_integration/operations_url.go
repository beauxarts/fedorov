package litres_integration

import (
	"net/url"
	"strconv"
)

const (
	OperationsLimit = 50
)

func OperationsUrl(page int) *url.URL {
	u := &url.URL{
		Scheme: httpsScheme,
		Host:   apiLitResHost,
		Path:   operationsPath,
	}

	q := u.Query()

	q.Set("limit", strconv.Itoa(OperationsLimit))
	if page > 1 {
		q.Set("offset", strconv.Itoa((page-1)*OperationsLimit))
	}

	u.RawQuery = q.Encode()

	return u
}
