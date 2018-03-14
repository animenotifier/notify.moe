package soundtracks

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

const (
	tracksFirstLoad = 12
	tracksPerScroll = 3
)

// Latest renders the latest soundtracks.
func Latest(ctx *aero.Context) string {
	// Fetch all eligible tracks
	tracks := fetchAll()

	// Sort the tracks by date
	arn.SortSoundTracksLatestFirst(tracks)

	// Render
	return render(ctx, tracks)
}
