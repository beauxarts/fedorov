package data

import (
	"fmt"
	"github.com/beauxarts/scrinium/litres_integration"
	"path/filepath"
	"strconv"
)

const (
	relImportedDir       = "_imported"
	relBackupDir         = "backup"
	relMyBookFreshDir    = "my_books_fresh"
	relMyBookDetailsDir  = "my_books_details"
	relLiveLibDetailsDir = "livelib_details"
	relDownloadsDir      = "downloads"

	relCookiesFilename = "cookies.txt"
	relExportFilename  = "export.txt"
	relImportFilename  = "import.txt"

	CoverExt = ".jpg"
)

var (
	absRootDir   = ""
	absReduxDir  = ""
	absCoversDir = ""
)

func ChRoot(d string) {
	absRootDir = d
}

func Pwd() string {
	return absRootDir
}

func SetReduxDir(d string) {
	absReduxDir = d
}

func SetCoversDir(d string) {
	absCoversDir = d
}

func AbsLitResMyBooksFreshDir() string {
	return filepath.Join(absRootDir, relMyBookFreshDir)
}

func AbsLitResMyBooksDetailsDir() string {
	return filepath.Join(absRootDir, relMyBookDetailsDir)
}

func AbsLiveLibDetailsDir() string {
	return filepath.Join(absRootDir, relLiveLibDetailsDir)
}

func AbsDownloadsDir() string {
	return filepath.Join(absRootDir, relDownloadsDir)
}

func AbsDownloadPath(id int64, file string) string {
	return filepath.Join(AbsDownloadsDir(), strconv.FormatInt(id, 10), file)
}

func AbsReduxDir() string {
	return absReduxDir
}

func AbsCoverPath(id int64, size litres_integration.CoverSize) string {
	return filepath.Join(AbsCoverDir(), RelCoverFilename(strconv.FormatInt(id, 10), size))
}

func AbsCoverDir() string {
	return absCoversDir
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
	return filepath.Join(absRootDir, relCookiesFilename)
}

func AbsExportFilename() string {
	return filepath.Join(absRootDir, relExportFilename)
}

func AbsImportFilename() string {
	return filepath.Join(absRootDir, relImportFilename)
}

func AbsImportedDir() string {
	return filepath.Join(absRootDir, relImportedDir)
}

func AbsBackupDir() string {
	return filepath.Join(absRootDir, relBackupDir)
}
