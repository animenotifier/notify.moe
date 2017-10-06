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

	for user := range arn.MustStreamUsers() {
		user.ProExpires = ""
		arn.PanicOnError(user.Save())
	}

	color.Green("Finished.")
}
