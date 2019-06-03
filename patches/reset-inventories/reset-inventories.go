package main

import (
	"flag"
	"fmt"

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
		color.Green("Please run this command with -confirm option.")
		return
	}

	color.Yellow("Resetting all inventories")
	defer arn.Node.Close()

	// Iterate over the stream
	for user := range arn.StreamUsers() {
		fmt.Println(user.Nick)

		inventory := arn.NewInventory(user.ID)
		inventory.Save()
	}

	color.Green("Finished.")
}
