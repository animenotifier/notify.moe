package soundtrack

import (
	"net/http"
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get track.
func Get(ctx *aero.Context) string {
	id := ctx.Get("id")
	track, err := arn.GetSoundTrack(id)
	user := utils.GetUser(ctx)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Track not found", err)
	}

	descriptionTags := []string{}

	for _, tag := range track.Tags {
		if strings.HasPrefix(tag, "anime:") {
			continue
		}

		descriptionTags = append(descriptionTags, tag)
	}

	openGraph := &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       track.Title,
			"og:description": track.MainAnime().Title.Canonical + " (" + strings.Join(descriptionTags, ", ") + ")",
			"og:url":         "https://" + ctx.App.Config.Domain + track.Link(),
			"og:site_name":   ctx.App.Config.Domain,
			"og:type":        "music.song",
		},
	}

	if track.MainAnime() != nil {
		openGraph.Tags["og:image"] = track.MainAnime().ImageLink("large")
	}

	if track.File != "" {
		openGraph.Tags["og:audio"] = "https://" + ctx.App.Config.Domain + "/audio/" + track.File
		openGraph.Tags["og:audio:type"] = "audio/vnd.facebook.bridge"
	}

	// Set video so that it can be played
	youtube := track.MediaByService("Youtube")

	if len(youtube) > 0 {
		openGraph.Tags["og:video"] = "https://www.youtube.com/v/" + youtube[0].ServiceID
	}

	ctx.Data = openGraph

	return ctx.HTML(components.SoundTrackPage(track, user))
}
