package compton_data

import "github.com/beauxarts/fedorov/data"

var PropertyTitles = map[string]string{
	data.TitleProperty:          "Название",
	data.HTMLAnnotationProperty: "Аннотация",

	data.AuthorsProperty:      "Автор",
	data.ReadersProperty:      "Читает",
	data.TranslatorsProperty:  "Перевод",
	data.IllustratorsProperty: "Иллюстратор",
	data.PerformersProperty:   "Исполнитель",
	data.PaintersProperty:     "Художник",

	data.ArtTypeProperty:      "Тип",
	data.GenresProperty:       "Жанры",
	data.TagsProperty:         "Тэги",
	data.LocalTagsProperty:    "Свои тэги",
	data.PublishersProperty:   "Издатель",
	data.RightholdersProperty: "Правообладатель",
	data.MinAgeProperty:       "Возрастное ограничение",
	data.ISBNProperty:         "ISBN",
	data.SeriesProperty:       "Серия",

	data.DateWrittenAtProperty:   "Написано",
	data.FirstTimeSaleAtProperty: "Начало продаж",
	data.PublicationDateProperty: "Опубликовано",
	data.RegisteredAtProperty:    "Зарегистрировано",
	data.TranslatedAtProperty:    "Переведено",

	data.SortProperty:                 "Порядок",
	data.ArtsHistoryEventTimeProperty: "По дате приобретения",
	data.DescendingProperty:           "По убыванию",

	data.BookCompletedProperty: "Прочитано",

	"true":  "Да",
	"false": "Нет",
}

var ShortPropertyTitles = map[string]string{
	data.IllustratorsProperty: "Иллюстр.",
}
