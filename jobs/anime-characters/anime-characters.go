package main

import (
	"fmt"
	"time"

	"github.com/akyoto/color"
	"github.com/animenotifier/arn"
)

func main() {
	color.Yellow("Refreshing anime characters...")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	rateLimiter := time.NewTicker(500 * time.Millisecond)

	for anime := range arn.StreamAnime() {
		<-rateLimiter.C

		chars, err := anime.RefreshAnimeCharacters()

		if err != nil {
			color.Red(err.Error())
			continue
		}

		fmt.Printf("%s %s (%d characters)\n", anime.ID, anime.Title.Canonical, len(chars.Items))
	}
}
