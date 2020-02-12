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

	relatedTracks := arn.FilterSoundTracks(func(t *arn.SoundTrack) bool {
		return !t.IsDraft && len(t.Media) > 0 && t.ID != track.ID && isRelated(track.Anime(), t)
	})

	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = getOpenGraph(track)
	return ctx.HTML(components.SoundTrackPage(track, relatedTracks, user))
}

func isRelated(anime []*arn.Anime, track *arn.SoundTrack) bool {
	for _, anime := range anime {
		if arn.Contains(track.Tags, "anime:"+anime.ID) {
			return true
		}
	}

	return false
}
