package main

import (
	"os"
	"path"

	"github.com/animenotifier/arn"
)

func main() {
	defer arn.Node.Close()

	for anime := range arn.StreamAnime() {
		base := path.Join(arn.Root, "/images/anime/original/", anime.ID)

		if _, err := os.Stat(base + ".png"); err == nil {
			anime.Image.Extension = ".png"
			anime.Save()
			continue
		}

		if _, err := os.Stat(base + ".jpg"); err == nil {
			anime.Image.Extension = ".jpg"
			anime.Save()
			continue
		}

		if _, err := os.Stat(base + ".jpeg"); err == nil {
			anime.Image.Extension = ".jpg"
			anime.Save()
			continue
		}

		if _, err := os.Stat(base + ".gif"); err == nil {
			anime.Image.Extension = ".gif"
			anime.Save()
			continue
		}

		if _, err := os.Stat(base + ".webp"); err == nil {
			anime.Image.Extension = ".webp"
			anime.Save()
			continue
		}
	}
}
