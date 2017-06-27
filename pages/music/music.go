package music

import (
	"net/http"
	"sort"

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

	sort.Slice(tracks, func(i, j int) bool {
		return tracks[i].Created > tracks[j].Created
	})

	if len(tracks) > maxTracks {
		tracks = tracks[:maxTracks]
	}

	return ctx.HTML(components.Music(tracks))
}
