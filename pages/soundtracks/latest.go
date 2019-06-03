package soundtracks

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Latest renders the latest soundtracks.
func Latest(ctx aero.Context) error {
	// Fetch all eligible tracks
	tracks := fetchAll()

	// Sort the tracks by date
	arn.SortSoundTracksLatestFirst(tracks)

	// Render
	return render(ctx, tracks)
}
