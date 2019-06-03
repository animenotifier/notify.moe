package main

import (
	"fmt"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Downloading soundtracks")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	for track := range arn.StreamSoundTracks() {
		if track.IsDraft {
			continue
		}

		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Println(track.Title)

		err := track.Download()

		if err != nil {
			color.Red(err.Error())
			continue
		} else {
			color.Green("Downloaded %s!", track.File)
		}

		// Save the file information
		track.Save()
	}
}
