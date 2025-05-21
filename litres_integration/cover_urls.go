package litres_integration

import (
	"net/url"
	"strconv"
	"strings"
)

type CoverSize string

const (
	Size250 CoverSize = "_250"
	Size415 CoverSize = "_415"
	SizeMax CoverSize = ""
)

func AllCoverSizes() []CoverSize {
	return []CoverSize{
		Size250,
		Size415,
		SizeMax,
	}
}

func CoverUrl(id int64, size CoverSize) *url.URL {
	path := strings.Replace(coverPathTemplate, "{size}", string(size), 1)
	path = strings.Replace(path, "{id}", strconv.FormatInt(id, 10), 1)

	return &url.URL{
		Scheme: httpsScheme,
		Host:   cvLitResHost,
		Path:   path,
	}
}
