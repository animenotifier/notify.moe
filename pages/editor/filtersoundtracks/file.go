package filtersoundtracks

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// File shows soundtracks without an audio file.
func File(ctx aero.Context) error {
	return editorList(
		ctx,
		"Soundtracks without an audio file",
		func(track *arn.SoundTrack) bool {
			return track.File == ""
		},
		nil,
	)
}
