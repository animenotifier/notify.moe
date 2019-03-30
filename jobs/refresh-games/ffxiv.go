package main

import (
	"fmt"
	"time"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/arn/stringutils"
	"github.com/blitzprog/color"
)

var tickerFFXIV = time.NewTicker(1100 * time.Millisecond)

func ffxiv(user *arn.User) {
	fmt.Println("[FFXIV]", user.Nick, user.Accounts.FinalFantasyXIV.Nick, user.Accounts.FinalFantasyXIV.Server)

	// Fetch new info
	err := user.RefreshFFXIVInfo()

	if err != nil {
		color.Red(err.Error())
		return
	}

	// Log it
	stringutils.PrettyPrint(user.Accounts.FinalFantasyXIV)

	// Save in database
	user.Save()

	// Wait for rate limiter
	<-tickerFFXIV.C
}
