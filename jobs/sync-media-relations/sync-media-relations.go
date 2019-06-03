package main

import (
	"fmt"
	"strings"

	"github.com/animenotifier/notify.moe/arn"

	"github.com/akyoto/color"
	"github.com/animenotifier/kitsu"
)

func main() {
	color.Yellow("Syncing media relations with Kitsu DB")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	kitsuMediaRelations := kitsu.StreamMediaRelations()
	relations := map[string]*arn.AnimeRelations{}

	for mediaRelation := range kitsuMediaRelations {
		// We only care about anime for now
		if mediaRelation.Relationships.Source.Data.Type != "anime" || mediaRelation.Relationships.Destination.Data.Type != "anime" {
			continue
		}

		relationType := strings.Replace(mediaRelation.Attributes.Role, "_", " ", -1)
		animeID := mediaRelation.Relationships.Source.Data.ID
		destinationAnimeID := mediaRelation.Relationships.Destination.Data.ID

		// Confirm that the anime IDs are valid
		if !arn.DB.Exists("Anime", animeID) {
			continue
		}

		if !arn.DB.Exists("Anime", destinationAnimeID) {
			continue
		}

		fmt.Printf(
			"%s %s has %s which is %s %s\n",
			mediaRelation.Relationships.Source.Data.Type,
			animeID,
			color.GreenString(relationType),
			mediaRelation.Relationships.Destination.Data.Type,
			destinationAnimeID,
		)

		// Add anime to the global map
		relationsList, found := relations[animeID]

		if !found {
			relationsList = &arn.AnimeRelations{
				AnimeID: animeID,
				Items:   []*arn.AnimeRelation{},
			}
			relations[animeID] = relationsList
		}

		relationsList.Items = append(relationsList.Items, &arn.AnimeRelation{
			AnimeID: destinationAnimeID,
			Type:    relationType,
		})

		// for _, item := range relationsList.Items {
		// 	fmt.Println("*", item.Type, item.AnimeID)
		// }
	}

	// Save relations map
	for _, animeRelations := range relations {
		animeRelations.Save()
	}
}
