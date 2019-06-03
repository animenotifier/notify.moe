package main

import (
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Deleting all AMV tags")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	for amv := range arn.StreamAMVs() {
		amv.Tags = []string{}
		amv.Save()
	}
}
