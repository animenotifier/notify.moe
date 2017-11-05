package main

import (
	"fmt"

	"github.com/animenotifier/arn"
)

func main() {
	for anime := range arn.StreamAnime() {
		if anime.NSFW == 1 || anime.Status != "current" || anime.StartDate == "" || anime.StartDate < "2017-09" || anime.StartDate > "2017-10-17" {
			continue
		}

		fmt.Printf("* [%s](/anime/%s)\n", anime.Title.Canonical, anime.ID)
	}
}
