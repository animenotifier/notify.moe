package soundtracks

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// FilterByTag renders the best soundtracks filtered by tag.
func FilterByTag(ctx aero.Context) error {
	tag := ctx.Get("tag")

	// Fetch all eligible tracks
	tracks := arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && len(track.Media) > 0 && track.HasTag(tag)
	})

	// Sort the tracks by number of likes
	arn.SortSoundTracksPopularFirst(tracks)

	// Render
	return render(ctx, tracks)
}
