package main

import (
	"fmt"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Adding inventories to users who don't have one")

	// Get a stream of all users
	allUsers, err := arn.StreamUsers()
	arn.PanicOnError(err)

	// Iterate over the stream
	for user := range allUsers {
		// exists, err := arn.DB.Exists("Inventory", user.ID)

		// if err != nil || exists {
		// 	continue
		// }

		fmt.Println(user.Nick)

		inventory := arn.NewInventory(user.ID)

		// TEST
		inventory.AddItem("anime-support-ticket", 50)
		inventory.AddItem("pro-account-24", 30)

		err = arn.DB.Set("Inventory", inventory.UserID, inventory)

		if err != nil {
			color.Red(err.Error())
		}
	}

	color.Green("Finished.")
}
