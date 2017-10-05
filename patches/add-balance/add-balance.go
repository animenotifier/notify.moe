package main

import (
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Adding balance to all users")

	// Get a stream of all users
	allUsers, err := arn.StreamUsers()
	arn.PanicOnError(err)

	// Iterate over the stream
	for user := range allUsers {
		user.Balance += 100000
		arn.PanicOnError(user.Save())
	}

	color.Green("Finished.")
}
