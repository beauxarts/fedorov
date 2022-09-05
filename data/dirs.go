package data

import (
	"path/filepath"
)

const (
	relReduxDir         = "_redux"
	relMyBookFreshDir   = "my_books_fresh"
	relMyBookDetailsDir = "my_books_details"
	relDownloadsDir     = "downloads"
	relCoversDir        = "covers"

	relCookiesFilename = "cookies.txt"
)

var rootDir = ""

func ChRoot(d string) {
	rootDir = d
}

func Pwd() string {
	return rootDir
}

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

func AbsCoverDir() string {
	return filepath.Join(rootDir, relCoversDir)
}

func AbsCookiesFilename() string {
	return filepath.Join(rootDir, relCookiesFilename)
}
