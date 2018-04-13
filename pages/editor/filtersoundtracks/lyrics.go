package filtersoundtracks

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Lyrics shows soundtracks without lyrics.
func Lyrics(ctx *aero.Context) string {
	return editorList(
		ctx,
		"Soundtracks without lyrics",
		func(track *arn.SoundTrack) bool {
			return !track.HasLyrics()
		},
		func(track *arn.SoundTrack) string {
			return "https://www.google.com/search?q=" + track.Title.String() + " lyrics site:animelyrics.com"
		},
	)
}
