package view_models

import "path/filepath"

type Downloads struct {
	Id     string
	Files  []string
	Titles []string
}

func NewDownloads(id string, links, titles []string) *Downloads {
	dvm := &Downloads{
		Id:     id,
		Files:  make([]string, 0, len(links)),
		Titles: titles,
	}

	for _, link := range links {
		_, filename := filepath.Split(link)
		dvm.Files = append(dvm.Files, filename)
	}
	return dvm
}
