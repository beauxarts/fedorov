package data

import (
	"path/filepath"
	"strconv"
)

const (
	relReduxDir              = "_redux"
	relDetailedDataRemoteDir = "detailed_data_remote"
	relDetailedDataLocalDir  = "detailed_data_local"
	relMyBookFreshDir        = "my_books_fresh"
	relCoversDir             = "covers"

	relCookiesFilename = "cookies.txt"

	coverExt = ".jpg"
)

var rootDir = ""

func ChRoot(d string) {
	rootDir = d
}

func Pwd() string {
	return rootDir
}
func AbsDetailedDataRemoteDir() string {
	return filepath.Join(rootDir, relDetailedDataRemoteDir)
}

func AbsDetailedDataLocalDir() string {
	return filepath.Join(rootDir, relDetailedDataLocalDir)
}

func AbsMyBooksFreshDir() string {
	return filepath.Join(rootDir, relMyBookFreshDir)
}

func AbsReduxDir() string {
	return filepath.Join(rootDir, relReduxDir)
}

func AbsCoverPath(id int64) string {
	dir := filepath.Join(rootDir, relCoversDir)

	if idstr := strconv.FormatInt(id, 10); len(idstr) > 0 {
		return filepath.Join(dir, idstr[:0], idstr+coverExt)
	}

	return ""
}

func AbsCookiesFilename() string {
	return filepath.Join(rootDir, relCookiesFilename)
}
