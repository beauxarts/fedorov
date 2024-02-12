package data

import (
	"fmt"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/pasu"
	"path/filepath"
	"strconv"
)

const DefaultFedorovRootDir = "/var/lib/fedorov"

const (
	Backups   pasu.AbsDir = "backups"
	Metadata  pasu.AbsDir = "metadata"
	Input     pasu.AbsDir = "input"
	Output    pasu.AbsDir = "output"
	Covers    pasu.AbsDir = "covers"
	Downloads pasu.AbsDir = "downloads"
	Imported  pasu.AbsDir = "_imported"
)

var AllAbsDirs = []pasu.AbsDir{
	Backups,
	Metadata,
	Input,
	Output,
	Covers,
	Downloads,
	Imported,
}

const (
	Redux    pasu.RelDir = "_redux"
	Contents pasu.RelDir = "contents"
)

var RelToAbsDirs = map[pasu.RelDir]pasu.AbsDir{
	Redux:    Metadata,
	Contents: Metadata,
}

const (
	relCookiesFilename = "cookies.txt"
	relExportFilename  = "export.txt"
	relImportFilename  = "import.txt"

	DefaultCoverExt = ".jpg"
)

func AbsDataTypeDir(stringer fmt.Stringer) (string, error) {
	absMetadataDir, err := pasu.GetAbsDir(Metadata)
	if err != nil {
		return "", err
	}
	return filepath.Join(absMetadataDir, stringer.String()), nil
}

func absStringerDir(stringer fmt.Stringer) (string, error) {
	absMetadataDir, err := pasu.GetAbsDir(Metadata)
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
	absDownloadsDir, err := pasu.GetAbsDir(Downloads)
	if err != nil {
		return "", err
	}

	return filepath.Join(absDownloadsDir, strconv.FormatInt(id, 10), file), nil
}

func AbsCoverImagePath(id int64, size litres_integration.CoverSize) (string, error) {
	absCoverDir, err := pasu.GetAbsDir(Covers)
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
	absInputDir, err := pasu.GetAbsDir(Input)
	if err != nil {
		return "", err
	}

	return filepath.Join(absInputDir, relCookiesFilename), nil
}

func AbsExportFilename() (string, error) {
	absOutputDir, err := pasu.GetAbsDir(Output)
	if err != nil {
		return "", err
	}

	return filepath.Join(absOutputDir, relExportFilename), nil
}

func AbsImportFilename() (string, error) {
	absInputDir, err := pasu.GetAbsDir(Input)
	if err != nil {
		return "", err
	}

	return filepath.Join(absInputDir, relImportFilename), nil
}
