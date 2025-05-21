package litres_integration

type ArtsType int

const (
	ArtsTypeUnknown ArtsType = iota
	ArtsTypeDetails
	ArtsTypeSimilar
	ArtsTypeQuotes
	ArtsTypeFiles
	ArtsTypeReviews
)

var artsTypesStrings = map[ArtsType]string{
	ArtsTypeUnknown: "arts-unknown",
	ArtsTypeDetails: "arts-details",
	ArtsTypeSimilar: "arts-similar",
	ArtsTypeQuotes:  "arts-quotes",
	ArtsTypeFiles:   "arts-files",
	ArtsTypeReviews: "arts-reviews",
}

func AllArtsTypes() []ArtsType {
	aat := make([]ArtsType, 0, len(artsTypesStrings)-1)
	for at := range artsTypesStrings {
		if at == ArtsTypeUnknown {
			continue
		}
		aat = append(aat, at)
	}

	return aat
}

func (at ArtsType) String() string {
	if ats, ok := artsTypesStrings[at]; ok {
		return ats
	}
	return artsTypesStrings[ArtsTypeUnknown]
}

func ParseArtsType(str string) ArtsType {
	for at, ats := range artsTypesStrings {
		if ats == str {
			return at
		}
	}
	return ArtsTypeUnknown
}
