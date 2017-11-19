package main

import (
	"os"
	"path"

	"github.com/animenotifier/arn"
)

func main() {
	defer arn.Node.Close()

	for anime := range arn.StreamAnime() {
		base := path.Join(os.Getenv("GOPATH"), "/src/github.com/animenotifier/notify.moe/images/anime/original/", anime.ID)

		if _, err := os.Stat(base + ".png"); err == nil {
			anime.ImageExtension = ".png"
			anime.Save()
			continue
		}

		if _, err := os.Stat(base + ".jpg"); err == nil {
			anime.ImageExtension = ".jpg"
			anime.Save()
			continue
		}

		if _, err := os.Stat(base + ".jpeg"); err == nil {
			anime.ImageExtension = ".jpg"
			anime.Save()
			continue
		}

		if _, err := os.Stat(base + ".gif"); err == nil {
			anime.ImageExtension = ".gif"
			anime.Save()
			continue
		}

		if _, err := os.Stat(base + ".webp"); err == nil {
			anime.ImageExtension = ".webp"
			anime.Save()
			continue
		}
	}
}
