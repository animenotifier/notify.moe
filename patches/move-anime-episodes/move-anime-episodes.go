package main

import (
	"github.com/animenotifier/arn"
)

func main() {
	for anime := range arn.MustStreamAnime() {
		arn.PanicOnError(arn.DB.Set("AnimeEpisodes", anime.ID, &arn.AnimeEpisodes{
			AnimeID: anime.ID,
			Items:   anime.Episodes,
		}))

		anime.Episodes = anime.Episodes[:0]
		anime.MustSave()
	}
}
