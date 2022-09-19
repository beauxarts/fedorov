package view_models

import (
	"html/template"
	"strings"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"linkFormat": LinkFormat,
		"formatDesc": formatDesc,
	}
}

func LinkFormat(link string) string {
	format := FormatUnknown

	if strings.HasSuffix(link, ".fb2.zip") {
		format = FormatFB2
	} else if strings.HasSuffix(link, ".ios.epub") {
		format = FormatIOSEPUB
	} else if strings.HasSuffix(link, ".epub") {
		format = FormatEPUB
	} else if strings.HasSuffix(link, ".txt.zip") {
		format = FormatTXTZIP
	} else if strings.HasSuffix(link, ".rtf.zip") {
		format = FormatRTF
	} else if strings.HasSuffix(link, ".a4.pdf") {
		format = FormatPDFA4
	} else if strings.HasSuffix(link, ".a6.pdf") {
		format = FormatPDFA6
	} else if strings.HasSuffix(link, ".pdf") {
		format = FormatPDF
	} else if strings.HasSuffix(link, ".mobi.prc") {
		format = FormatMOBI
	} else if strings.HasSuffix(link, ".fb3") {
		format = FormatFB3
	} else if strings.HasSuffix(link, ".txt") {
		format = FormatTXT
	} else if strings.HasSuffix(link, ".mp3.zip") {
		format = FormatMP3
	} else if strings.HasSuffix(link, ".mp3") {
		format = FormatMP3
	} else if strings.HasSuffix(link, ".m4b") {
		format = FormatMP4
	} else if strings.HasSuffix(link, ".zip") {
		format = FormatZIP
	} else if strings.HasSuffix(link, ".azw3") {
		format = FormatAZW3
	}

	return format
}

func formatDesc(format string) string {
	return formatDescriptors[format]
}
