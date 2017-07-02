package main

import (
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Updating user references")

	arn.DB.DeleteTable("NickToUser")
	arn.DB.DeleteTable("EmailToUser")
	arn.DB.DeleteTable("GoogleToUser")
	arn.DB.DeleteTable("FacebookToUser")

	// Get a stream of all users
	allUsers, err := arn.StreamUsers()

	if err != nil {
		panic(err)
	}

	// Iterate over the stream
	count := 0
	for user := range allUsers {
		count++
		println(count, user.Nick)

		user.ForceSetNick(user.Nick)

		if user.Email != "" {
			user.SetEmail(user.Email)
		}

		if user.Accounts.Google.ID != "" {
			user.ConnectGoogle(user.Accounts.Google.ID)
		}

		if user.Accounts.Facebook.ID != "" {
			user.ConnectFacebook(user.Accounts.Facebook.ID)
		}
	}

	color.Green("Finished.")
}
