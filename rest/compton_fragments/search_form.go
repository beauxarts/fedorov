package compton_fragments

import (
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/input_types"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/kevlar"
	"strings"
)

func SearchForm(r compton.Registrar, query map[string][]string, searchQuery compton.Element, rdx kevlar.ReadableRedux) compton.Element {

	form := compton.Form("/search", "GET")
	formStack := compton.FlexItems(r, direction.Column)
	form.Append(formStack)

	if searchQuery != nil {
		formStack.Append(searchQuery)
	}

	submitRow := compton.FlexItems(r, direction.Row).JustifyContent(align.Center)
	submit := compton.InputValue(r, input_types.Submit, "Искать")
	submitRow.Append(submit)
	formStack.Append(submitRow)

	inputsGrid := compton.GridItems(r).JustifyContent(align.Center).GridTemplateRows(size.XLarge)
	formStack.Append(inputsGrid)

	searchInputs(r, query, inputsGrid, rdx)

	// duplicating Submit button after inputs at the end
	formStack.Append(submitRow)

	return form
}

func searchInputs(r compton.Registrar, query map[string][]string, container compton.Element, rdx kevlar.ReadableRedux) {
	for _, property := range compton_data.SearchProperties {
		title := compton_data.PropertyTitles[property]
		value := strings.Join(query[property], ", ")
		titleInput := compton.TISearchValue(r, title, property, value)

		var datalist map[string]string
		var listId string

		//if slices.Contains(compton_data.BinaryDigestProperties, property) {
		//	datalist = binDatalist
		//	listId = "bin-list"
		//} else if slices.Contains(compton_data.DigestProperties, property) {
		//	switch property {
		//	case vangogh_local_data.TypesProperty:
		//		datalist = typesDatalist()
		//	case vangogh_local_data.OperatingSystemsProperty:
		//		datalist = operatingSystemsDatalist()
		//	case vangogh_local_data.SortProperty:
		//		datalist = sortDatalist()
		//	case vangogh_local_data.ProductTypeProperty:
		//		datalist = productTypesDatalist()
		//	case vangogh_local_data.SteamDeckAppCompatibilityCategoryProperty:
		//		datalist = steamDeckDatalist()
		//	case vangogh_local_data.LanguageCodeProperty:
		//		datalist = languagesDatalist()
		//	case vangogh_local_data.TagIdProperty:
		//		datalist = tagsDatalist(rdx)
		//	}
		//}

		if len(datalist) > 0 {
			titleInput.SetDatalist(datalist, listId)
		}

		container.Append(titleInput)
	}
}
