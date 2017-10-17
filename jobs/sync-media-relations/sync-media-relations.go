package main

import (
	"fmt"
	"strings"

	"github.com/animenotifier/kitsu"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Syncing media relations with Kitsu DB")

	kitsuMediaRelations := kitsu.StreamMediaRelations()

	for mediaRelation := range kitsuMediaRelations {
		// We only care about anime for now
		if mediaRelation.Relationships.Source.Data.Type != "anime" || mediaRelation.Relationships.Destination.Data.Type != "anime" {
			continue
		}

		relationType := strings.Replace(mediaRelation.Attributes.Role, "_", " ", -1)

		fmt.Printf(
			"%s %s has a %s which is %s %s\n",
			mediaRelation.Relationships.Source.Data.Type,
			mediaRelation.Relationships.Source.Data.ID,
			color.GreenString(relationType),
			mediaRelation.Relationships.Destination.Data.Type,
			mediaRelation.Relationships.Destination.Data.ID,
		)
	}

	color.Green("Finished.")
}
