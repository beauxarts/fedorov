package compton_data

import "github.com/beauxarts/fedorov/data"

var PropertyTitles = map[string]string{
	data.TitleProperty:          "Название",
	data.HTMLAnnotationProperty: "Аннотация",

	data.AuthorsProperty:      "Автор",
	data.ReadersProperty:      "Читает",
	data.TranslatorsProperty:  "Перевод",
	data.IllustratorsProperty: "Иллюстр.",
	data.PerformersProperty:   "Исполнитель",
	data.PaintersProperty:     "Художник",
	data.PublishersProperty:   "Издатель",
	data.RightholdersProperty: "Правообладатель",

	data.PriceProperty:           "Цена",
	data.RatedAvgProperty:        "Рейтинг (ЛитРес)",
	data.LivelibRatedAvgProperty: "Рейтинг (LiveLib)",

	data.ArtTypeProperty: "Тип",
	data.GenresProperty:  "Жанры",
	data.TagsProperty:    "Тэги",
	data.MinAgeProperty:  "Возрастное ограничение",
	data.ISBNProperty:    "ISBN",
	data.SeriesProperty:  "Серия",

	data.DateWrittenAtProperty:   "Написано",
	data.PublicationDateProperty: "Опубликовано",
	data.TranslatedAtProperty:    "Переведено",
	data.RegisteredAtProperty:    "Зарегистрировано",
	data.AvailableFromProperty:   "Доступно",
	data.FirstTimeSaleAtProperty: "Начало продаж",
	data.LastReleasedAtProperty:  "Выпущено",
	data.LastUpdatedAtProperty:   "Обновлено",

	data.CurrentPagesOrSecondsProperty: "Объем",

	data.SortProperty:                 "Порядок",
	data.ArtsHistoryEventTimeProperty: "По дате приобретения",
	data.DescendingProperty:           "По убыванию",

	data.LitresBookLinksProperty:         "Книги",
	data.LitresAuthorLinksProperty:       "Авторы",
	data.LitresSeriesLinksProperty:       "Серии",
	data.LitresPublishersLinksProperty:   "Издатели",
	data.LitresRightholdersLinksProperty: "Правообладатели",
	data.LitresGenresLinksProperty:       "Жанры",
	data.LitresTagsLinksProperty:         "Tэги",

	"true":  "Да",
	"false": "Нет",
}

var ShortPropertyTitles = map[string]string{
	data.IllustratorsProperty: "Иллюстр.",
}
