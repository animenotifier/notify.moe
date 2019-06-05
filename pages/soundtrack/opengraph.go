package soundtrack

import (
	"strings"

	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/assets"
)

func getOpenGraph(track *arn.SoundTrack) *arn.OpenGraph {
	openGraph := &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":     track.Title.ByUser(nil),
			"og:url":       "https://" + assets.Domain + track.Link(),
			"og:site_name": assets.Domain,
			"og:type":      "music.song",
		},
	}

	descriptionTags := []string{}

	for _, tag := range track.Tags {
		if strings.HasPrefix(tag, "anime:") {
			continue
		}

		descriptionTags = append(descriptionTags, tag)
	}

	if track.MainAnime() != nil {
		openGraph.Tags["og:image"] = track.MainAnime().ImageLink("large")
		openGraph.Tags["og:description"] = track.MainAnime().Title.Canonical + " (" + strings.Join(descriptionTags, ", ") + ")"
	}

	if track.File != "" {
		openGraph.Tags["og:audio"] = "https://" + assets.Domain + "/audio/" + track.File
		openGraph.Tags["og:audio:type"] = "audio/vnd.facebook.bridge"
	}

	// Set video so that it can be played
	youtube := track.MediaByService("Youtube")

	if len(youtube) > 0 {
		openGraph.Tags["og:video"] = "https://www.youtube.com/v/" + youtube[0].ServiceID
	}

	return openGraph
}
