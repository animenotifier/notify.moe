package main

import (
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
	"github.com/jinzhu/copier"
)

func main() {
	color.Yellow("Updating user struct")

	// Iterate over the stream
	for user := range arn.MustStreamUsers() {
		newUser := &arn.UserNew{}

		copier.Copy(newUser, user)
		newUser.Avatar.Extension = user.Avatar

		// Save in DB
		arn.PanicOnError(arn.DB.Set("User", user.ID, newUser))
	}

	color.Green("Finished.")
}
