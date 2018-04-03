package main

import (
	"fmt"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Adding Kitsu mappings")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	for mapping := range arn.StreamKitsuMappings() {
		if mapping.Relationships.Item.Data.Type != "anime" {
			continue
		}

		if mapping.Attributes.ExternalSite != "trakt" && mapping.Attributes.ExternalSite != "anidb" {
			continue
		}

		anime, _ := arn.GetAnime(mapping.Relationships.Item.Data.ID)

		if anime == nil {
			continue
		}

		fmt.Println(anime.ID, mapping.Attributes.ExternalSite, mapping.Attributes.ExternalID)

		anime.ImportKitsuMapping(mapping)
		anime.Save()
	}
}
