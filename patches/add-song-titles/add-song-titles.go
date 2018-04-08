package main

import (
	"github.com/animenotifier/arn"
)

func main() {
	defer arn.Node.Close()

	for track := range arn.StreamSoundTracks() {
		if arn.ContainsUnicodeLetters(track.Title) {
			track.NewTitle.Native = track.Title
		} else {
			track.NewTitle.Canonical = track.Title
		}

		track.Save()
	}
}
