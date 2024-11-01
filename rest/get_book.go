package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_pages"
	"github.com/beauxarts/scrinium/litres_integration"
	"github.com/boggydigital/nod"
	"net/http"
)

var artTypeSection = map[litres_integration.ArtsType]string{
	litres_integration.ArtsTypeSimilar: compton_data.SimilarSection,
	litres_integration.ArtsTypeReviews: compton_data.ReviewsSection,
	litres_integration.ArtsTypeFiles:   compton_data.FilesSection,
}

type NewBookViewModel struct {
	Id        string
	Title     string
	Authors   []string
	Downloads []*DownloadViewModel
}

type DownloadViewModel struct {
	Id          string
	Filename    string
	Description string
}

func GetBook(w http.ResponseWriter, r *http.Request) {

	// GET /book?id

	id := r.URL.Query().Get("id")

	var err error
	if rdx, err = rdx.RefreshReader(); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	hasSections := make([]string, 0, 3)
	// every book is expected to have at least those sections
	hasSections = append(hasSections, compton_data.InformationSection, compton_data.AnnotationSection, compton_data.ExternalLinksSection)

	artsTypes := []litres_integration.ArtsType{
		litres_integration.ArtsTypeSimilar,
		litres_integration.ArtsTypeReviews,
	}

	for _, at := range artsTypes {
		if hasArtsType(id, at) {
			hasSections = append(hasSections, artTypeSection[at])
		}
	}

	if videos, ok := rdx.GetAllValues(data.YouTubeVideosProperty, id); ok && len(videos) > 0 {
		hasSections = append(hasSections, compton_data.VideosSection)
	}

	if contentsUrl, ok := rdx.GetLastVal(data.ContentsUrlProperty, id); ok && contentsUrl != "" {
		hasSections = append(hasSections, compton_data.ContentsSection)
	}

	if hasArtsType(id, litres_integration.ArtsTypeFiles) {
		hasSections = append(hasSections, artTypeSection[litres_integration.ArtsTypeFiles])
	}

	//contentsReader, err := data.NewArtsReader(compton_data.ArtTypes)

	if p := compton_pages.Book(id, hasSections, rdx); p != nil {
		if err := p.WriteResponse(w); err != nil {
			http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
			return
		}
	}

	//title := ""
	//if t, ok := rdx.GetLastVal(data.TitleProperty, id); ok {
	//	title = t
	//}
	//
	//var authors []string
	//if aus, err := authorsFullNames(id, rdx); err == nil {
	//	authors = aus
	//}
	//
	//kv, err := data.NewArtsReader(litres_integration.ArtsTypeFiles)
	//if err != nil {
	//	http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//artFiles, err := kv.ArtsFiles(id)
	//if err != nil {
	//	http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//nbvm := &NewBookViewModel{
	//	Id:      id,
	//	Title:   title,
	//	Authors: authors,
	//}
	//
	//for _, dt := range artFiles.DownloadsTypes() {
	//
	//	fn := dt.Filename
	//	if ext := dt.Extension; ext != nil {
	//		fn = strings.Replace(fn, "zip", *ext, 1)
	//	}
	//
	//	dvm := &DownloadViewModel{
	//		Id:          id,
	//		Filename:    fn,
	//		Description: dt.TypeDescription(),
	//	}
	//
	//	nbvm.Files = append(nbvm.Files, dvm)
	//}
	//
	//if err := tmpl.ExecuteTemplate(w, "new_book", nbvm); err != nil {
	//	http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
	//	return
	//}

}

func hasArtsType(id string, at litres_integration.ArtsType) bool {
	if reader, err := data.NewArtsReader(at); err == nil {
		if ok, err := reader.Has(id); ok && err == nil {
			return true
		}
	}
	return false
}
