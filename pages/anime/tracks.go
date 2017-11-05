package anime

import (
	"net/http"

	"github.com/animenotifier/notify.moe/components"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Tracks ...
func Tracks(ctx *aero.Context) string {
	id := ctx.Get("id")

	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	tracks, err := arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && len(track.Media) > 0 && arn.Contains(track.Tags, "anime:"+anime.ID)
	})

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Error fetching soundtracks", err)
	}

	return ctx.HTML(components.AnimeTracks(anime, tracks))
}
