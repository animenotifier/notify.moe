package soundtrack

import (
	"math/rand"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Random returns a random soundtrack.
func Random(ctx *aero.Context) string {
	tracks := arn.AllSoundTracks()
	index := rand.Intn(len(tracks))
	track := tracks[index]

	return ctx.JSON(track)
}

// Next returns the next soundtrack for the audio player.
func Next(ctx *aero.Context) string {
	tracks := arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return track.File != ""
	})

	index := rand.Intn(len(tracks))
	track := tracks[index]

	return ctx.JSON(track)
}
