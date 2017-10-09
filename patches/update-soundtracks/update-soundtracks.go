package main

import (
	"github.com/animenotifier/arn"
)

func main() {
	for track := range arn.MustStreamSoundTracks() {
		arn.PanicOnError(track.Save())
	}
}
