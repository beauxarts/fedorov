package compton_fragments

import (
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/size"
)

func GitHubLink(r compton.Registrar) compton.Element {
	gitHubLink := compton.A("https://github.com/beauxarts")
	gitHubLink.Append(compton.Fspan(r, "東京からこんにちは").FontWeight(font_weight.Bolder).FontSize(size.XSmall))
	return gitHubLink
}
