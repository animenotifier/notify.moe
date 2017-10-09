package main

import (
	"flag"
	"fmt"

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
		color.Green("Please run this command with -confirm option.")
		return
	}

	color.Yellow("Resetting all inventories")

	// Get a stream of all users
	allUsers, err := arn.StreamUsers()
	arn.PanicOnError(err)

	// Iterate over the stream
	for user := range allUsers {
		fmt.Println(user.Nick)

		inventory := arn.NewInventory(user.ID)
		err = inventory.Save()

		if err != nil {
			color.Red(err.Error())
		}
	}

	color.Green("Finished.")
}
