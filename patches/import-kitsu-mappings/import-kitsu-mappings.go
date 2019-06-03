package main

import (
	"fmt"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Adding Kitsu mappings")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	finder := arn.NewAnimeFinder("kitsu/anime")

	for mapping := range arn.StreamKitsuMappings() {
		if mapping.Relationships.Item.Data.Type != "anime" {
			continue
		}

		if mapping.Attributes.ExternalSite != "trakt" && mapping.Attributes.ExternalSite != "anidb" {
			continue
		}

		anime := finder.GetAnime(mapping.Relationships.Item.Data.ID)

		if anime == nil {
			continue
		}

		fmt.Println(anime.ID, mapping.Attributes.ExternalSite, mapping.Attributes.ExternalID)

		anime.ImportKitsuMapping(mapping)
		anime.Save()
	}
}
