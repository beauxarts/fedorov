package compton_pages

import (
	"encoding/xml"
	"os"

	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/redux"
)

func Contents(id string, rdx redux.Readable) compton.PageElement {

	s := compton_fragments.ProductSection(compton_data.ContentsSection, id, rdx)

	contentsDir := data.Pwd.AbsRelDirPath(data.Contents, data.Metadata)

	contReader, err := kevlar.New(contentsDir, kevlar.XmlExt)
	if err != nil {
		s.Error(err)
		return s
	}

	var contents *litres_integration.Contents

	contXml, err := contReader.Get(id)
	if err != nil {
		if !os.IsNotExist(err) {
			s.Error(err)
			return s
		}
	} else {
		if contXml != nil {
			if err = xml.NewDecoder(contXml).Decode(&contents); err != nil {
				s.Error(err)
				return s
			}
		}
	}

	if contentsElement := compton_fragments.Contents(s, contents); contents != nil {
		s.Append(contentsElement)
	}
	return s
}
