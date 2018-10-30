package main

import (
	"flag"
	"path"

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

		readAnimeFile(path.Join(arn.Root, "jobs/mal-download/anime", "anime-"+anime.GetMapping("myanimelist/anime")+".html"))
		return true
	}

	return false
}
