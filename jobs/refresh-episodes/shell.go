package main

import (
	"flag"

	"github.com/animenotifier/notify.moe/arn"
)

// Shell parameters
var animeID string
var queue string

// Shell flags
func init() {
	flag.StringVar(&animeID, "id", "", "ID of the anime you want to refresh")
	flag.StringVar(&queue, "queue", "", "Queue type you want to refresh (high, medium, low)")
	flag.Parse()
}

// InvokeShellArgs ...
func InvokeShellArgs() bool {
	if animeID != "" {
		anime, err := arn.GetAnime(animeID)

		if err != nil {
			panic(err)
		}

		refresh(anime)
		return true
	}

	return false
}
