package cli

import (
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/litres_integration"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

func ReduceDetailedData(since int64) error {

	rdda := nod.Begin("reducing detailed data...")
	defer rdda.End()

	ddrp := filepath.Join(data.AbsDetailedDataRemoteDir(), litres_integration.DetailedDataFilename)
	stat, err := os.Stat(ddrp)
	if err != nil {
		return rdda.EndWithError(err)
	}

	if stat.ModTime().Unix() < since {
		// no changes since provided time
		rdda.EndWithResult("unchanged")
		return nil
	}

	lxa := nod.NewProgress(" loading data...")
	defer lxa.End()

	ddf, err := os.Open(ddrp)
	if err != nil {
		return rdda.EndWithError(err)
	}
	defer ddf.Close()

	tr := io.TeeReader(ddf, lxa)
	lxa.Total(uint64(stat.Size()))

	gr, err := gzip.NewReader(tr)
	if err != nil {
		return rdda.EndWithError(err)
	}

	var lru litres_integration.LitResUpdates

	if err := xml.NewDecoder(gr).Decode(&lru); err != nil {
		return rdda.EndWithError(err)
	}

	lxa.EndWithResult("done")

	properties := []string{
		data.ArtTitlesProperty,
		data.ArtAuthorsProperty,
		data.ArtGenresProperty,
		data.ArtPublisherProperty,
		data.ArtISBNPropertyProperty,
		data.ArtSrcUrlsProperty,
		data.ArtYearProperty,
		data.ArtSequenceNameProperty,
		data.ArtSequenceNumberProperty,
		data.ArtRelationsProperty,
	}

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), nil, properties...)
	if err != nil {
		rdda.EndWithError(err)
	}

	rdx := make(map[string]map[string][]string, len(properties))
	for _, p := range properties {
		rdx[p] = make(map[string][]string, len(lru.Arts))
	}

	raa := nod.NewProgress(" reducing arts...")
	defer raa.End()

	raa.TotalInt(len(lru.Arts))

	for _, art := range lru.Arts {
		id := strconv.FormatInt(art.IntId, 10)
		for prop, vals := range reduceArt(&art) {
			rdx[prop][id] = vals
		}
		raa.Increment()
	}

	sra := nod.NewProgress(" saving reductions...")
	defer sra.End()

	sra.TotalInt(len(rdx))

	for prop, reds := range rdx {
		if err := rxa.BatchReplaceValues(prop, reds); err != nil {
			return rdda.EndWithError(err)
		}
		sra.Increment()
	}

	sra.EndWithResult("done")

	return nil
}

func reduceArt(art *litres_integration.Art) map[string][]string {

	pi := art.TextDescription.Hidden.PublishInfo
	di := art.TextDescription.Hidden.DocumentInfo
	ti := art.TextDescription.Hidden.TitleInfo

	rdx := map[string][]string{
		data.ArtTitlesProperty:         {pi.BookName},
		data.ArtGenresProperty:         ti.Genres,
		data.ArtPublisherProperty:      {pi.Publisher},
		data.ArtISBNPropertyProperty:   {pi.ISBN},
		data.ArtSrcUrlsProperty:        di.SrcUrls,
		data.ArtYearProperty:           {strconv.Itoa(pi.Year)},
		data.ArtSequenceNameProperty:   {pi.Sequence.Name},
		data.ArtSequenceNumberProperty: {strconv.Itoa(pi.Sequence.Number)},
	}

	authors := make([]string, 0, len(art.Authors.Authors))
	for _, author := range art.Authors.Authors {
		authors = append(authors, fmt.Sprintf("%s %s", author.FirstName, author.LastName))
	}
	rdx[data.ArtAuthorsProperty] = authors

	relations := make([]string, 0, len(art.ArtsRelations.ArtsRelations))
	for _, rel := range art.ArtsRelations.ArtsRelations {
		relations = append(relations, strconv.FormatInt(rel.RelArt, 10))
	}
	rdx[data.ArtRelationsProperty] = relations

	return rdx
}
