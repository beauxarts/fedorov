package compton_pages

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/rest/compton_data"
	"github.com/beauxarts/fedorov/rest/compton_fragments"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/redux"
	"net/url"
)

func Videos(id string, rdx redux.Readable) compton.PageElement {

	s := compton_fragments.ProductSection(compton_data.VideosSection, id, rdx)

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

	pageStack := compton.FlexItems(s, direction.Column)

	s.Append(pageStack)

	if len(videoIds) == 0 {
		fs := compton.Fspan(s, "Для данной книги нет видео").
			ForegroundColor(color.RepGray)
		pageStack.Append(compton.FICenter(s, fs))
	}

	for ii, videoId := range videoIds {
		videoTitle := videoTitles[videoId]
		videoDuration := videoDurations[videoId]
		pageStack.Append(compton_fragments.VideoOriginLink(s, videoId, videoTitle, videoDuration))

		if ii != len(videoIds)-1 {
			pageStack.Append(compton.Hr())
		}
	}

	return s

}
