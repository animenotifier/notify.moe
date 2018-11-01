package main

import (
	"fmt"
	"time"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/arn/stringutils"
	"github.com/fatih/color"
)

var tickerOsu = time.NewTicker(500 * time.Millisecond)

func osu(user *arn.User) {
	fmt.Println("[Osu]", user.Nick, user.Accounts.Osu.Nick)

	// Fetch new info
	err := user.RefreshOsuInfo()

	if err != nil {
		color.Red(err.Error())
		return
	}

	// Log it
	stringutils.PrettyPrint(user.Accounts.Osu)

	// Save in database
	user.Save()

	// Wait for rate limiter
	<-tickerOsu.C
}
