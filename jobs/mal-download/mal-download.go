package main

import (
	"fmt"
	"os"
	"time"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"

	"github.com/aerogo/crawler"
)

const (
	// The maximum age of files we accept until we force a refresh.
	maxAge               = 24 * time.Hour
	delayBetweenRequests = 1100 * time.Millisecond
	userAgent            = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.20 Safari/537.36"
	animeDirectory       = "anime"
	characterDirectory   = "character"
)

var headers = map[string]string{
	"User-Agent":      userAgent,
	"Accept-Encoding": "gzip",
}

func main() {
	defer color.Green("Finished.")

	// Called with arguments?
	if InvokeShellArgs() {
		return
	}

	// Filter anime with MAL ID
	var animes []*arn.Anime

	if objectType == "all" || objectType == "anime" {
		animes = arn.FilterAnime(func(anime *arn.Anime) bool {
			return anime.GetMapping("myanimelist/anime") != ""
		})

		color.Yellow("Found %d anime", len(animes))

		// Sort so that we download the most important ones first
		arn.SortAnimeByQuality(animes)

		// Create anime directory if it's missing
		os.Mkdir(animeDirectory, 0777)
	}

	// Filter characters with MAL ID
	var characters []*arn.Character

	if objectType == "all" || objectType == "character" {
		characters = arn.FilterCharacters(func(character *arn.Character) bool {
			return character.GetMapping("myanimelist/character") != ""
		})

		color.Yellow("Found %d characters", len(characters))

		// Sort so that we download the most important ones first
		arn.SortCharactersByLikes(characters)

		// Create character directory if it's missing
		os.Mkdir(characterDirectory, 0777)
	}

	// We don't need the database anymore
	arn.Node.Close()

	// Create crawler
	malCrawler := crawler.New(
		headers,
		delayBetweenRequests,
		len(animes)+len(characters),
	)

	// Queue up URLs
	count := 0

	for _, anime := range animes {
		queueAnime(anime, malCrawler)
		count++
	}

	for _, character := range characters {
		queueCharacter(character, malCrawler)
		count++
	}

	// Log number of links
	color.Yellow("Queued up %d links", count)

	// Wait for completion
	malCrawler.Wait()
}

func queueAnime(anime *arn.Anime, malCrawler *crawler.Crawler) {
	malID := anime.GetMapping("myanimelist/anime")
	url := "https://myanimelist.net/anime/" + malID
	filePath := fmt.Sprintf("%s/%s.html.gz", animeDirectory, malID)
	fileInfo, err := os.Stat(filePath)

	if err == nil && time.Since(fileInfo.ModTime()) <= maxAge {
		// fmt.Println(color.YellowString(url), "skip")
		return
	}

	malCrawler.Queue(&crawler.Task{
		URL:         url,
		Destination: filePath,
		Raw:         true,
	})
}

func queueCharacter(character *arn.Character, malCrawler *crawler.Crawler) {
	malID := character.GetMapping("myanimelist/character")
	url := "https://myanimelist.net/character/" + malID
	filePath := fmt.Sprintf("%s/%s.html.gz", characterDirectory, malID)
	fileInfo, err := os.Stat(filePath)

	if err == nil && time.Since(fileInfo.ModTime()) <= maxAge {
		// fmt.Println(color.YellowString(url), "skip")
		return
	}

	malCrawler.Queue(&crawler.Task{
		URL:         url,
		Destination: filePath,
		Raw:         true,
	})
}
