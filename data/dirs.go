package data

import (
	"fmt"
	"github.com/beauxarts/litres_integration"
	"path/filepath"
	"strconv"
)

const (
	relReduxDir         = "_redux"
	relImportedDir      = "_imported"
	relBackupDir        = "backup"
	relMyBookFreshDir   = "my_books_fresh"
	relMyBookDetailsDir = "my_books_details"
	relDownloadsDir     = "downloads"
	relCoversDir        = "covers"

	relCookiesFilename = "cookies.txt"
	relExportFilename  = "export.txt"
	relImportFilename  = "import.txt"

	CoverExt = ".jpg"
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

func AbsDownloadPath(id int64, file string) string {
	return filepath.Join(AbsDownloadsDir(), strconv.FormatInt(id, 10), file)
}

func AbsReduxDir() string {
	return filepath.Join(rootDir, relReduxDir)
}

func AbsCoverPath(id int64, size litres_integration.CoverSize) string {
	return filepath.Join(AbsCoverDir(), RelCoverFilename(strconv.FormatInt(id, 10), size))
}

func AbsCoverDir() string {
	return filepath.Join(rootDir, relCoversDir)
}

func RelCoverFilename(id string, size litres_integration.CoverSize) string {
	switch size {
	case litres_integration.SizeMax:
		return id + CoverExt
	default:
		return fmt.Sprintf("%s_%d%s", id, size, CoverExt)
	}
}

func AbsCookiesFilename() string {
	return filepath.Join(rootDir, relCookiesFilename)
}

func AbsExportFilename() string {
	return filepath.Join(rootDir, relExportFilename)
}

func AbsImportFilename() string {
	return filepath.Join(rootDir, relImportFilename)
}

func AbsImportedDir() string {
	return filepath.Join(rootDir, relImportedDir)
}

func AbsBackupDir() string {
	return filepath.Join(rootDir, relBackupDir)
}
