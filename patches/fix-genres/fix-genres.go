package main

import (
	"fmt"
	"strings"

	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	defer arn.Node.Close()

	for anime := range arn.StreamAnime() {
		for i := 0; i < len(anime.Genres); i++ {
			old := anime.Genres[i]
			genre := strings.TrimSpace(anime.Genres[i])

			if genre != old {
				fmt.Println(anime.Title.Canonical, genre)
			}

			anime.Genres[i] = genre
		}

		anime.Save()
	}
}
