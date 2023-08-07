package stencil_app

import "github.com/beauxarts/fedorov/data"

var PropertyTitles = map[string]string{
	data.TitleProperty:            "Название",
	data.BookTypeProperty:         "Тип",
	data.BookCompletedProperty:    "Прочитано",
	data.AuthorsProperty:          "Автор(ы)",
	data.CoauthorsProperty:        "Cоавтор(ы)",
	data.TranslatorsProperty:      "Переводчик(и)",
	data.ReadersProperty:          "Чтец(ы)",
	data.IllustratorsProperty:     "Иллюстратор(ы)",
	data.ComposersProperty:        "Композитор(ы)",
	data.AdapterProperty:          "Адаптация",
	data.PerformersProperty:       "Исполнители",
	data.DirectorsProperty:        "Режиссер(ы)",
	data.SoundDirectorsProperty:   "Звукорежиссер(ы)",
	data.DescriptionProperty:      "Описание",
	data.CopyrightHoldersProperty: "Правообладатели",
	data.PublishersProperty:       "Издатели",
	data.HrefProperty:             "Ссылки",
	data.SequenceNameProperty:     "Серия",
	data.SequenceNumberProperty:   "Номер в серии",
	data.DateReleasedProperty:     "Опубликовано",
	data.DateTranslatedProperty:   "Переведено",
	data.DateCreatedProperty:      "Написано",
	data.AgeRatingProperty:        "Для возраста",
	data.DurationProperty:         "Длительность",
	data.VolumeProperty:           "Объем",
	data.ISBNProperty:             "ISBN",
	data.TotalSizeProperty:        "Общий размер",
	data.TotalPagesProperty:       "Всего страниц",
	data.GenresProperty:           "Жанры",
	data.TagsProperty:             "Теги",
	data.LocalTagsProperty:        "Свои теги",
	data.ImportedProperty:         "Импортированная",
	data.LanguageProperty:         "Язык",
	// sorting properties
	data.SortProperty:       "Порядок",
	data.DescendingProperty: "По убыванию",
	// sorting values
	data.MyBooksOrderProperty: "Мои книги",
	// Yes/No
	"true":  "Да",
	"false": "Нет",
}
