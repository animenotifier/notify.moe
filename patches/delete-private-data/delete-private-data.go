package main

import (
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Deleting private user data")
	defer arn.Node.Close()

	emailToUser := arn.DB.Collection("EmailToUser")

	testRecord, _ := emailToUser.Get("e.urbach@gmail.com")
	emailToUser.Clear()
	emailToUser.Set("e.urbach@gmail.com", testRecord)

	// Iterate over the stream
	count := 0

	for user := range arn.StreamUsers() {
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

		user.PushSubscriptions().Items = []*arn.PushSubscription{}
		user.PushSubscriptions().Save()

		// Save in DB
		user.Save()
	}

	color.Green("Finished.")
}
