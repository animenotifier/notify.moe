package main

import (
	"fmt"
	"time"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/arn/stringutils"
)

var tickerOW = time.NewTicker(1100 * time.Millisecond)

func overwatch(user *arn.User) {
	fmt.Println("[Overwatch]", user.Nick, user.Accounts.Overwatch.BattleTag)

	// Fetch new info
	err := user.RefreshOverwatchInfo()

	if err != nil {
		color.Red(err.Error())
		return
	}

	// Log it
	stringutils.PrettyPrint(user.Accounts.Overwatch)

	// Save in database
	user.Save()

	// Wait for rate limiter
	<-tickerOW.C
}
