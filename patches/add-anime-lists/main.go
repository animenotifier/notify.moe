package main

import (
	"fmt"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Adding empty anime lists to users who don't have one")

	// Get a stream of all users
	allUsers, err := arn.AllUsers()

	if err != nil {
		panic(err)
	}

	// Iterate over the stream
	for user := range allUsers {
		if user.AnimeList() == nil {
			fmt.Println(user.Nick)

			err := arn.DB.Set("AnimeList", user.ID, &arn.AnimeList{
				UserID: user.ID,
				Items:  make([]*arn.AnimeListItem, 0),
			})

			if err != nil {
				color.Red(err.Error())
			}
		}
	}

	color.Green("Finished.")
}
