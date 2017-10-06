package main

import (
	"flag"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

// Shell parameters
var confirmed bool

// Shell flags
func init() {
	flag.BoolVar(&confirmed, "confirm", false, "Confirm that you really want to execute this.")
	flag.Parse()
}

func main() {
	if !confirmed {
		color.Green("Please run this command with -confirm option if you really want to delete all pro subscriptions.")
		return
	}

	color.Yellow("Deleting all pro subscriptions")

	// Get a stream of all users
	allUsers, err := arn.StreamUsers()
	arn.PanicOnError(err)

	// Iterate over the stream
	for user := range allUsers {
		user.Balance = 0
		arn.PanicOnError(user.Save())
	}

	color.Green("Finished.")
}
