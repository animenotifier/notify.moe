package main

import (
	"fmt"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	defer arn.Node.Close()

	for mapping := range arn.StreamKitsuMappings() {
		fmt.Printf(
			"Kitsu Anime %s: %s mapped to %s\n",
			mapping.Relationships.Item.Data.ID,
			color.YellowString(mapping.Attributes.ExternalSite),
			color.GreenString(mapping.Attributes.ExternalID),
		)
	}
}
