package soundtracks

import (
	"github.com/animenotifier/arn"
)

// fetchAll returns all soundtracks
func fetchAll() []*arn.SoundTrack {
	return arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && track.File != ""
	})
}
