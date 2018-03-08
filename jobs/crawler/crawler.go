package main

import (
	"fmt"
	"time"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"

	"github.com/aerogo/crawler"
)

const userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.166 Safari/537.36"

func main() {
	defer arn.Node.Close()

	malCrawler := crawler.New(userAgent, 1*time.Second, 20000)
	count := 0

	for anime := range arn.StreamAnime() {
		malID := anime.GetMapping("myanimelist/anime")

		if malID == "" {
			continue
		}

		malCrawler.Queue(&crawler.Task{
			URL:         "https://myanimelist.net/anime/" + malID,
			Destination: fmt.Sprintf("mal/anime-%s.html", malID),
		})

		count++
	}

	color.Yellow("Queued up %d links", count)
	malCrawler.Wait()
}
