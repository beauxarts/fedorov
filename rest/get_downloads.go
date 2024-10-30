package rest

import (
	"github.com/beauxarts/fedorov/rest/compton_pages"
	"github.com/boggydigital/nod"
	"net/http"
)

func GetFiles(w http.ResponseWriter, r *http.Request) {

	// GET /files?id

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, nod.ErrorStr("missing required book id"), http.StatusInternalServerError)
		return
	}

	var err error
	if rdx, err = rdx.RefreshReader(); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	if p := compton_pages.Files(id); p != nil {
		if err := p.WriteResponse(w); err != nil {
			http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		}
	}

	//links, ok := rdx.GetAllValues("" /*data.DownloadLinksProperty*/, idstr)
	//titles, _ := rdx.GetAllValues("" /*data.DownloadTitlesProperty*/, idstr)
	//
	//if !ok {
	//	http.Error(w, nod.ErrorStr("book has no downloads"), http.StatusInternalServerError)
	//	return
	//}
	//
	//files := make([]string, 0, len(links))
	//
	//if id, err := strconv.ParseInt(idstr, 10, 64); err == nil {
	//	for _, link := range links {
	//		filename := filepath.Base(link)
	//		absDownloadFilename, err := data.AbsFileDownloadPath(id, filename)
	//		if err != nil {
	//			http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
	//			return
	//		}
	//		if _, err := os.Stat(absDownloadFilename); err == nil {
	//			files = append(files, filename)
	//		} else {
	//			nod.Log(err.Error())
	//		}
	//	}
	//}
	//
	//sb := &strings.Builder{}
	//dvm := view_models.NewDownloads(idstr, files, titles)
	//
	//if err := tmpl.ExecuteTemplate(sb, "downloads", dvm); err != nil {
	//	http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//if err := app.RenderSection(idstr, stencil_app.DownloadsSection, sb.String(), w); err != nil {
	//	http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
	//	return
	//}

}
