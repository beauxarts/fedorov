package view_models

import "html/template"

type Description struct {
	Id      string
	Content template.HTML
}

func NewDescription(id, desc string) *Description {
	return &Description{
		Id:      id,
		Content: template.HTML(desc),
	}
}
