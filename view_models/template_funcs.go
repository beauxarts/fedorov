package view_models

import (
	"html/template"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		//"linkFormat": data.LinkFormat,
		//"formatDesc": data.FormatDesc,
	}
}
