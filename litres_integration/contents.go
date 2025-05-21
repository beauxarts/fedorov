package litres_integration

import "encoding/xml"

type Contents struct {
	XMLName xml.Name `xml:"toc"`
	Text    string   `xml:",chardata"`
	TocItem []struct {
		Text string `xml:",chardata"`
		N    string `xml:"n,attr"`
		Deep string `xml:"deep,attr"`
		ID   string `xml:"id,attr"`
	} `xml:"toc-item"`
}
