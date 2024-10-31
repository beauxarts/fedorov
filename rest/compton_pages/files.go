package compton_pages

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/direction"
)

func Files(id string) compton.PageElement {
	s := compton_fragments.ProductSection(compton_data.FilesSection)

	filesReader, err := data.NewArtsReader(litres_integration.ArtsTypeFiles)
	if err != nil {
		return s.Error(err)
	}

	artFiles, err := filesReader.ArtsFiles(id)
	if err != nil {
		return s.Error(err)
	}

	pageStack := compton.FlexItems(s, direction.Column)
	s.Append(pageStack)

	dlTypes := artFiles.PreferredDownloadsTypes()
	for ii, dt := range dlTypes {
		if dte := compton_fragments.DownloadType(s, id, dt); dte != nil {
			pageStack.Append(dte)
		}
		if ii != len(dlTypes)-1 {
			pageStack.Append(compton.Hr())
		}
	}

	return s
}
