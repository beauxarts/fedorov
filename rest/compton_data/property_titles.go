package compton_data

import "github.com/beauxarts/fedorov/data"

var PropertyTitles = map[string]string{
	data.TitleProperty: "Название",

	data.AuthorsProperty:      "Автор",
	data.ReadersProperty:      "Чтец",
	data.TranslatorsProperty:  "Перевод",
	data.IllustratorsProperty: "Иллюстратор",
	data.PerformersProperty:   "Исполнитель",
	data.PaintersProperty:     "Художник",
	data.PublishersProperty:   "Издатель",

	data.ArtTypeProperty:       "Тип",
	data.DateWrittenAtProperty: "Написано",
	data.MinAgeProperty:        "Возрастное ограничение",

	data.SortProperty:             "Порядок",
	data.ArtsHistoryOrderProperty: "По дате приобретения",
	data.DescendingProperty:       "По убыванию",

	data.BookCompletedProperty: "Прочитано",
	data.LocalTagsProperty:     "Свои тэги",
}

var ShortPropertyTitles = map[string]string{
	data.IllustratorsProperty: "Иллюстр.",
}
