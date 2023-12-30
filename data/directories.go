package data

import (
	"fmt"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/pathology"
	"path/filepath"
	"strconv"
)

const DefaultFedorovRootDir = "/var/lib/fedorov"

const (
	Backups   pathology.AbsDir = "backups"
	Metadata  pathology.AbsDir = "metadata"
	Input     pathology.AbsDir = "input"
	Output    pathology.AbsDir = "output"
	Covers    pathology.AbsDir = "covers"
	Downloads pathology.AbsDir = "downloads"
	Imported  pathology.AbsDir = "_imported"
)

var AllAbsDirs = []pathology.AbsDir{
	Backups,
	Metadata,
	Input,
	Output,
	Covers,
	Downloads,
	Imported,
}

const (
	Redux pathology.RelDir = "_redux"
)

var RelToAbsDirs = map[pathology.RelDir]pathology.AbsDir{
	Redux: Metadata,
}

const (
	relCookiesFilename = "cookies.txt"
	relExportFilename  = "export.txt"
	relImportFilename  = "import.txt"

	DefaultCoverExt = ".jpg"
)

func AbsDataTypeDir(stringer fmt.Stringer) (string, error) {
	absMetadataDir, err := pathology.GetAbsDir(Metadata)
	if err != nil {
		return "", err
	}
	return filepath.Join(absMetadataDir, stringer.String()), nil
}

func AbsArtsTypeDir(at litres_integration.ArtsType) (string, error) {
	absMetadataDir, err := pathology.GetAbsDir(Metadata)
	if err != nil {
		return "", err
	}
	return filepath.Join(absMetadataDir, at.String()), nil
}

func AbsFileDownloadPath(id int64, file string) (string, error) {
	absDownloadsDir, err := pathology.GetAbsDir(Downloads)
	if err != nil {
		return "", err
	}

	return filepath.Join(absDownloadsDir, strconv.FormatInt(id, 10), file), nil
}

func AbsCoverImagePath(id int64, size litres_integration.CoverSize) (string, error) {
	absCoverDir, err := pathology.GetAbsDir(Covers)
	if err != nil {
		return "", err
	}

	return filepath.Join(absCoverDir, RelCoverFilename(strconv.FormatInt(id, 10), size)), nil
}

func RelCoverFilename(id string, size litres_integration.CoverSize) string {
	switch size {
	case litres_integration.SizeMax:
		return id + DefaultCoverExt
	default:
		return fmt.Sprintf("%s_%d%s", id, size, DefaultCoverExt)
	}
}

func AbsCookiesFilename() (string, error) {
	absInputDir, err := pathology.GetAbsDir(Input)
	if err != nil {
		return "", err
	}

	return filepath.Join(absInputDir, relCookiesFilename), nil
}

func AbsExportFilename() (string, error) {
	absOutputDir, err := pathology.GetAbsDir(Output)
	if err != nil {
		return "", err
	}

	return filepath.Join(absOutputDir, relExportFilename), nil
}

func AbsImportFilename() (string, error) {
	absInputDir, err := pathology.GetAbsDir(Input)
	if err != nil {
		return "", err
	}

	return filepath.Join(absInputDir, relImportFilename), nil
}
