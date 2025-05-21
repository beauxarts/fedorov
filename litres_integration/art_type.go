package litres_integration

import "strconv"

type ArtType int

const (
	ArtTypeText  ArtType = 0
	ArtTypeAudio ArtType = 1
	ArtTypePDF   ArtType = 4
)

var artTypeStrings = map[ArtType]string{
	ArtTypeText:  "Текст",
	ArtTypeAudio: "Аудио",
	ArtTypePDF:   "PDF",
}

var artTypeValues = map[ArtType]string{
	ArtTypeText:  strconv.Itoa(int(ArtTypeText)),
	ArtTypeAudio: strconv.Itoa(int(ArtTypeAudio)),
	ArtTypePDF:   strconv.Itoa(int(ArtTypePDF)),
}

func (at ArtType) String() string {
	return artTypeStrings[at]
}

func ParseArtType(ats string) ArtType {
	for at, str := range artTypeValues {
		if str == ats {
			return at
		}
	}
	return ArtTypeText
}
