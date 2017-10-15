package soundtracks

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxTracks = 9

// Get renders the soundtracks page.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	tracks, err := arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && len(track.Media) > 0
	})

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching soundtracks", err)
	}

	arn.SortSoundTracksLatestFirst(tracks)

	if len(tracks) > maxTracks {
		tracks = tracks[:maxTracks]
	}

	return ctx.HTML(components.SoundTracks(tracks, user))
}
