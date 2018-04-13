package filtersoundtracks

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Tags shows soundtracks with less than 3 tags.
func Tags(ctx *aero.Context) string {
	return editorList(
		ctx,
		"Soundtracks with less than 3 tags",
		func(track *arn.SoundTrack) bool {
			return len(track.Tags) < 3
		},
		nil,
	)
}
