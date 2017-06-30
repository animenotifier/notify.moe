package main

import (
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Refreshing track titles")

	// Get a stream of all soundtracks
	soundtracks, err := arn.StreamSoundTracks()

	if err != nil {
		panic(err)
	}

	// Iterate over the stream
	for track := range soundtracks {
		sync(track)
	}

	color.Green("Finished.")
}

func sync(track *arn.SoundTrack) {
	for _, media := range track.Media {
		media.RefreshMetaData()
		println(media.Service, media.Title)
	}

	err := track.Save()

	if err != nil {
		panic(err)
	}
}
