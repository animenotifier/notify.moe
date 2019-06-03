package main

import (
	"time"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/shoboi"
)

func main() {
	color.Yellow("Syncing Shoboi Anime")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	// Priority queues
	highPriority := []*arn.Anime{}
	mediumPriority := []*arn.Anime{}
	lowPriority := []*arn.Anime{}

	for anime := range arn.StreamAnime() {
		if anime.GetMapping("shoboi/anime") != "" {
			continue
		}

		switch anime.Status {
		case "current":
			highPriority = append(highPriority, anime)
		case "upcoming":
			mediumPriority = append(mediumPriority, anime)
		default:
			lowPriority = append(lowPriority, anime)
		}
	}

	color.Cyan("High priority queue (%d):", len(highPriority))
	refreshQueue(highPriority)

	color.Cyan("Medium priority queue (%d):", len(mediumPriority))
	refreshQueue(mediumPriority)

	color.Cyan("Low priority queue (%d):", len(lowPriority))
	refreshQueue(lowPriority)

	// This is a lazy hack: Wait 5 minutes for goroutines to finish their remaining work.
	time.Sleep(5 * time.Minute)
}

func refreshQueue(queue []*arn.Anime) {
	count := 0

	for _, anime := range queue {
		if sync(anime) {
			anime.Save()
			count++
		}
	}

	color.Green("Added Shoboi IDs for %d anime", count)
}

func sync(anime *arn.Anime) bool {
	// If we already have the ID, nothing to do here
	if anime.GetMapping("shoboi/anime") != "" {
		return false
	}

	// Log ID and title
	print(anime.ID + " | [JP] " + anime.Title.Japanese + " | [EN] " + anime.Title.Canonical)

	// Search Japanese title
	if anime.GetMapping("shoboi/anime") == "" && anime.Title.Japanese != "" {
		search(anime, anime.Title.Japanese)
	}

	// Search English title
	if anime.GetMapping("shoboi/anime") == "" && anime.Title.English != "" {
		search(anime, anime.Title.English)
	}

	// Did we get the ID?
	if anime.GetMapping("shoboi/anime") != "" {
		println(color.GreenString("✔"))
		return true
	}

	println(color.RedString("✘"))
	return false
}

// Search for a specific title
func search(anime *arn.Anime, title string) {
	shoboi, err := shoboi.SearchAnime(title)

	if err != nil {
		color.Red(err.Error())
		return
	}

	if shoboi == nil {
		return
	}

	// Copy titles
	if shoboi.TitleJapanese != "" {
		anime.Title.Japanese = shoboi.TitleJapanese
	}

	if shoboi.TitleHiragana != "" {
		anime.Title.Hiragana = shoboi.TitleHiragana
	}

	if shoboi.FirstChannel != "" {
		anime.FirstChannel = shoboi.FirstChannel
	}

	// This will start a goroutine that saves the anime
	anime.SetMapping("shoboi/anime", shoboi.TID)
}
