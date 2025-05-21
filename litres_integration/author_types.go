package litres_integration

type AuthorType int

const (
	AuthorUnknown AuthorType = iota
	AuthorDetails
	AuthorArts
)

var authorTypesStrings = map[AuthorType]string{
	AuthorUnknown: "author-unknown",
	AuthorDetails: "author-details",
	AuthorArts:    "author-similar",
}

func AllAuthorTypes() []AuthorType {
	aat := make([]AuthorType, 0, len(authorTypesStrings)-1)
	for at := range authorTypesStrings {
		if at == AuthorUnknown {
			continue
		}
		aat = append(aat, at)
	}

	return aat
}

func (at AuthorType) String() string {
	if ats, ok := authorTypesStrings[at]; ok {
		return ats
	}
	return authorTypesStrings[AuthorUnknown]
}

func ParseAuthorType(str string) AuthorType {
	for at, ats := range authorTypesStrings {
		if ats == str {
			return at
		}
	}
	return AuthorUnknown
}
