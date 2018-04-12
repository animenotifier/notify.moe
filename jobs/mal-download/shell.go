package main

import (
	"flag"

	"github.com/aerogo/crawler"
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

		// Create crawler
		malCrawler := crawler.New(
			headers,
			delayBetweenRequests,
			1,
		)

		// Queue
		queue(anime, malCrawler)

		// Wait
		malCrawler.Wait()
		return true
	}

	return false
}
