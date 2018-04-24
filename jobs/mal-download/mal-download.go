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
	userAgent            = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36"
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
	animes := []*arn.Anime{}

	for anime := range arn.StreamAnime() {
		malID := anime.GetMapping("myanimelist/anime")

		if malID == "" {
			continue
		}

		animes = append(animes, anime)
	}

	color.Yellow("Found %d anime", len(animes))

	// We don't need the database anymore
	arn.Node.Close()

	// Create files directory if it's missing
	os.Mkdir("files", 0777)

	// Create crawler
	malCrawler := crawler.New(
		headers,
		delayBetweenRequests,
		len(animes),
	)

	// Sort so that we download the most important ones first
	arn.SortAnimeByQuality(animes)

	// Queue up URLs
	count := 0

	for _, anime := range animes {
		queue(anime, malCrawler)
		count++
	}

	// Log number of links
	color.Yellow("Queued up %d links", count)

	// Wait for completion
	malCrawler.Wait()
}

func queue(anime *arn.Anime, malCrawler *crawler.Crawler) {
	malID := anime.GetMapping("myanimelist/anime")
	url := "https://myanimelist.net/anime/" + malID
	filePath := fmt.Sprintf("anime/anime-%s.html", malID)
	fileInfo, err := os.Stat(filePath)

	if err == nil && time.Since(fileInfo.ModTime()) <= maxAge {
		// fmt.Println(color.YellowString(url), "skip")
		return
	}

	malCrawler.Queue(&crawler.Task{
		URL:         url,
		Destination: filePath,
	})
}
