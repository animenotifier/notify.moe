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
		color.Green("Please run this command with -confirm option if you really want to delete all pro subscriptions.")
		return
	}

	color.Yellow("Deleting all pro subscriptions")
	defer arn.Node.Close()

	for user := range arn.StreamUsers() {
		user.ProExpires = ""
		user.Save()
	}

	color.Green("Finished.")
}
