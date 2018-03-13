package soundtracks

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/animenotifier/notify.moe/utils/infinitescroll"
)

const maxTracks = 12

// Latest renders the latest soundtracks.
func Latest(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	index, _ := ctx.GetInt("index")

	// Fetch all eligible tracks
	allTracks := fetchAll()

	// Sort the tracks by date
	arn.SortSoundTracksLatestFirst(allTracks)

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
