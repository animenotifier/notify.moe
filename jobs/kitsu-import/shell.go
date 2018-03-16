package main

import (
	"errors"
	"flag"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/kitsu"
	"github.com/fatih/color"
)

// Shell parameters
var animeID string
var verbose bool

// Shell flags
func init() {
	flag.StringVar(&animeID, "id", "", "ID of the anime that you want to refresh")
	flag.BoolVar(&verbose, "v", false, "Verbose output")
	flag.Parse()
}

// InvokeShellArgs ...
func InvokeShellArgs() bool {
	if animeID != "" {
		kitsuAnime, err := kitsu.GetAnime(animeID)

		if err != nil {
			panic(err)
		}

		if kitsuAnime.ID != animeID {
			panic(errors.New("Anime ID is not the same"))
		}

		anime := importKitsuAnime(kitsuAnime)

		if verbose {
			color.Cyan("Kitsu:")
			arn.PrettyPrint(kitsuAnime)

			color.Cyan("ARN:")
			arn.PrettyPrint(anime)
		}

		return true
	}

	return false
}
