package main

import (
	"flag"
	"path"

	"github.com/animenotifier/arn"
)

// Shell parameters
var objectType string
var objectID string

// Shell flags
func init() {
	flag.StringVar(&objectType, "type", "all", "all | anime | character")
	flag.StringVar(&objectID, "id", "", "ID of the notify.moe anime/character you want to refresh")
	flag.Parse()
}

// InvokeShellArgs ...
func InvokeShellArgs() bool {
	if objectID != "" {
		switch objectType {
		case "anime":
			anime, err := arn.GetAnime(objectID)
			arn.PanicOnError(err)

			if anime.GetMapping("myanimelist/anime") == "" {
				panic("No MAL ID")
			}

			readAnimeFile(path.Join(arn.Root, "jobs", "mal-download", "anime", anime.GetMapping("myanimelist/anime")+".html.gz"))

		case "character":
			character, err := arn.GetCharacter(objectID)
			arn.PanicOnError(err)

			if character.GetMapping("myanimelist/character") == "" {
				panic("No MAL ID")
			}

			readCharacterFile(path.Join(arn.Root, "jobs", "mal-download", "character", character.GetMapping("myanimelist/character")+".html.gz"))
		}

		return true
	}

	return false
}
