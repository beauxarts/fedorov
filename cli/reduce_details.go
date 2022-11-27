package cli

import (
	"bytes"
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/match_node"
	"github.com/boggydigital/nod"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/url"
	"strings"
)

func ReduceDetailsHandler(_ *url.URL) error {
	return ReduceDetails()
}

func ReduceDetails() error {

	rmbda := nod.NewProgress("reducing details...")
	defer rmbda.End()

	reduxProps := data.ReduxProperties()

	reductions := make(map[string]map[string][]string, len(reduxProps))
	for _, p := range reduxProps {
		reductions[p] = make(map[string][]string)
	}

	missingDetails := make([]string, 0)

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), nil, reduxProps...)
	if err != nil {
		return rmbda.EndWithError(err)
	}

	kv, err := kvas.ConnectLocal(data.AbsMyBooksDetailsDir(), kvas.HtmlExt)
	if err != nil {
		return rmbda.EndWithError(err)
	}

	ids := kv.Keys()

	rmbda.TotalInt(len(ids))

	for _, id := range ids {

		det, err := kv.Get(id)
		if err != nil {
			det.Close()
			return rmbda.EndWithError(err)
		}

		body, err := html.Parse(det)
		if err != nil {
			det.Close()
			return rmbda.EndWithError(err)
		}

		rdx, err := reduceDetails(body)
		if err != nil {
			det.Close()
			return rmbda.EndWithError(err)
		}

		if isEmpty(rdx) {
			missingDetails = append(missingDetails, id)
		}

		for p, vals := range rdx {
			reductions[p][id] = vals
		}

		det.Close()
		rmbda.Increment()
	}

	sra := nod.NewProgress(" saving reductions...")
	defer sra.End()

	sra.TotalInt(len(reductions))

	if err := rxa.ReplaceValues(data.MissingDetailsIdsProperty, data.MissingDetailsIdsProperty, missingDetails...); err != nil {
		rmbda.EndWithError(err)
	}

	for prop, rdx := range reductions {
		if err := rxa.BatchReplaceValues(prop, rdx); err != nil {
			return rmbda.EndWithError(err)
		}
		sra.Increment()
	}

	sra.EndWithResult("done")
	rmbda.EndWithResult("done")

	return nil
}

func isEmpty(rdx map[string][]string) bool {
	isEmpty := true

	for _, vals := range rdx {
		isEmpty = isEmpty && len(vals) == 0
	}

	return isEmpty
}

