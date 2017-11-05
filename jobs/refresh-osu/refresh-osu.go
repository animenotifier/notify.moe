package main

import (
	"time"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Refreshing osu information")
	defer arn.Node.Close()

	ticker := time.NewTicker(500 * time.Millisecond)

	for user := range arn.StreamUsers() {
		// Get osu info
		if user.RefreshOsuInfo() == nil {
			arn.PrettyPrint(user.Accounts.Osu)

			// Fetch user again to prevent writing old data
			updatedUser, _ := arn.GetUser(user.ID)
			updatedUser.Accounts.Osu = user.Accounts.Osu
			updatedUser.Save()
		}

		// Wait for rate limiter
		<-ticker.C
	}

	color.Green("Finished.")
}
