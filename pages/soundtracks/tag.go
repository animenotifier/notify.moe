package soundtracks

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/animenotifier/notify.moe/utils/infinitescroll"
)

// FilterByTag renders the best soundtracks filtered by tag.
func FilterByTag(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	tag := ctx.Get("tag")
	index, _ := ctx.GetInt("index")

	// Fetch all eligible tracks
	allTracks := arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && len(track.Media) > 0 && track.HasTag(tag)
	})

	// Sort the tracks by number of likes
	arn.SortSoundTracksPopularFirst(allTracks)

	// Slice the part that we need
	tracks := allTracks[index:]

	if len(tracks) > maxTracks {
		tracks = tracks[:maxTracks]
	}

	// Next index
	nextIndex := infinitescroll.NextIndex(ctx, len(allTracks), maxTracks, index)

	// In case we're scrolling, send soundtracks only (without the page frame)
	if index > 0 {
		return ctx.HTML(components.SoundTracksScrollable(tracks, user))
	}

	// Otherwise, send the full page
	return ctx.HTML(components.SoundTracks(tracks, nextIndex, "", user))
}
