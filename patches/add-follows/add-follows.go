package main

import (
	"fmt"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Adding user follows to users who don't have one")

	// Get a stream of all users
	allUsers, err := arn.StreamUsers()

	if err != nil {
		panic(err)
	}

	// Iterate over the stream
	for user := range allUsers {
		// exists, err := arn.DB.Exists("UserFollows", user.ID)

		// if err != nil || exists {
		// 	continue
		// }

		fmt.Println(user.Nick)

		follows := &arn.UserFollows{}
		follows.UserID = user.ID
		follows.Items = user.Following

		err = arn.DB.Set("UserFollows", follows.UserID, follows)

		if err != nil {
			color.Red(err.Error())
		}
	}

	color.Green("Finished.")
}
