package main

import (
	"flag"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
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
		color.Green("Please run this command with -confirm option if you really want to reset the balance of all users.")
		return
	}

	color.Yellow("Resetting balance of all users to 0")
	defer arn.Node.Close()

	// Iterate over the stream
	for user := range arn.StreamUsers() {
		user.Balance = 0
		user.Save()
	}

	color.Green("Finished.")
}
