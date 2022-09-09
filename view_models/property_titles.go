package view_models

import "github.com/beauxarts/fedorov/data"

var propertyTitles = map[string]string{
	data.AnyTextProperty:          "Любой текст",
	data.TitleProperty:            "Название",
	data.BookTypeProperty:         "Тип",
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
	data.ISBNPropertyProperty:     "ISBN",
	data.TotalSizeProperty:        "Общий размер",
	data.TotalPagesProperty:       "Всего страниц",
	data.GenresProperty:           "Жанры",
	data.TagsProperty:             "Теги",
	// sorting properties
	data.SortProperty:       "Порядок",
	data.DescendingProperty: "По убыванию",
}
