package compton_pages

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/kevlar"
	"net/url"
	"strconv"
)

var (
	textType  = strconv.Itoa(int(litres_integration.ArtTypeText))
	audioType = strconv.Itoa(int(litres_integration.ArtTypeAudio))
)

func ExternalLinks(id string, rdx kevlar.ReadableRedux) compton.PageElement {
	s := compton_fragments.ProductSection(compton_data.ExternalLinksSection)
	if links := compton_fragments.ExternalLinks(s, externalLinks(id, rdx)); links != nil {
		s.Append(links)
	}
	return s
}

func externalLinks(id string, rdx kevlar.ReadableRedux) map[string][]string {

	links := make(map[string][]string)

	var bookType, altType string
	if at, ok := rdx.GetLastVal(data.ArtTypeProperty, id); ok {
		bookType = compton_data.ArtTypes[at]
		switch at {
		case textType:
			altType = compton_data.ArtTypes[audioType]
		case audioType:
			altType = compton_data.ArtTypes[textType]
		}
	}

	links[data.LitresBookLinksProperty] = append(links[data.LitresBookLinksProperty],
		encodeLink(bookType, "/book/"+id))

	if avId, ok := rdx.GetLastVal(data.AlternativeVersionsProperty, id); ok && avId != "0" {
		links[data.LitresBookLinksProperty] = append(links[data.LitresBookLinksProperty],
			encodeLink(altType, "/book/"+avId))
	}

	if personsIds, ok := rdx.GetAllValues(data.PersonsIdsProperty, id); ok {
		for _, personId := range personsIds {
			name, _ := rdx.GetLastVal(data.PersonFullNameProperty, personId)
			path, _ := rdx.GetLastVal(data.PersonUrlProperty, personId)
			links[data.LitresAuthorLinksProperty] = append(links[data.LitresAuthorLinksProperty],
				encodeLink(name, path))
		}
	}

	if seriesIds, ok := rdx.GetAllValues(data.SeriesIdProperty, id); ok {
		for _, seriesId := range seriesIds {
			name, _ := rdx.GetLastVal(data.SeriesNameProperty, seriesId)
			path, _ := rdx.GetLastVal(data.SeriesUrlProperty, seriesId)
			links[data.LitresSeriesLinksProperty] = append(links[data.LitresSeriesLinksProperty],
				encodeLink(name, path))
		}
	}

	if publishersIds, ok := rdx.GetAllValues(data.PublisherIdProperty, id); ok {
		for _, publisherId := range publishersIds {
			name, _ := rdx.GetLastVal(data.PublisherNameProperty, publisherId)
			path, _ := rdx.GetLastVal(data.PublisherUrlProperty, publisherId)
			links[data.LitresPublishersLinksProperty] = append(links[data.LitresPublishersLinksProperty],
				encodeLink(name, path))
		}
	}

	if rightholdersIds, ok := rdx.GetAllValues(data.RightholdersIdsProperty, id); ok {
		for _, rightholderId := range rightholdersIds {
			name, _ := rdx.GetLastVal(data.RightholderNameProperty, rightholderId)
			path, _ := rdx.GetLastVal(data.RightholderUrlProperty, rightholderId)
			links[data.LitresRightholdersLinksProperty] = append(links[data.LitresRightholdersLinksProperty],
				encodeLink(name, path))
		}
	}

	if genresIds, ok := rdx.GetAllValues(data.GenresIdsProperty, id); ok {
		for _, genreId := range genresIds {
			name, _ := rdx.GetLastVal(data.GenreNameProperty, genreId)
			path, _ := rdx.GetLastVal(data.GenreUrlProperty, genreId)
			links[data.LitresGenresLinksProperty] = append(links[data.LitresGenresLinksProperty],
				encodeLink(name, path))
		}
	}

	if tagsIds, ok := rdx.GetAllValues(data.TagsIdsProperty, id); ok {
		for _, tagId := range tagsIds {
			name, _ := rdx.GetLastVal(data.TagNameProperty, tagId)
			path, _ := rdx.GetLastVal(data.TagUrlProperty, tagId)
			links[data.LitresTagsLinksProperty] = append(links[data.LitresTagsLinksProperty],
				encodeLink(name, path))
		}
	}

	return links
}

func encodeLink(name, path string) string {

	absUrl := url.URL{
		Scheme: "https",
		Host:   litres_integration.LitResHost,
		Path:   path,
	}

	return fmt.Sprintf("%s=%s", name, absUrl.String())
}
