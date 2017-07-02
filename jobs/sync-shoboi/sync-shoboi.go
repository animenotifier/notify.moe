package main

import (
	"time"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/shoboi"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Syncing Shoboi Anime")

	// Get a slice of all anime
	allAnime, _ := arn.AllAnime()

	// Iterate over the slice
	count := 0
	for _, anime := range allAnime {
		if sync(anime) {
			count++
		}
	}

	// Log
	color.Green("Successfully added Shoboi IDs for %d anime", count)

	// This is a lazy hack: Wait 5 minutes for goroutines to finish their remaining work.
	time.Sleep(5 * time.Minute)

	color.Green("Finished.")
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
		time.Sleep(2 * time.Second)
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
	anime.AddMapping("shoboi/anime", shoboi.TID, "")
}
