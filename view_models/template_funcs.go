package view_models

import (
	"github.com/beauxarts/fedorov/data"
	"html/template"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"linkFormat": data.LinkFormat,
		"formatDesc": data.FormatDesc,
	}
}
