package soundtracks

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils/infinitescroll"
)

const (
	tracksFirstLoad = 12
	tracksPerScroll = 9
)

// render renders the soundracks page with the given tracks.
func render(ctx aero.Context, allTracks []*arn.SoundTrack) error {
	user := arn.GetUserFromContext(ctx)
	index, _ := ctx.GetInt("index")
	tag := ctx.Get("tag")

	// Slice the part that we need
	tracks := allTracks[index:]
	maxLength := tracksFirstLoad

	if index > 0 {
		maxLength = tracksPerScroll
	}

	if len(tracks) > maxLength {
		tracks = tracks[:maxLength]
	}

	// Next index
	nextIndex := infinitescroll.NextIndex(ctx, len(allTracks), maxLength, index)

	// In case we're scrolling, send soundtracks only (without the page frame)
	if index > 0 {
		return ctx.HTML(components.SoundTracksScrollable(tracks, user))
	}

	// Otherwise, send the full page
	return ctx.HTML(components.SoundTracks(tracks, nextIndex, tag, user))
}
