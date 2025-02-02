package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathways"
	"github.com/boggydigital/redux"
	"github.com/boggydigital/yet_urls/youtube_urls"
	"net/http"
	"net/url"
)

const limitVideoRequests = 1000

func GetVideosMetadataHandler(u *url.URL) error {
	force := u.Query().Has("force")

	return GetVideosMetadata(force)
}

func GetVideosMetadata(force bool) error {

	gvma := nod.NewProgress("getting video metadata...")
	defer gvma.End()

	rp, err := pathways.GetAbsRelDir(data.Redux)
	if err != nil {
		return gvma.EndWithError(err)
	}

	rdx, err := redux.NewWriter(rp, data.ReduxProperties()...)
	if err != nil {
		return gvma.EndWithError(err)
	}

	videoIds := make([]string, 0)

	for _, id := range rdx.Keys(data.YouTubeVideosProperty) {
		if vids, ok := rdx.GetAllValues(data.YouTubeVideosProperty, id); ok {
			for _, vid := range vids {

				videoId := ""

				if vidUrl, err := url.Parse(vid); err == nil {
					if videoId = vidUrl.Query().Get("v"); videoId != "" {
					}
				}

				if videoId == "" {
					continue
				}
				if rdx.HasKey(data.VideoTitleProperty, videoId) && !force {
					continue
				}
				if rdx.HasKey(data.VideoErrorProperty, videoId) && !force {
					continue
				}

				videoIds = append(videoIds, videoId)
			}
		}
	}

	if len(videoIds) > limitVideoRequests {
		gvma.EndWithResult("limiting number of videos to avoid IP blacklisting")
		gvma = nod.NewProgress("getting %d videos metadata...", limitVideoRequests)
		videoIds = videoIds[:limitVideoRequests]
	}

	gvma.TotalInt(len(videoIds))
	videoTitles := make(map[string][]string)
	videoDurations := make(map[string][]string)
	videoErrors := make(map[string][]string)

	for _, videoId := range videoIds {

		ipr, err := youtube_urls.GetVideoPage(http.DefaultClient, videoId)
		if err != nil {
			videoErrors[videoId] = append(videoErrors[videoId], err.Error())
			gvma.Error(err)
			gvma.Increment()
			continue
		}

		videoTitles[videoId] = []string{ipr.VideoDetails.Title}
		videoDurations[videoId] = []string{ipr.VideoDetails.LengthSeconds}

		gvma.Increment()
	}

	if err := rdx.BatchAddValues(data.VideoTitleProperty, videoTitles); err != nil {
		return gvma.EndWithError(err)
	}

	if err := rdx.BatchAddValues(data.VideoDurationProperty, videoDurations); err != nil {
		return gvma.EndWithError(err)
	}

	if err := rdx.BatchAddValues(data.VideoErrorProperty, videoErrors); err != nil {
		return gvma.EndWithError(err)
	}

	gvma.EndWithResult("done")

	return nil
}
