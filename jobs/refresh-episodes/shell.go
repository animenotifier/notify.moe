package main

import (
	"flag"

	"github.com/animenotifier/arn"
)

// Shell parameters
var animeID string

// Shell flags
func init() {
	flag.StringVar(&animeID, "id", "", "ID of the anime you want to refresh")
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
