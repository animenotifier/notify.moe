package soundtracks

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxTracks = 9

// Get renders the soundtracks page.
func Get(ctx *aero.Context) string {
	tracks, err := arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft
	})

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching soundtracks", err)
	}

	arn.SortSoundTracksLatestFirst(tracks)

	if len(tracks) > maxTracks {
		tracks = tracks[:maxTracks]
	}

	return ctx.HTML(components.SoundTracks(tracks))
}
