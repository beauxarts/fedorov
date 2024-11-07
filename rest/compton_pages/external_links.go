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
	if ats, ok := rdx.GetLastVal(data.ArtTypeProperty, id); ok {
		at := litres_integration.ParseArtType(ats)
		switch at {
		case litres_integration.ArtTypeText:
			altType = litres_integration.ArtTypeAudio.String()
		case litres_integration.ArtTypeAudio:
			altType = litres_integration.ArtTypeText.String()
		}
	}

	links[data.LitresBookLinksProperty] = append(links[data.LitresBookLinksProperty],
		encodeLink(bookType, "/book/"+id))

	if avId, ok := rdx.GetLastVal(data.AlternativeVersionsProperty, id); ok && avId != "0" {
		links[data.LitresBookLinksProperty] = append(links[data.LitresBookLinksProperty],
			encodeLink(altType, "/book/"+avId))
	}

	appendLink(links, id, rdx,
		data.PersonsIdsProperty,
		data.PersonFullNameProperty,
		data.PersonUrlProperty,
		data.LitresAuthorLinksProperty)

	appendLink(links, id, rdx,
		data.SeriesIdProperty,
		data.SeriesNameProperty,
		data.SeriesUrlProperty,
		data.LitresSeriesLinksProperty)

	appendLink(links, id, rdx,
		data.PublisherIdProperty,
		data.PublisherNameProperty,
		data.PublisherUrlProperty,
		data.LitresPublishersLinksProperty)

	appendLink(links, id, rdx,
		data.RightholdersIdsProperty,
		data.RightholderNameProperty,
		data.RightholderUrlProperty,
		data.LitresRightholdersLinksProperty)

	appendLink(links, id, rdx,
		data.GenresIdsProperty,
		data.GenreNameProperty,
		data.GenreUrlProperty,
		data.LitresGenresLinksProperty)

	appendLink(links, id, rdx,
		data.TagsIdsProperty,
		data.TagNameProperty,
		data.TagUrlProperty,
		data.LitresTagsLinksProperty)

	return links
}

func appendLink(links map[string][]string, id string, rdx kevlar.ReadableRedux, idsProperty, nameProperty, urlProperty, linkProperty string) {
	if pids, ok := rdx.GetAllValues(idsProperty, id); ok {
		for _, pid := range pids {
			name, _ := rdx.GetLastVal(nameProperty, pid)
			path, _ := rdx.GetLastVal(urlProperty, pid)

			links[linkProperty] = append(links[linkProperty], encodeLink(name, path))
		}
	}
}

func encodeLink(name, path string) string {

	absUrl := url.URL{
		Scheme: "https",
		Host:   litres_integration.LitResHost,
		Path:   path,
	}

	return fmt.Sprintf("%s=%s", name, absUrl.String())
}
