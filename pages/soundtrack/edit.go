package soundtrack

import (
	"net/http"

	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/animenotifier/notify.moe/utils/editform"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Edit track.
func Edit(ctx *aero.Context) string {
	id := ctx.Get("id")
	track, err := arn.GetSoundTrack(id)
	user := utils.GetUser(ctx)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Track not found", err)
	}

	ctx.Data = &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":     track.NewTitle.ByUser(user),
			"og:url":       "https://" + ctx.App.Config.Domain + track.Link(),
			"og:site_name": "notify.moe",
			"og:type":      "music.song",
		},
	}

	if track.MainAnime() != nil {
		ctx.Data.(*arn.OpenGraph).Tags["og:image"] = track.MainAnime().ImageLink("large")
	}

	return ctx.HTML(components.SoundTrackTabs(track, user) + editform.Render(track, "Edit soundtrack", user))
}
