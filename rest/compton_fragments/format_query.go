package compton_fragments

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/litres_integration"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"slices"
)

func FormatQuery(q map[string][]string) map[string][]string {

	fq := make(map[string][]string)

	for p, vals := range q {
		for _, val := range vals {
			if pv, ok := compton_data.PropertyTitles[val]; ok {
				fq[p] = append(fq[p], pv)
			} else if slices.Contains(compton_data.BinaryProperties, p) {
				fq[p] = append(fq[p], compton_data.BinaryTitles[val])
			} else {
				switch p {
				case data.ArtTypeProperty:
					at := litres_integration.ParseArtType(val)
					fq[p] = append(fq[p], at.String())
				default:
					fq[p] = append(fq[p], val)
				}
			}
		}
	}
	return fq
}
