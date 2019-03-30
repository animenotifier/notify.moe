package main

import (
	"github.com/animenotifier/arn"
	"github.com/blitzprog/color"
)

func main() {
	color.Yellow("Resetting location privacy to enabled")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	for user := range arn.StreamUsers() {
		settings := user.Settings()
		settings.Privacy.ShowLocation = true
		settings.Save()
	}
}
