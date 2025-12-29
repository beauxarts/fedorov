package data

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/boggydigital/pathways"
)

const (
	setPathwaysFilename = "directories.txt"
	rootPathwaysDir     = "/var/lib/fedorov"
)

const (
	Backups   pathways.AbsDir = "backups"
	Metadata  pathways.AbsDir = "metadata"
	Input     pathways.AbsDir = "input"
	Covers    pathways.AbsDir = "covers"
	Downloads pathways.AbsDir = "downloads"
)

const (
	Redux    pathways.RelDir = "_redux"   // Metadata
	Contents pathways.RelDir = "contents" // Metadata
)

var Pwd pathways.Pathway

const (
	relCookiesFilename = "cookies_litres_ru.json"
	DefaultCoverExt    = ".jpg"
)

func AbsDataTypeDir(stringer fmt.Stringer) string {
	return filepath.Join(Pwd.AbsDirPath(Metadata), stringer.String())
}

func absStringerDir(stringer fmt.Stringer) string {
	return filepath.Join(Pwd.AbsDirPath(Metadata), stringer.String())
}

func AbsArtsTypeDir(at litres_integration.ArtsType) string {
	return absStringerDir(at)
}

func AbsSeriesTypeDir(st litres_integration.SeriesType) string {
	return absStringerDir(st)
}

func AbsAuthorTypeDir(at litres_integration.AuthorType) string {
	return absStringerDir(at)
}

func AbsFileDownloadPath(id int64, file string) string {
	return filepath.Join(Pwd.AbsDirPath(Downloads), strconv.FormatInt(id, 10), file)
}

func AbsCoverImagePath(id int64, size litres_integration.CoverSize) string {
	return filepath.Join(Pwd.AbsDirPath(Covers), RelCoverFilename(strconv.FormatInt(id, 10), size))
}

func RelCoverFilename(id string, size litres_integration.CoverSize) string {
	fn := ""
	switch size {
	case litres_integration.SizeMax:
		fn = id + DefaultCoverExt
	default:
		fn = fmt.Sprintf("%s%s%s", id, size, DefaultCoverExt)
	}

	if len(fn) > 0 {
		fn = filepath.Join(fn[:2], fn)
	}

	return fn
}

func AbsCookiesFilename() (string, error) {
	return filepath.Join(Pwd.AbsDirPath(Input), relCookiesFilename), nil
}

func InitPathways() error {

	var setExists bool
	if _, err := os.Stat(setPathwaysFilename); err == nil {
		setExists = true
	}

	var err error
	switch setExists {
	case true:
		Pwd, err = pathways.ReadSet(setPathwaysFilename)
	default:
		Pwd, err = pathways.NewRoot(rootPathwaysDir)
	}

	return err
}
