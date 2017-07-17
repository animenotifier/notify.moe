package main

import (
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Updating user struct")

	// // Iterate over the stream
	// for user := range arn.MustStreamUsers() {
	// 	newUser := &arn.UserNew{}

	// 	copier.Copy(newUser, user)
	// 	newUser.Avatar.Extension = user.Avatar

	// 	// Save in DB
	// 	arn.PanicOnError(arn.DB.Set("User", user.ID, newUser))
	// }

	color.Green("Finished.")
}
