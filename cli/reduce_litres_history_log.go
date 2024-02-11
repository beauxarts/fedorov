package cli

import (
	"fmt"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/match_node"
	"github.com/boggydigital/nod"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	historyEventsPerPage = 40
)

func ReduceLitResHistoryLogHandler(_ *url.URL) error {
	return ReduceLitResHistoryLog()
}

func ReduceLitResHistoryLog() error {

	rhla := nod.NewProgress("reducing history log...")
	defer rhla.End()

	absLitResHistoryLogDir, err := data.AbsDataTypeDir(litres_integration.LitResHistoryLog)
	if err != nil {
		return rhla.EndWithError(err)
	}

	kv, err := kvas.ConnectLocal(absLitResHistoryLogDir, kvas.HtmlExt)
	if err != nil {
		return rhla.EndWithError(err)
	}

	totalPages := len(kv.Keys())
	artsHistoryOrder := make([]string, 0, totalPages*historyEventsPerPage)
	artsHistoryEventTime := make(map[string][]string, totalPages*historyEventsPerPage)
	//artsIds := make(map[string][]string, totalPages*historyEventsPerPage)

	booksBoxEtc := match_node.NewEtc(atom.Div, "books_box", true)
	historyEventEtc := match_node.NewEtc(atom.Div, "history-event", true)

	for p := 1; p <= totalPages; p++ {
		page, err := kv.Get(strconv.FormatInt(int64(p), 10))
		if err != nil {
			return rhla.EndWithError(err)
		}

		body, err := html.Parse(page)
		if err != nil {
			return rhla.EndWithError(err)
		}

		for _, bb := range match_node.Matches(body, booksBoxEtc, -1) {
			for _, he := range match_node.Matches(bb, historyEventEtc, -1) {

				if contentArts := eventContentArts(he); len(contentArts) > 0 {

					artsHistoryOrder = append(artsHistoryOrder, contentArts...)
					et := eventTime(he)

					for _, ca := range contentArts {
						artsHistoryEventTime[ca] = []string{et.Format(time.RFC3339)}
					}

				}
			}
		}
	}

	rdx, err := data.NewReduxWriter(
		data.ArtsHistoryOrderProperty,
		data.ArtsHistoryEventTimeProperty,
		data.ImportedProperty)
	if err != nil {
		return rhla.EndWithError(err)
	}

	sra := nod.Begin(" saving redux...")
	defer sra.End()

	// add previously imported book to my books before saving that set
	artsHistoryOrder = append(artsHistoryOrder, rdx.Keys(data.ImportedProperty)...)

	if err := rdx.ReplaceValues(data.ArtsHistoryOrderProperty, data.ArtsHistoryOrderProperty, artsHistoryOrder...); err != nil {
		return rhla.EndWithError(err)
	}

	if err := rdx.BatchReplaceValues(data.ArtsHistoryEventTimeProperty, artsHistoryEventTime); err != nil {
		return rhla.EndWithError(err)
	}

	sra.EndWithResult("done")
	rhla.EndWithResult("done")

	return nil
}

func eventTime(historyEvent *html.Node) time.Time {
	dateStr, timeStr := "", ""
	eventDateTimeEtc := match_node.NewEtc(atom.Div, "event-datetime", false)
	eventDateEtc := match_node.NewEtc(atom.P, "event-date", true)
	eventTimeEtc := match_node.NewEtc(atom.P, "event-time", true)
	if edt := match_node.Match(historyEvent, eventDateTimeEtc); edt != nil {
		if ed := match_node.Match(edt, eventDateEtc); ed != nil && ed.FirstChild != nil {
			dateStr = ed.FirstChild.Data
		}
		if et := match_node.Match(edt, eventTimeEtc); et != nil && et.FirstChild != nil {
			timeStr = et.FirstChild.Data
		}

		if t, err := time.Parse("02.01.06 15:04", fmt.Sprintf("%s %s", dateStr, timeStr)); err == nil {
			return t
		}
	}
	return time.Now()
}

func eventContentArts(historyEvent *html.Node) []string {
	contentArts := make([]string, 0)
	contentArtEtc := match_node.NewEtc(atom.Div, "content-art", true)
	for _, ca := range match_node.Matches(historyEvent, contentArtEtc, -1) {
		if ca.FirstChild != nil && ca.FirstChild.DataAtom == atom.A {
			for _, attr := range ca.FirstChild.Attr {
				if attr.Key == "href" {
					contentArts = append(contentArts, artFromHref(attr.Val))
				}
			}
		}
	}
	return contentArts
}

func artFromHref(href string) string {
	if strings.Contains(href, "biblio_book") {
		if _, id, ok := strings.Cut(href, "="); ok {
			return id
		}
	} else {
		if parts := strings.Split(strings.TrimSuffix(href, "/"), "-"); len(parts) > 0 {
			return parts[len(parts)-1]
		}
	}
	return href
}

func idHref(node *html.Node) (string, string) {
	id, href := "", ""
	for _, attr := range node.Attr {
		if attr.Key == "data-art" {
			id = attr.Val
		}
		if attr.Key == "href" {
			href = attr.Val
		}
		if id != "" && href != "" {
			break
		}
	}
	return id, href
}
