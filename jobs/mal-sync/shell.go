package main

import (
	"flag"

	"github.com/animenotifier/arn"
)

// Shell parameters
var animeID string

// Shell flags
func init() {
	flag.StringVar(&animeID, "id", "", "ID of the notify.moe anime you want to refresh")
	flag.Parse()
}

// InvokeShellArgs ...
func InvokeShellArgs() bool {
	if animeID != "" {
		anime, err := arn.GetAnime(animeID)

		if err != nil {
			panic(err)
		}

		if anime.GetMapping("myanimelist/anime") == "" {
			panic("No MAL ID")
		}

		syncAnime(anime, anime.GetMapping("myanimelist/anime"))
		return true
	}

	return false
}
