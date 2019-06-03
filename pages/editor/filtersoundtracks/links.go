package filtersoundtracks

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Links shows soundtracks without links.
func Links(ctx aero.Context) error {
	return editorList(
		ctx,
		"Soundtracks without links",
		func(track *arn.SoundTrack) bool {
			return len(track.Links) == 0
		},
		func(track *arn.SoundTrack) string {
			youtubeMedia := track.MediaByService("Youtube")

			if len(youtubeMedia) > 0 {
				return "https://song.link/https://youtu.be/" + youtubeMedia[0].ServiceID
			}

			return ""
		},
	)
}
