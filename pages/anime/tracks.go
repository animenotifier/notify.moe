package anime

import (
	"net/http"

	"github.com/animenotifier/notify.moe/utils"

	"github.com/animenotifier/notify.moe/components"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Tracks ...
func Tracks(ctx *aero.Context) string {
	id := ctx.Get("id")
	user := utils.GetUser(ctx)

	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	tracks := arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && len(track.Media) > 0 && arn.Contains(track.Tags, "anime:"+anime.ID)
	})

	return ctx.HTML(components.AnimeTracks(anime, tracks, user, true))
}
