package main

import (
	"time"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Refreshing osu information")

	ticker := time.NewTicker(500 * time.Millisecond)

	for user := range arn.MustStreamUsers() {
		// Get osu info
		if user.RefreshOsuInfo() == nil {
			arn.PrettyPrint(user.Accounts.Osu)
			user.Save()
		}

		// Wait for rate limiter
		<-ticker.C
	}

	color.Green("Finished.")
}
