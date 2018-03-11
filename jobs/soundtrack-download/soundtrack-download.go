package main

import (
	"fmt"
	"time"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

const delayBetweenRequests = 1000

func main() {
	color.Yellow("Downloading soundtracks")
	defer arn.Node.Close()

	for track := range arn.StreamSoundTracks() {
		fmt.Println(track.Title)

		err := track.Download()

		if err != nil {
			color.Red(err.Error())
			continue
		}

		// Save the file information
		track.Save()

		// Delay a little
		time.Sleep(delayBetweenRequests)
	}

	color.Green("Finished.")
}
