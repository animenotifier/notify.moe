package main

import (
	"fmt"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Adding push subscriptions to users who don't have one")

	// Get a stream of all users
	allUsers, err := arn.StreamUsers()

	if err != nil {
		panic(err)
	}

	// Iterate over the stream
	for user := range allUsers {
		exists, err := arn.DB.Exists("PushSubscriptions", user.ID)

		if err == nil && !exists {
			fmt.Println(user.Nick)

			err := arn.DB.Set("PushSubscriptions", user.ID, &arn.PushSubscriptions{
				UserID: user.ID,
				Items:  make([]*arn.PushSubscription, 0),
			})

			if err != nil {
				color.Red(err.Error())
			}
		}
	}

	color.Green("Finished.")
}
