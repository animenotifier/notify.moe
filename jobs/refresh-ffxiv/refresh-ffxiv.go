package main

import (
	"time"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/arn/stringutils"
	"github.com/fatih/color"
)

var ticker = time.NewTicker(1100 * time.Millisecond)

func main() {
	color.Yellow("Refreshing FFXIV information")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	for user := range arn.StreamUsers() {
		if user.Accounts.FinalFantasyXIV.Nick == "" || user.Accounts.FinalFantasyXIV.Server == "" {
			continue
		}

		// Fetch new info
		err := user.RefreshFFXIVInfo()

		if err != nil {
			color.Red(err.Error())
			continue
		}

		// Log it
		stringutils.PrettyPrint(user.Accounts.FinalFantasyXIV)

		// Save in database
		user.Save()

		// Wait for rate limiter
		<-ticker.C
	}
}
