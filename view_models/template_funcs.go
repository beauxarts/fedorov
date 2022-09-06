package view_models

import "html/template"

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"propertyTitle": propertyTitle,
	}
}

func propertyTitle(p string) string {
	return propertyTitles[p]
}
