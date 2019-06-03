package soundtrack

import (
	"math/rand"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Random returns a random soundtrack.
func Random(ctx aero.Context) error {
	tracks := arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft
	})

	index := rand.Intn(len(tracks))
	track := tracks[index]

	return ctx.JSON(track)
}

// Next returns the next soundtrack for the audio player.
func Next(ctx aero.Context) error {
	tracks := arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && track.File != ""
	})

	index := rand.Intn(len(tracks))
	track := tracks[index]

	return ctx.JSON(track)
}
