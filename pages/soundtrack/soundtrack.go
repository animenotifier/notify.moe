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

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Track not found", err)
	}

	user := arn.GetUserFromContext(ctx)
	animes := track.Anime()

	relatedTracks := arn.FilterSoundTracks(func(t *arn.SoundTrack) bool {
		return !t.IsDraft && len(t.Media) > 0 && t.ID != track.ID && isRelated(animes, t)
	})

	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = getOpenGraph(track)
	return ctx.HTML(components.SoundTrackPage(track, relatedTracks, user))
}

func isRelated(animes []*arn.Anime, track *arn.SoundTrack) bool {
	for _, anime := range animes {
		if arn.Contains(track.Tags, "anime:"+anime.ID) {
			return true
		}
	}

	return false
}
