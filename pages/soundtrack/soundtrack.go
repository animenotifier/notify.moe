package soundtrack

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get track.
func Get(ctx *aero.Context) string {
	id := ctx.Get("id")
	track, err := arn.GetSoundTrack(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Track not found", err)
	}

	openGraph := &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":     track.Title,
			"og:url":       "https://" + ctx.App.Config.Domain + track.Link(),
			"og:site_name": "notify.moe",
			"og:type":      "music.song",
		},
	}

	if track.MainAnime() != nil {
		openGraph.Tags["og:image"] = track.MainAnime().Image.Large
	}

	// Set video so that it can be played
	youtube := track.MediaByName("Youtube")
	if len(youtube) > 0 {
		openGraph.Tags["og:video"] = "https://www.youtube.com/v/" + youtube[0].ServiceID
	}

	ctx.Data = openGraph

	return ctx.HTML(components.Track(track))
}
