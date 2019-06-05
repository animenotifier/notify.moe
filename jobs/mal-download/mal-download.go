package main

import (
	"fmt"
	"os"
	"time"

	"github.com/animenotifier/notify.moe/arn/osutils"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"

	"github.com/aerogo/crawler"
)

const (
	// The maximum age of files we accept until we force a refresh.
	maxAge               = 7 * 24 * time.Hour
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

	// Create directories in case they're missing
	err := os.Mkdir(animeDirectory, 0777)

	if err != nil {
		panic(err)
	}

	err = os.Mkdir(characterDirectory, 0777)

	if err != nil {
		panic(err)
	}

	// Called with arguments?
	if InvokeShellArgs() {
		return
	}

	// Filter anime with MAL ID
	var animes []*arn.Anime

	if objectType == "all" || objectType == "anime" {
		animes = arn.FilterAnime(func(anime *arn.Anime) bool {
			malID := anime.GetMapping("myanimelist/anime")

			if malID == "" {
				return false
			}

			return !newOnly || !osutils.Exists(animeFilePath(malID))
		})

		color.Yellow("Found %d anime", len(animes))

		// Sort so that we download the most important ones first
		arn.SortAnimeByQuality(animes)
	}

	// Filter characters with MAL ID
	var characters []*arn.Character

	if objectType == "all" || objectType == "character" {
		characters = arn.FilterCharacters(func(character *arn.Character) bool {
			malID := character.GetMapping("myanimelist/character")

			if malID == "" {
				return false
			}

			return !newOnly || !osutils.Exists(characterFilePath(malID))
		})

		color.Yellow("Found %d characters", len(characters))

		// Sort so that we download the most important ones first
		arn.SortCharactersByLikes(characters)
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

func animeFilePath(malID string) string {
	return fmt.Sprintf("%s/%s.html.gz", animeDirectory, malID)
}

func characterFilePath(malID string) string {
	return fmt.Sprintf("%s/%s.html.gz", characterDirectory, malID)
}

func queueAnime(anime *arn.Anime, malCrawler *crawler.Crawler) {
	malID := anime.GetMapping("myanimelist/anime")
	url := "https://myanimelist.net/anime/" + malID
	filePath := animeFilePath(malID)
	fileInfo, err := os.Stat(filePath)

	if err == nil && time.Since(fileInfo.ModTime()) <= maxAge {
		// fmt.Println(color.YellowString(url), "skip")
		return
	}

	err = malCrawler.Queue(&crawler.Task{
		URL:         url,
		Destination: filePath,
		Raw:         true,
	})

	if err != nil {
		panic(err)
	}
}

func queueCharacter(character *arn.Character, malCrawler *crawler.Crawler) {
	malID := character.GetMapping("myanimelist/character")
	url := "https://myanimelist.net/character/" + malID
	filePath := characterFilePath(malID)
	fileInfo, err := os.Stat(filePath)

	if err == nil && time.Since(fileInfo.ModTime()) <= maxAge {
		// fmt.Println(color.YellowString(url), "skip")
		return
	}

	err = malCrawler.Queue(&crawler.Task{
		URL:         url,
		Destination: filePath,
		Raw:         true,
	})

	if err != nil {
		panic(err)
	}
}
