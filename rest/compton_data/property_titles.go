package compton_data

import "github.com/beauxarts/fedorov/data"

var PropertyTitles = map[string]string{
	data.AuthorsProperty:      "Автор",
	data.ReadersProperty:      "Чтец",
	data.TranslatorsProperty:  "Перевод",
	data.IllustratorsProperty: "Иллюстратор",
	data.PerformersProperty:   "Исполнитель",
	data.PaintersProperty:     "Художник",
	data.PublishersProperty:   "Издатель",

	data.DateWrittenAtProperty: "Написано",
}

var ShortPropertyTitles = map[string]string{
	data.IllustratorsProperty: "Иллюстр.",
}
