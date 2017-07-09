package main

import (
	"sort"
	"strconv"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/twist"
	"github.com/fatih/color"
)

func main() {
	// Replace this with ID list from twist.moe later
	animeIDs := []string{
		"13274",
		"10902",
	}

	for _, animeID := range animeIDs {
		feed, err := twist.GetFeedByKitsuID(animeID)

		if err != nil {
			color.Red("Error querying ID %s: %v", animeID, err)
			continue
		}

		episodes := feed.Episodes

		// Sort by episode number
		sort.Slice(episodes, func(a, b int) bool {
			epsA, _ := strconv.Atoi(episodes[a].Number)
			epsB, _ := strconv.Atoi(episodes[b].Number)
			return epsA < epsB
		})

		arn.PrettyPrint(episodes)
	}
}
