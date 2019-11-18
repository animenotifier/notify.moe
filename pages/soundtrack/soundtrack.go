package soundtrack

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/server/middleware"
)

// Get track.
func Get(ctx aero.Context) error {
	id := ctx.Get("id")
	track, err := arn.GetSoundTrack(id)
	user := arn.GetUserFromContext(ctx)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Track not found", err)
	}

	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = getOpenGraph(track)
	return ctx.HTML(components.SoundTrackPage(track, user))
}
