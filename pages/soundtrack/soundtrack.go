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

	relatedTracks := make([]*arn.SoundTrack, 0, 5)
	for _, anime := range track.Anime() {
		anime := anime
		tracks := arn.FilterSoundTracks(func(t *arn.SoundTrack) bool {
			return !t.IsDraft && len(t.Media) > 0 && t.ID != track.ID && arn.Contains(t.Tags, "anime:"+anime.ID)
		})
		relatedTracks = append(relatedTracks, tracks...)
	}

	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = getOpenGraph(track)
	return ctx.HTML(components.SoundTrackPage(track, relatedTracks, user))
}
