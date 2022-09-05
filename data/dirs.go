package data

import (
	"path/filepath"
	"strconv"
)

const (
	relReduxDir         = "_redux"
	relDetailedDataDir  = "detailed_data"
	relMyBookFreshDir   = "my_books_fresh"
	relMyBookDetailsDir = "my_books_details"
	relDownloadsDir     = "downloads"
	relCoversDir        = "covers"

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

//func AbsDetailedDataDir() string {
//	return filepath.Join(rootDir, relDetailedDataDir)
//}

func AbsMyBooksFreshDir() string {
	return filepath.Join(rootDir, relMyBookFreshDir)
}

func AbsMyBooksDetailsDir() string {
	return filepath.Join(rootDir, relMyBookDetailsDir)
}

func AbsDownloadsDir() string {
	return filepath.Join(rootDir, relDownloadsDir)
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
