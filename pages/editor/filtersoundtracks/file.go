package filtersoundtracks

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// File shows soundtracks without an audio file.
func File(ctx *aero.Context) string {
	return editorList(
		ctx,
		"Soundtracks without an audio file",
		func(track *arn.SoundTrack) bool {
			return track.File == ""
		},
		nil,
	)
}
