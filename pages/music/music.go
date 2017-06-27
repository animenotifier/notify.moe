package music

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxTracks = 10

// Get renders the music page.
func Get(ctx *aero.Context) string {
	tracks, err := arn.AllSoundTracks()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching soundtracks", err)
	}

	arn.SortSoundTracksLatestFirst(tracks)

	if len(tracks) > maxTracks {
		tracks = tracks[:maxTracks]
	}

	return ctx.HTML(components.Music(tracks))
}
