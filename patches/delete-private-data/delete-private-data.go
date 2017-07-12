package main

import (
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Deleting private user data")

	// Get a stream of all users
	allUsers, err := arn.StreamUsers()

	if err != nil {
		panic(err)
	}

	arn.DB.DeleteTable("EmailToUser")
	arn.DB.DeleteTable("GoogleToUser")

	// Iterate over the stream
	count := 0
	for user := range allUsers {
		count++
		println(count, user.Nick)

		// Delete private data
		user.Email = ""
		user.Gender = ""
		user.FirstName = ""
		user.LastName = ""
		user.IP = ""
		user.Accounts.Facebook.ID = ""
		user.Accounts.Google.ID = ""
		user.AgeRange = arn.UserAgeRange{}
		user.Location = arn.UserLocation{}

		// Save in DB
		user.Save()
	}

	color.Green("Finished.")
}
