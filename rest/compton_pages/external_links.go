package compton_pages

import (
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/kevlar"
)

func ExternalLinks(id string, rdx kevlar.ReadableRedux) compton.PageElement {
	s := compton_fragments.ProductSection(compton_data.FilesSection)
	if links := compton_fragments.ExternalLinks(s, externalLinks(id, rdx)); links != nil {
		s.Append(links)
	}
	return s
}

func externalLinks(id string, rdx kevlar.ReadableRedux) map[string][]string {

	links := make(map[string][]string)

	//rdx.GetLastVal(data.Url)

	return links
}
