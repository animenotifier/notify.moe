package main

import (
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Updating user references")
	defer arn.Node.Close()

	arn.DB.Clear("NickToUser")
	arn.DB.Clear("EmailToUser")
	arn.DB.Clear("GoogleToUser")
	arn.DB.Clear("FacebookToUser")

	// Iterate over the stream
	count := 0

	for user := range arn.StreamUsers() {
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
