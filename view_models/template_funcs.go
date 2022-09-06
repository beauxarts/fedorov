package view_models

import (
	"html/template"
	"strings"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"propertyTitle": propertyTitle,
		"linkFormat":    linkFormat,
		"formatDesc":    formatDesc,
	}
}

func propertyTitle(p string) string {
	return propertyTitles[p]
}

func linkFormat(link string) string {
	if strings.HasSuffix(link, ".fb2.zip") {
		return FormatFB2
	} else if strings.HasSuffix(link, ".ios.epub") {
		return FormatIOSEPUB
	} else if strings.HasSuffix(link, ".epub") {
		return FormatEPUB
	} else if strings.HasSuffix(link, ".txt.zip") {
		return FormatTXTZIP
	} else if strings.HasSuffix(link, ".rtf.zip") {
		return FormatRTF
	} else if strings.HasSuffix(link, ".a4.pdf") {
		return FormatPDFA4
	} else if strings.HasSuffix(link, ".a6.pdf") {
		return FormatPDFA6
	} else if strings.HasSuffix(link, ".pdf") {
		return FormatPDF
	} else if strings.HasSuffix(link, ".mobi.prc") {
		return FormatMOBI
	} else if strings.HasSuffix(link, ".fb3") {
		return FormatFB3
	} else if strings.HasSuffix(link, ".txt") {
		return FormatTXT
	} else if strings.HasSuffix(link, ".mp3.zip") {
		return FormatMP3
	} else if strings.HasSuffix(link, ".m4b") {
		return FormatMP4
	} else if strings.HasSuffix(link, ".zip") {
		return FormatZIP
	}

	return FormatUnknown
}

func formatDesc(format string) string {
	return formatDescriptors[format]
}
