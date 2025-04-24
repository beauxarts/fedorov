package compton_fragments

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/input_types"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/redux"
	"golang.org/x/net/html/atom"
	"slices"
	"strings"
)

func SearchForm(r compton.Registrar, query map[string][]string, searchQuery *compton.FrowElement, rdx redux.Readable) compton.Element {

	form := compton.Form("/search", "GET")
	formStack := compton.FlexItems(r, direction.Column)
	form.Append(formStack)

	if searchQuery != nil {
		formStack.Append(compton.FICenter(r, searchQuery))
	}

	submitRow := compton.FlexItems(r, direction.Row).JustifyContent(align.Center)
	submit := compton.InputValue(r, input_types.Submit, "Искать")
	submitRow.Append(submit)
	formStack.Append(submitRow)

	inputsGrid := compton.GridItems(r).JustifyContent(align.Center).RowGap(size.Normal)
	formStack.Append(inputsGrid)

	searchInputs(r, query, inputsGrid, rdx)

	// duplicating Submit button after inputs at the end
	formStack.Append(submitRow)

	return form
}

func searchInputs(r compton.Registrar, query map[string][]string, container compton.Element, rdx redux.Readable) {
	for ii, property := range compton_data.SearchProperties {
		title := compton_data.PropertyTitles[property]
		value := strings.Join(query[property], ", ")
		titleInput := compton.TISearchValue(r, title, property, value)
		titleInput.RowGap(size.XSmall)

		if ii == 0 {
			if input := titleInput.GetFirstElementByTagName(atom.Input); input != nil {
				input.SetAttribute("autofocus", "")
			}
		}

		var datalist map[string]string
		var listId string

		if slices.Contains(compton_data.BinaryProperties, property) {
			datalist = binDatalist
			listId = "bin-list"
		} else if slices.Contains(compton_data.DigestProperties, property) {
			switch property {
			case data.ArtTypeProperty:
				datalist = make(map[string]string)
				for _, ats := range []string{
					litres_integration.ArtTypeText.String(),
					litres_integration.ArtTypeAudio.String(),
					litres_integration.ArtTypePDF.String(),
				} {
					datalist[ats] = ats
				}
			case data.SortProperty:
				datalist = sortDatalist
			default:
				datalist = propertyDatalist(property, rdx)
			}
		}

		if len(datalist) > 0 {
			titleInput.SetDatalist(datalist, listId)
		}

		container.Append(titleInput)
	}
}

var binDatalist = map[string]string{
	"true":  "Да",
	"false": "Нет",
}

var sortDatalist = map[string]string{
	data.ArtsOperationsOrderProperty: compton_data.PropertyTitles[data.ArtsOperationsEventTimeProperty],
}

func propertyDatalist(property string, rdx redux.Readable) map[string]string {
	values := make(map[string]string)
	for id := range rdx.Keys(property) {
		if vals, ok := rdx.GetAllValues(property, id); ok {
			for _, val := range vals {
				values[val] = val
			}
		}
	}
	return values
}
