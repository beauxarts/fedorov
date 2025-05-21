package litres_integration

import (
	"bytes"
	"github.com/boggydigital/match_node"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
)

const (
	TitleProperty           = "Название"
	TypeProperty            = "Тип книги"
	AuthorsProperty         = "Авторы"
	DownloadLinksProperty   = "Загрузки"
	DescriptionProperty     = "Описание"
	SequenceNameProperty    = "Название серии"
	SequenceNumberProperty  = "Номер в серии"
	GenresProperty          = "Жанры"
	TagsProperty            = "Тэги"
	KnownIrrelevantProperty = "known-irrelevant-property"
)

func Reduce(body *html.Node) (map[string][]string, error) {

	rdx := make(map[string][]string)

	bookNameEtc := match_node.NewEtc(atom.Div, "biblio_book_name biblio-book__title-block", true)
	if bbn := match_node.Match(body, bookNameEtc); bbn != nil {
		for n := bbn.FirstChild; n != nil; n = n.NextSibling {
			if n.DataAtom == atom.H1 {
				rdx[TitleProperty] = []string{n.FirstChild.Data}
			}
		}
		labelEtc := match_node.NewEtc(atom.Span, "label", false)
		if label := match_node.Match(bbn, labelEtc); label != nil {
			rdx[TypeProperty] = []string{strings.ToLower(label.FirstChild.Data)}
		}
	}

	authorsEtc := match_node.NewEtc(atom.Div, "biblio_book_author", true)
	if an := match_node.Match(body, authorsEtc); an != nil {
		authorEtc := match_node.NewEtc(atom.A, "biblio_book_author__link", true)
		authorLinks := match_node.Matches(an, authorEtc, -1)
		authors := make([]string, 0, len(authorLinks))
		for _, al := range authorLinks {
			if span := al.FirstChild; span != nil {
				authors = append(authors, span.FirstChild.Data)
			}
		}
		rdx[AuthorsProperty] = authors
	}

	downloadsEtc := match_node.NewEtc(atom.Div, "book_download", false)
	if bdn := match_node.Match(body, downloadsEtc); bdn != nil {
		downloadLinkEtc := match_node.NewEtc(atom.A, "biblio_book_download_file__link", true)
		downloadLinks := match_node.Matches(bdn, downloadLinkEtc, -1)
		links := make([]string, 0, len(downloadLinks))
		for _, dl := range downloadLinks {
			links = append(links, getAttribute(dl, "href"))
		}
		rdx[DownloadLinksProperty] = links
	}

	if len(rdx[DownloadLinksProperty]) == 0 {
		// check for PDF links
		newButtonEtc := match_node.NewEtc(atom.Div, "bb_newbutton", true)
		for _, nb := range match_node.Matches(body, newButtonEtc, -1) {
			if getAttribute(nb, "data-type") == "pdf" {
				rdx[DownloadLinksProperty] = []string{getAttribute(nb.FirstChild, "href")}
			}
		}
	}

	// check for additional materials
	additionalMaterialsEtc := match_node.NewEtc(atom.Div, "bb_newbutton bb_newbutton_add-materials", true)
	if amn := match_node.Match(body, additionalMaterialsEtc); amn != nil {
		linkEtc := match_node.NewEtc(atom.A, "bb_newbutton_inner_link", true)
		if link := match_node.Match(amn, linkEtc); link != nil {
			rdx[DownloadLinksProperty] = append(rdx[DownloadLinksProperty], getAttribute(link, "href"))
		}
	}

	bookDescrEtc := match_node.NewEtc(atom.Div, "biblio_book_descr_publishers", true)
	if bd := match_node.Match(body, bookDescrEtc); bd != nil {

		buf := new(bytes.Buffer)
		if err := html.Render(buf, bd); err != nil {
			return rdx, err
		}
		rdx[DescriptionProperty] = append(rdx[DescriptionProperty], buf.String())
	}

	sequenceNames := make([]string, 0)
	sequenceNumbers := make([]string, 0)

	sequencesEtc := match_node.NewEtc(atom.Div, "biblio_book_sequences", true)
	for _, bbsn := range match_node.Matches(body, sequencesEtc, -1) {

		nameEtc := match_node.NewEtc(atom.A, "biblio_book_sequences__link", true)
		if nan := match_node.Match(bbsn, nameEtc); nan != nil {
			sequenceNames = append(sequenceNames, nan.FirstChild.Data)
		}
		numberEtc := match_node.NewEtc(atom.Span, "number", true)
		if nun := match_node.Match(bbsn, numberEtc); nun != nil {
			sequenceNumbers = append(sequenceNumbers, strings.TrimSpace(nun.FirstChild.Data))
		} else {
			sequenceNumbers = append(sequenceNumbers, "")
		}
	}

	rdx[SequenceNameProperty] = sequenceNames
	rdx[SequenceNumberProperty] = sequenceNumbers

	detailedInfoLeftEtc := match_node.NewEtc(atom.Div, "biblio_book_info_detailed_left", true)
	if din := match_node.Match(body, detailedInfoLeftEtc); din != nil {
		for key, values := range getBookInfoItems(din) {
			rdx[key] = append(rdx[key], values...)
		}
	}

	detailedInfoRightEtc := match_node.NewEtc(atom.Div, "biblio_book_info_detailed_right", true)
	if din := match_node.Match(body, detailedInfoRightEtc); din != nil {
		for key, values := range getBookInfoItems(din) {
			rdx[key] = append(rdx[key], values...)
		}
	}

	bookInfoProperties := []string{GenresProperty, TagsProperty}
	bookInfoEtc := match_node.NewEtc(atom.Div, "biblio_book_info", true)
	if bin := match_node.Match(body, bookInfoEtc); bin != nil {
		liEtc := match_node.NewEtc(atom.Li, "", true)
		bilEtc := match_node.NewEtc(atom.A, "biblio_info__link", true)
		pi := 0
		for _, li := range match_node.Matches(bin, liEtc, -1) {
			ans := match_node.Matches(li, bilEtc, -1)
			property := bookInfoProperties[pi]
			for _, n := range ans {
				rdx[property] = append(rdx[property], bookInfoLinkTextContent(n))
			}
			if len(ans) > 0 {
				pi++
			}
			if pi >= len(bookInfoProperties) {
				break
			}
		}
	}

	return rdx, nil
}

