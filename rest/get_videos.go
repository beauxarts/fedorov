package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_pages"
	"github.com/boggydigital/nod"
	"net/http"
	"net/url"
)

func GetVideos(w http.ResponseWriter, r *http.Request) {

	// GET /videos?id

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, nod.ErrorStr("missing book id"), http.StatusInternalServerError)
		return
	}

	var err error
	if rdx, err = rdx.RefreshReader(); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	var videoIds []string
	if vids, ok := rdx.GetAllValues(data.YouTubeVideosProperty, id); ok {
		for _, vid := range vids {
			if vidUrl, err := url.Parse(vid); err == nil {
				if v := vidUrl.Query().Get("v"); v != "" {
					if !rdx.HasKey(data.VideoErrorProperty, v) {
						videoIds = append(videoIds, v)
					}
				}
			}
		}
	}

	videoTitles := make(map[string]string)
	videoDurations := make(map[string]string)

	for _, videoId := range videoIds {
		if vt, ok := rdx.GetLastVal(data.VideoTitleProperty, videoId); ok {
			videoTitles[videoId] = vt
		}
		if vd, ok := rdx.GetLastVal(data.VideoDurationProperty, videoId); ok {
			videoDurations[videoId] = vd
		}
	}

	if p := compton_pages.Videos(videoIds, videoTitles, videoDurations); p != nil {
		if err := p.WriteResponse(w); err != nil {
			http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		}
	}

}
