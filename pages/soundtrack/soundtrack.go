package soundtrack

import (
	"net/http"

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

	ctx.Data = getOpenGraph(ctx, track)

	// Comments
	posts := track.Posts()
	arn.SortPostsLatestFirst(posts)

	return ctx.HTML(components.SoundTrackPage(track, posts, user))
}