func bookInfoLinkTextContent(biln *html.Node) string {
	tc := ""
	i := 0
	for n := biln.FirstChild; n != nil; n = n.NextSibling {
		switch i {
		case 0:
			tc = n.FirstChild.Data
		case 1:
			tc += n.Data
		default:
			break
		}
		i++
	}
	return tc
}

func getBookInfoItems(node *html.Node) map[string][]string {
	infoItems := make(map[string][]string)
	bii := match_node.NewEtc(atom.Dl, "biblio_book_info_item", false)
	for _, biin := range match_node.Matches(node, bii, -1) {
		p := ""
		for n := biin.FirstChild; n != nil; n = n.NextSibling {
			switch n.DataAtom {
			case atom.Dt:
				switch n.FirstChild.Type {
				case html.TextNode:
					p = n.FirstChild.Data
				case html.ElementNode:
					switch n.FirstChild.DataAtom {
					case atom.Strong:
						p = n.FirstChild.FirstChild.Data
					case atom.A:
						// do nothing
					default:
						panic("unknown dt property container")
					}
				default:
					panic("unknown dt node type")
				}

			case atom.Dd:
				if p == KnownIrrelevantProperty {
					continue
				}
				if p == "" {
					panic("attempt to set unknown property")
				}
				switch n.FirstChild.Type {
				case html.TextNode:
					infoItems[p] = []string{n.FirstChild.Data}
				case html.ElementNode:
					bidl := match_node.NewEtc(atom.Span, "biblio_info_detailed__link", false)
					for _, s := range match_node.Matches(n, bidl, -1) {
						infoItems[p] = append(infoItems[p], s.FirstChild.Data)
					}
				}

			default:
				panic("unknown node type")
			}
		}
	}
	return infoItems
}