func reduceDetails(body *html.Node) (map[string][]string, error) {

	rdx := make(map[string][]string)

	bookNameEtc := match_node.NewEtc(atom.Div, "biblio_book_name biblio-book__title-block", true)
	if bbn := match_node.Match(body, bookNameEtc); bbn != nil {
		for n := bbn.FirstChild; n != nil; n = n.NextSibling {
			if n.DataAtom == atom.H1 {
				rdx[data.TitleProperty] = []string{n.FirstChild.Data}
			}
		}
		labelEtc := match_node.NewEtc(atom.Span, "label", false)
		if label := match_node.Match(bbn, labelEtc); label != nil {
			rdx[data.BookTypeProperty] = []string{strings.ToLower(label.FirstChild.Data)}
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
		rdx[data.AuthorsProperty] = authors
	}

	downloadsEtc := match_node.NewEtc(atom.Div, "book_download", false)
	if bdn := match_node.Match(body, downloadsEtc); bdn != nil {
		downloadLinkEtc := match_node.NewEtc(atom.A, "biblio_book_download_file__link", true)
		downloadLinks := match_node.Matches(bdn, downloadLinkEtc, -1)
		links := make([]string, 0, len(downloadLinks))
		for _, dl := range downloadLinks {
			links = append(links, getAttribute(dl, "href"))
		}
		rdx[data.DownloadLinksProperty] = links
	}

	if len(rdx[data.DownloadLinksProperty]) == 0 {
		// check for PDF links
		newButtonEtc := match_node.NewEtc(atom.Div, "bb_newbutton", true)
		for _, nb := range match_node.Matches(body, newButtonEtc, -1) {
			if getAttribute(nb, "data-type") == "pdf" {
				rdx[data.DownloadLinksProperty] = []string{getAttribute(nb.FirstChild, "href")}
			}
		}
	}

	// check for additional materials
	additionalMaterialsEtc := match_node.NewEtc(atom.Div, "bb_newbutton bb_newbutton_add-materials", true)
	if amn := match_node.Match(body, additionalMaterialsEtc); amn != nil {
		linkEtc := match_node.NewEtc(atom.A, "bb_newbutton_inner_link", true)
		if link := match_node.Match(amn, linkEtc); link != nil {
			rdx[data.DownloadLinksProperty] = append(rdx[data.DownloadLinksProperty], getAttribute(link, "href"))
		}
	}

	bookDescrEtc := match_node.NewEtc(atom.Div, "biblio_book_descr", true)
	if bd := match_node.Match(body, bookDescrEtc); bd != nil {

		buf := new(bytes.Buffer)
		if err := html.Render(buf, bd); err != nil {
			return rdx, err
		}
		rdx[data.DescriptionProperty] = append(rdx[data.DescriptionProperty], buf.String())
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

	rdx[data.SequenceNameProperty] = sequenceNames
	rdx[data.SequenceNumberProperty] = sequenceNumbers

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

	bookInfoProperties := []string{data.GenresProperty, data.TagsProperty}
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

func getAttribute(node *html.Node, attrName string) string {
	for _, attr := range node.Attr {
		if attr.Key == attrName {
			return attr.Val
		}
	}
	return ""
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
					p = propertyByTitle(n.FirstChild.Data)
				case html.ElementNode:
					switch n.FirstChild.DataAtom {
					case atom.Strong:
						p = propertyByTitle(n.FirstChild.FirstChild.Data)
					case atom.A:
						// do nothing
					default:
						panic("unknown dt property container")
					}
				default:
					panic("unknown dt node type")
				}

			case atom.Dd:
				if p == data.KnownIrrelevantProperty {
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

func propertyByTitle(key string) string {
	property := ""
	switch key {
	case "Соавтор:":
		property = data.CoauthorsProperty
	case "Возрастное ограничение:":
		property = data.AgeRatingProperty
	case "Объем:":
		property = data.VolumeProperty
	case "Длительность:":
		property = data.DurationProperty
	case "Дата выхода на ЛитРес:":
		property = data.DateReleasedProperty
	case "Дата перевода:":
		property = data.DateTranslatedProperty
	case "Дата написания:":
		property = data.DateCreatedProperty
	case "ISBN:":
		property = data.ISBNPropertyProperty
	case "Переводчики:":
		fallthrough
	case "Переводчик:":
		property = data.TranslatorsProperty
	case "Чтецы:":
		fallthrough
	case "Чтец:":
		property = data.ReadersProperty
	case "Художники:":
		fallthrough
	case "Художник:":
		property = data.IllustratorsProperty
	case "Правообладатели:":
		fallthrough
	case "Правообладатель:":
		property = data.CopyrightHoldersProperty
	case "Композиторы:":
		fallthrough
	case "Композитор:":
		property = data.ComposersProperty
	case "Адаптация:":
		property = data.AdapterProperty
	case "Исполнители:":
		property = data.PerformersProperty
	case "Режиссер:":
		property = data.DirectorsProperty
	case "Звукорежиссер:":
		property = data.SoundDirectorsProperty
	case "Издатели:":
		fallthrough
	case "Издатель:":
		property = data.PublishersProperty
	case "Общий размер:":
		property = data.TotalSizeProperty
	case "Общее кол-во страниц:":
		property = data.TotalPagesProperty

	case "Оглавление":
		property = data.KnownIrrelevantProperty
	case "Размер страницы:":
		property = data.KnownIrrelevantProperty

	default:
		nod.Log("unknown detailed info key: %s", key)
		return ""
	}
	return property
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
