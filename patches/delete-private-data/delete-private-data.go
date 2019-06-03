package main

import (
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
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
		user.Accounts.Twitter.ID = ""
		user.Accounts.Twitter.Nick = ""
		user.Location = &arn.Location{}
		user.BirthDay = ""

		user.PushSubscriptions().Items = []*arn.PushSubscription{}
		user.PushSubscriptions().Save()

		// Save in DB
		user.Save()
	}

	color.Green("Finished.")
}
