package data

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/boggydigital/camino"
)

const (
	directoriesFilename = "directories.txt"
	fedorovRootDir      = "/var/lib/fedorov"
)

const (
	Backups camino.AbsDir = iota
	Metadata
	Input
	Covers
	Downloads
)

var absDirNames = map[camino.AbsDir]string{
	Backups:   "backups",
	Metadata:  "metadata",
	Input:     "input",
	Covers:    "covers",
	Downloads: "downloads",
}

const (
	Redux camino.RelDir = iota
	Contents
)

var relDirNames = map[camino.RelDir]string{
	Redux:    "_redux",
	Contents: "contents",
}

var relAbsParents = map[camino.RelDir][]camino.AbsDir{
	Redux:    {Metadata},
	Contents: {Metadata},
}

const (
	relCookiesFilename = "cookies_litres_ru.json"
	DefaultCoverExt    = ".jpg"
)

func AbsDataTypeDir(stringer fmt.Stringer) string {
	return filepath.Join(camino.GetAbs(Metadata), stringer.String())
}

func absStringerDir(stringer fmt.Stringer) string {
	return filepath.Join(camino.GetAbs(Metadata), stringer.String())
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
	return filepath.Join(camino.GetAbs(Downloads), strconv.FormatInt(id, 10), file)
}

func AbsCoverImagePath(id int64, size litres_integration.CoverSize) string {
	return filepath.Join(camino.GetAbs(Covers), RelCoverFilename(strconv.FormatInt(id, 10), size))
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
	return filepath.Join(camino.GetAbs(Input), relCookiesFilename), nil
}

func InitFedorovCamino() error {

	var overrides map[string]string

	if _, err := os.Stat(directoriesFilename); err == nil {
		if overrides, err = camino.ReadOverrides(directoriesFilename); err != nil {
			return err
		}
	}

	fads := make(map[camino.AbsDir]string)

	for ad := range absDirNames {
		var apd string
		var ok bool

		if apd, ok = absDirNames[ad]; !ok {
			return errors.New("fedorov abs dir path not set")
		}

		var dir string
		if dir, ok = overrides[apd]; !ok {
			fads[ad] = filepath.Join(fedorovRootDir, apd)
		} else {
			fads[ad] = dir
		}
	}

	vrds := make(map[camino.RelDir]string)

	for vrp := range relAbsParents {
		var ok bool
		if vrds[vrp], ok = relDirNames[vrp]; !ok {
			return errors.New("fedorov rel dir path not set")
		}
	}

	return camino.Register(fads, relDirNames, relAbsParents)
}
