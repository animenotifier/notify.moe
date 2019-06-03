package main

import (
	"fmt"
	"strings"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Checking Kitsu for anilist mappings")
	defer arn.Node.Close()

	confirmed := 0
	added := 0
	conflicted := 0

	finder := arn.NewAnimeFinder("kitsu/anime")

	for mapping := range arn.StreamKitsuMappings() {
		if mapping.Relationships.Item.Data.Type != "anime" {
			continue
		}

		if mapping.Attributes.ExternalSite != "anilist" {
			continue
		}

		externalID := mapping.Attributes.ExternalID
		externalID = strings.TrimPrefix(externalID, "anime/")

		anime := finder.GetAnime(mapping.Relationships.Item.Data.ID)

		if anime == nil {
			continue
		}

		currentID := anime.GetMapping("anilist/anime")

		if currentID == "" {
			added++
			// color.Yellow("Added: %s (%v) on %s is %s", anime.ID, anime, mapping.Attributes.ExternalSite, externalID)
			// color.Yellow("Added:\n * https://notify.moe/anime/%s\n * https://anilist.co/anime/%s\n\n", anime.ID, externalID)
		} else {
			if currentID == externalID {
				confirmed++
				// color.Green("Confirmed: %s (%v) on %s is %s", anime.ID, anime, mapping.Attributes.ExternalSite, externalID)
			} else if currentID != externalID {
				conflicted++
				// color.Red("Conflict: %s (%v) on %s is %s but current value is %s", anime.ID, anime, mapping.Attributes.ExternalSite, externalID, currentID)
				color.Red("Conflict (#%d):\n * https://notify.moe/anime/%s\n * https://anilist.co/anime/%s (current)\n * https://anilist.co/anime/%s (suggested)\n\n", conflicted, anime.ID, externalID, currentID)
			}
		}

		// anime.SetMapping("anilist/anime", externalID)
	}

	fmt.Printf("%d confirmed, %d added, %d conflicted\n", confirmed, added, conflicted)
}
