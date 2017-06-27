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
			arn.DB.Set("GoogleToUser", user.Accounts.Google.ID, &arn.GoogleToUser{
				ID:     user.Accounts.Google.ID,
				UserID: user.ID,
			})
		}
	}

	color.Green("Finished.")
}
