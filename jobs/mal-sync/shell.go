package main

import (
	"flag"

	"github.com/animenotifier/notify.moe/arn"
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

			syncAnime(anime, anime.GetMapping("myanimelist/anime"))

		case "character":
			character, err := arn.GetCharacter(objectID)
			arn.PanicOnError(err)

			if character.GetMapping("myanimelist/character") == "" {
				panic("No MAL ID")
			}

			syncCharacter(character, character.GetMapping("myanimelist/character"))
		}

		return true
	}

	return false
}
