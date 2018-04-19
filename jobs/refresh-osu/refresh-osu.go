package main

import (
	"time"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/arn/stringutils"
	"github.com/fatih/color"
)

var ticker = time.NewTicker(500 * time.Millisecond)

func main() {
	color.Yellow("Refreshing osu information")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	for user := range arn.StreamUsers() {
		if user.Accounts.Osu.Nick == "" {
			continue
		}

		// Fetch new info
		err := user.RefreshOsuInfo()

		if err != nil {
			color.Red(err.Error())
			continue
		}

		// Log it
		stringutils.PrettyPrint(user.Accounts.Osu)

		// Save in database
		user.Save()

		// Wait for rate limiter
		<-ticker.C
	}
}
