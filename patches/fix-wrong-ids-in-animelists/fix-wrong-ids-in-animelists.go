package main

import (
	"fmt"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Fixing anime IDs in anime lists")

	defer color.Green("Finished")
	defer arn.Node.Close()

	finder := arn.NewAnimeFinder("kitsu/anime")

	for animeList := range arn.StreamAnimeLists() {
		modified := false

		for _, item := range animeList.Items {
			anime := item.Anime()

			if anime != nil {
				continue
			}

			anime = finder.GetAnime(item.AnimeID)

			if anime != nil {
				item.AnimeID = anime.ID
				modified = true
			}

			fmt.Println(item.AnimeID, anime, animeList.User().Nick)
		}

		if modified {
			animeList.Save()
		}
	}
}
