package main

import (
	"fmt"
	"strings"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Refreshing episode information for each anime.")
	defer arn.Node.Close()

	if InvokeShellArgs() {
		return
	}

	highPriority := []*arn.Anime{}
	mediumPriority := []*arn.Anime{}
	lowPriority := []*arn.Anime{}

	for anime := range arn.StreamAnime() {
		if anime.GetMapping("shoboi/anime") == "" {
			continue
		}

		// The rest gets sorted by airing status
		switch anime.Status {
		case "current":
			highPriority = append(highPriority, anime)
		case "upcoming":
			mediumPriority = append(mediumPriority, anime)
		default:
			lowPriority = append(lowPriority, anime)
		}
	}

	switch queue {
	case "high":
		refreshQueue(highPriority)

	case "medium":
		refreshQueue(mediumPriority)

	case "low":
		refreshQueue(lowPriority)

	default:
		color.Cyan("High priority queue (%d):", len(highPriority))
		refreshQueue(highPriority)

		color.Cyan("Medium priority queue (%d):", len(mediumPriority))
		refreshQueue(mediumPriority)

		color.Cyan("Low priority queue (%d):", len(lowPriority))
		refreshQueue(lowPriority)
	}

	color.Green("Finished.")
}

func refreshQueue(queue []*arn.Anime) {
	for _, anime := range queue {
		refresh(anime)
	}
}

func refresh(anime *arn.Anime) {
	fmt.Println(anime.ID, "|", anime.Title.Canonical, "|", anime.GetMapping("shoboi/anime"))

	episodeCount := len(anime.Episodes().Items)
	availableEpisodeCount := anime.Episodes().AvailableCount()

	err := anime.RefreshEpisodes()

	if err != nil {
		if strings.Contains(err.Error(), "missing a Shoboi ID") {
			return
		}

		color.Red(err.Error())
	} else {
		faint := color.New(color.Faint).SprintFunc()
		episodes := anime.Episodes()

		fmt.Println(faint(episodes.ListString()))
		fmt.Printf("+%d episodes | +%d available (%d total)\n", len(episodes.Items)-episodeCount, episodes.AvailableCount()-availableEpisodeCount, len(episodes.Items))
		println()
	}
}
