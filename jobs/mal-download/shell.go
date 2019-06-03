package main

import (
	"flag"

	"github.com/aerogo/crawler"
	"github.com/animenotifier/notify.moe/arn"
)

// Shell parameters
var objectType string
var objectID string
var newOnly bool

// Shell flags
func init() {
	flag.StringVar(&objectType, "type", "all", "all | anime | character")
	flag.StringVar(&objectID, "id", "", "ID of the notify.moe anime/character you want to refresh")
	flag.BoolVar(&newOnly, "new", false, "Skip existing entries and only download new ones")
	flag.Parse()
}

// InvokeShellArgs ...
func InvokeShellArgs() bool {
	if objectID != "" {
		// Create crawler
		malCrawler := crawler.New(
			headers,
			delayBetweenRequests,
			1,
		)

		switch objectType {
		case "anime":
			anime, err := arn.GetAnime(objectID)
			arn.PanicOnError(err)
			queueAnime(anime, malCrawler)

		case "character":
			character, err := arn.GetCharacter(objectID)
			arn.PanicOnError(err)
			queueCharacter(character, malCrawler)
		}

		// Wait
		malCrawler.Wait()
		return true
	}

	return false
}
