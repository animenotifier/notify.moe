package main

import (
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
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
