package main

import (
	"time"

	"github.com/aerogo/crawler"
)

const userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.166 Safari/537.36"

func main() {
	malCrawler := crawler.New(userAgent, 1*time.Second, 20000)

	malCrawler.Queue(&crawler.Task{
		URL:         "https://github.com/animenotifier/notify.moe",
		Destination: "file.html",
	})

	malCrawler.Wait()
}
