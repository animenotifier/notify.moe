package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Updating anime thumbnails")

	defer arn.Node.Close()
	defer color.Green("Finished.")

	for anime := range arn.StreamAnime() {
		base := path.Join(arn.Root, "/images/anime/original/", anime.ID)

		if _, err := os.Stat(base + ".png"); err == nil {
			update(anime, base+".png")
			continue
		}

		if _, err := os.Stat(base + ".jpg"); err == nil {
			update(anime, base+".jpg")
			continue
		}

		if _, err := os.Stat(base + ".jpeg"); err == nil {
			update(anime, base+".jpg")
			continue
		}

		if _, err := os.Stat(base + ".gif"); err == nil {
			update(anime, base+".gif")
			continue
		}

		if _, err := os.Stat(base + ".webp"); err == nil {
			update(anime, base+".webp")
			continue
		}
	}
}

func update(anime *arn.Anime, filePath string) {
	fmt.Println(anime.ID, anime)

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		return
	}

	anime.SetImageBytes(data)
	anime.Save()
}
