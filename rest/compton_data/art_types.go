package compton_data

import (
	"github.com/beauxarts/scrinium/litres_integration"
	"strconv"
)

var ArtTypes = map[string]string{
	strconv.Itoa(int(litres_integration.ArtTypeText)):  "Текст",
	strconv.Itoa(int(litres_integration.ArtTypeAudio)): "Аудио",
	strconv.Itoa(int(litres_integration.ArtTypePDF)):   "PDF",
}
