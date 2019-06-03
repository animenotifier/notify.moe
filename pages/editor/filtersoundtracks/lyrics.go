package filtersoundtracks

import (
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// MissingLyrics shows soundtracks without lyrics.
func MissingLyrics(ctx aero.Context) error {
	return editorList(
		ctx,
		"Soundtracks without lyrics",
		func(track *arn.SoundTrack) bool {
			if !track.HasTag("vocal") {
				return false
			}

			return !track.HasLyrics()
		},
		func(track *arn.SoundTrack) string {
			return "https://www.google.com/search?q=" + track.Title.String() + " lyrics site:animelyrics.com"
		},
	)
}

// UnalignedLyrics shows soundtracks with unaligned lyrics.
func UnalignedLyrics(ctx aero.Context) error {
	return editorList(
		ctx,
		"Soundtracks with unaligned lyrics",
		func(track *arn.SoundTrack) bool {
			if !track.HasTag("vocal") {
				return false
			}

			return track.Lyrics.Native != "" && track.Lyrics.Romaji != "" && strings.Count(track.Lyrics.Native, "\n") != strings.Count(track.Lyrics.Romaji, "\n")
		},
		nil,
	)
}
