package main

import (
	"fmt"

	"github.com/animenotifier/arn"
)

func main() {
	defer arn.Node.Close()

	for anime := range arn.StreamAnime() {
		relations, _ := arn.GetAnimeRelations(anime.ID)

		if relations == nil {
			relations := &arn.AnimeRelations{
				AnimeID: anime.ID,
				Items:   []*arn.AnimeRelation{},
			}

			arn.DB.Set("AnimeRelations", anime.ID, relations)
			fmt.Println(anime.ID, anime.Title.Canonical)
		}
	}
}
