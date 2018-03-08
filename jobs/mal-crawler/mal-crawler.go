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
	maxAge               = 30 * 24 * time.Hour
	delayBetweenRequests = 1000 * time.Millisecond
	userAgent            = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.166 Safari/537.36"
)

func main() {
	defer arn.Node.Close()

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

	// Create crawler
	malCrawler := crawler.New(
		map[string]string{
			"User-Agent":      userAgent,
			"Accept-Encoding": "gzip",
		},
		delayBetweenRequests,
		len(animes),
	)

	// Sort so that we download the most important ones first
	arn.SortAnimeByQuality(animes, "")

	// Queue up URLs
	count := 0

	for _, anime := range animes {
		malID := anime.GetMapping("myanimelist/anime")
		url := "https://myanimelist.net/anime/" + malID
		filePath := fmt.Sprintf("files/anime-%s.html", malID)
		fileInfo, err := os.Stat(filePath)

		if err == nil && time.Since(fileInfo.ModTime()) <= maxAge {
			// fmt.Println(color.YellowString(url), "skip")
			continue
		}

		malCrawler.Queue(&crawler.Task{
			URL:         url,
			Destination: filePath,
		})

		count++
	}

	color.Yellow("Queued up %d links", count)
	malCrawler.Wait()
}
