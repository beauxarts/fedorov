package cli

import (
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

	rmbda := nod.NewProgress("reducing my books details...")
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

		rdx := reduceDetails(body)

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

func reduceDetails(body *html.Node) map[string][]string {

	rdx := make(map[string][]string)

	bookNameEtc := match_node.NewEtc(atom.Div, "biblio_book_name")
	if bbn := match_node.Match(body, bookNameEtc); bbn != nil {
		for n := bbn.FirstChild; n != nil; n = n.NextSibling {
			if n.DataAtom == atom.H1 {
				rdx[data.TitleProperty] = []string{n.FirstChild.Data}
			}
		}
		labelEtc := match_node.NewEtc(atom.Span, "label")
		if label := match_node.Match(bbn, labelEtc); label != nil {
			rdx[data.BookTypeProperty] = []string{strings.ToLower(label.FirstChild.Data)}
		}
	}

	authorsEtc := match_node.NewEtc(atom.Div, "biblio_book_author")
	if an := match_node.Match(body, authorsEtc); an != nil {
		authorEtc := match_node.NewEtc(atom.A, "biblio_book_author__link")
		authorLinks := match_node.Matches(an, authorEtc, -1)
		authors := make([]string, 0, len(authorLinks))
		for _, al := range authorLinks {
			if span := al.FirstChild; span != nil {
				authors = append(authors, span.FirstChild.Data)
			}
		}
		rdx[data.AuthorsProperty] = authors
	}

	downloadsEtc := match_node.NewEtc(atom.Div, "book_download")
	if bdn := match_node.Match(body, downloadsEtc); bdn != nil {
		downloadLinkEtc := match_node.NewEtc(atom.A, "biblio_book_download_file__link")
		downloadLinks := match_node.Matches(bdn, downloadLinkEtc, -1)
		links := make([]string, 0, len(downloadLinks))
		for _, dl := range downloadLinks {
			links = append(links, getAttribute(dl, "href"))
		}
		rdx[data.DownloadLinksProperty] = links
	}

	if len(rdx[data.DownloadLinksProperty]) == 0 {
		// check for PDF links
		newButtonEtc := match_node.NewEtc(atom.Div, "bb_newbutton")
		for _, nb := range match_node.Matches(body, newButtonEtc, -1) {
			if getAttribute(nb, "data-type") == "pdf" {
				rdx[data.DownloadLinksProperty] = []string{getAttribute(nb.FirstChild, "href")}
			}
		}
	}

	// check for additional materials
	additionalMaterialsEtc := match_node.NewEtc(atom.Div, "bb_newbutton bb_newbutton_add-materials")
	if amn := match_node.Match(body, additionalMaterialsEtc); amn != nil {
		linkEtc := match_node.NewEtc(atom.A, "bb_newbutton_inner_link")
		if link := match_node.Match(amn, linkEtc); link != nil {
			rdx[data.DownloadLinksProperty] = append(rdx[data.DownloadLinksProperty], getAttribute(link, "href"))
		}
	}

	sequenceNames := make([]string, 0)
	sequenceNumbers := make([]string, 0)

	sequencesEtc := match_node.NewEtc(atom.Div, "biblio_book_sequences")
	for _, bbsn := range match_node.Matches(body, sequencesEtc, -1) {

		nameEtc := match_node.NewEtc(atom.A, "biblio_book_sequences__link")
		if nan := match_node.Match(bbsn, nameEtc); nan != nil {
			sequenceNames = append(sequenceNames, nan.FirstChild.Data)
		}
		numberEtc := match_node.NewEtc(atom.Span, "number")
		if nun := match_node.Match(bbsn, numberEtc); nun != nil {
			sequenceNumbers = append(sequenceNumbers, strings.TrimSpace(nun.FirstChild.Data))
		} else {
			sequenceNumbers = append(sequenceNumbers, "")
		}
	}

	rdx[data.SequenceNameProperty] = sequenceNames
	rdx[data.SequenceNumberProperty] = sequenceNumbers

	detailedInfoLeftEtc := match_node.NewEtc(atom.Ul, "biblio_book_info_detailed_left")
	if din := match_node.Match(body, detailedInfoLeftEtc); din != nil {
		for n := din.FirstChild; n != nil; n = n.NextSibling {
			if n.FirstChild == nil {
				continue
			}
			strong := n.FirstChild
			property := propertyByStrongTitle(strong.FirstChild.Data)
			if property == "" {
				continue
			}

			values := make([]string, 0)

			linksEtc := match_node.NewEtc(atom.Span, "biblio_info_detailed__link")
			for _, link := range match_node.Matches(n, linksEtc, -1) {
				values = append(values, link.FirstChild.Data)
			}

			if len(values) == 0 {
				values = []string{strings.TrimSpace(strong.NextSibling.Data)}
			}

			rdx[property] = append(rdx[property], values...)
		}
	}

	detailedInfoRightEtc := match_node.NewEtc(atom.Ul, "biblio_book_info_detailed_right")
	if din := match_node.Match(body, detailedInfoRightEtc); din != nil {
		for n := din.FirstChild; n != nil; n = n.NextSibling {
			if n.FirstChild == nil {
				continue
			}
			strong := n.FirstChild
			property := propertyByStrongTitle(strong.FirstChild.Data)
			if property == "" {
				continue
			}

			values := make([]string, 0)

			linksEtc := match_node.NewEtc(atom.Span, "biblio_info_detailed__link")
			for _, link := range match_node.Matches(n, linksEtc, -1) {
				values = append(values, link.FirstChild.Data)
			}

			if len(values) == 0 {
				val := strong.NextSibling.Data
				if property == data.ISBNPropertyProperty {
					val = strong.NextSibling.NextSibling.FirstChild.Data
				}

				values = []string{strings.TrimSpace(val)}
			}

			rdx[property] = append(rdx[property], values...)
		}
	}

	return rdx
}

func getAttribute(node *html.Node, attrName string) string {
	for _, attr := range node.Attr {
		if attr.Key == attrName {
			return attr.Val
		}
	}
	return ""
}

func propertyByStrongTitle(key string) string {
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
		// do nothing
	case "Размер страницы:":
		// do nothing

	default:
		nod.Log("unknown detailed info key: %s", key)
		return ""
	}
	return property
}
