package soundtrack

import (
	"net/http"

	"github.com/animenotifier/notify.moe/assets"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/middleware"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/animenotifier/notify.moe/utils/editform"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Edit track.
func Edit(ctx aero.Context) error {
	id := ctx.Get("id")
	track, err := arn.GetSoundTrack(id)
	user := utils.GetUser(ctx)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Track not found", err)
	}

	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":     track.Title.ByUser(user),
			"og:url":       "https://" + assets.Domain + track.Link(),
			"og:site_name": "notify.moe",
			"og:type":      "music.song",
		},
	}

	if track.MainAnime() != nil {
		customCtx.OpenGraph.Tags["og:image"] = track.MainAnime().ImageLink("large")
	}

	return ctx.HTML(components.SoundTrackTabs(track, user) + editform.Render(track, "Edit soundtrack", user))
}
