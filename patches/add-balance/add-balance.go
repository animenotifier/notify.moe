package main

import (
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Adding balance to all users")

	// Iterate over the stream
	for user := range arn.StreamUsers() {
		user.Balance += 100000
		user.Save()
	}

	color.Green("Finished.")
}
