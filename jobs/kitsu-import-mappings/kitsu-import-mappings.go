package main

import (
	"fmt"

	"github.com/akyoto/color"
	"github.com/animenotifier/kitsu"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Importing Kitsu mappings")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	// Iterate over all mappings
	for mapping := range kitsu.StreamMappings() {
		sync(mapping)
	}
}

func sync(mapping *kitsu.Mapping) {
	// Skip mappings for anything that's not anime
	if mapping.Relationships.Item.Data.Type != "anime" {
		return
	}

	fmt.Printf("[MappingID: %s] [AnimeID: %s] %s %s\n", mapping.ID, mapping.Relationships.Item.Data.ID, color.YellowString(mapping.Attributes.ExternalSite), mapping.Attributes.ExternalID)
	arn.Kitsu.Set("Mapping", mapping.ID, mapping)
}
