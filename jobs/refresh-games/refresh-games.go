package main

import (
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Refreshing game information")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	for user := range arn.StreamUsers() {
		if user.Accounts.Osu.Nick != "" {
			osu(user)
		}

		if user.Accounts.Overwatch.BattleTag != "" {
			overwatch(user)
		}

		if user.Accounts.FinalFantasyXIV.Nick != "" && user.Accounts.FinalFantasyXIV.Server != "" {
			ffxiv(user)
		}
	}
}
