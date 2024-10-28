package compton_fragments

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/input_types"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/kevlar"
	"golang.org/x/exp/slices"
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

		if slices.Contains(compton_data.BinaryProperties, property) {
			datalist = binDatalist
			listId = "bin-list"
		} else if slices.Contains(compton_data.DigestProperties, property) {
			switch property {
			case data.ArtTypeProperty:
				datalist = compton_data.ArtTypes
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
	data.ArtsHistoryOrderProperty: compton_data.PropertyTitles[data.ArtsHistoryOrderProperty],
}

func propertyDatalist(property string, rdx kevlar.ReadableRedux) map[string]string {
	values := make(map[string]string)
	for _, id := range rdx.Keys(property) {
		if vals, ok := rdx.GetAllValues(property, id); ok {
			for _, val := range vals {
				values[val] = val
			}
		}
	}
	return values
}
