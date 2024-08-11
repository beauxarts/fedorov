package data

import (
	"fmt"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/pathways"
	"path/filepath"
	"strconv"
)

const DefaultFedorovRootDir = "/var/lib/fedorov"

const (
	Backups   pathways.AbsDir = "backups"
	Metadata  pathways.AbsDir = "metadata"
	Input     pathways.AbsDir = "input"
	Covers    pathways.AbsDir = "covers"
	Downloads pathways.AbsDir = "downloads"
)

var AllAbsDirs = []pathways.AbsDir{
	Backups,
	Metadata,
	Input,
	Covers,
	Downloads,
}

const (
	Redux    pathways.RelDir = "_redux"
	Contents pathways.RelDir = "contents"
)

var RelToAbsDirs = map[pathways.RelDir]pathways.AbsDir{
	Redux:    Metadata,
	Contents: Metadata,
}

const (
	relCookiesFilename = "cookies.txt"

	DefaultCoverExt = ".jpg"
)

func AbsDataTypeDir(stringer fmt.Stringer) (string, error) {
	absMetadataDir, err := pathways.GetAbsDir(Metadata)
	if err != nil {
		return "", err
	}
	return filepath.Join(absMetadataDir, stringer.String()), nil
}

func absStringerDir(stringer fmt.Stringer) (string, error) {
	absMetadataDir, err := pathways.GetAbsDir(Metadata)
	if err != nil {
		return "", err
	}
	return filepath.Join(absMetadataDir, stringer.String()), nil
}

func AbsArtsTypeDir(at litres_integration.ArtsType) (string, error) {
	return absStringerDir(at)
}

func AbsSeriesTypeDir(st litres_integration.SeriesType) (string, error) {
	return absStringerDir(st)
}

func AbsAuthorTypeDir(at litres_integration.AuthorType) (string, error) {
	return absStringerDir(at)
}

func AbsFileDownloadPath(id int64, file string) (string, error) {
	absDownloadsDir, err := pathways.GetAbsDir(Downloads)
	if err != nil {
		return "", err
	}

	return filepath.Join(absDownloadsDir, strconv.FormatInt(id, 10), file), nil
}

func AbsCoverImagePath(id int64, size litres_integration.CoverSize) (string, error) {
	absCoverDir, err := pathways.GetAbsDir(Covers)
	if err != nil {
		return "", err
	}

	return filepath.Join(absCoverDir, RelCoverFilename(strconv.FormatInt(id, 10), size)), nil
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
	absInputDir, err := pathways.GetAbsDir(Input)
	if err != nil {
		return "", err
	}

	return filepath.Join(absInputDir, relCookiesFilename), nil
}
