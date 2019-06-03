package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"path"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Updating anime image average colors")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	// Parse flags
	var animeID string
	flag.StringVar(&animeID, "id", "", "ID of the anime that you want to refresh")
	flag.Parse()

	// Refresh 1 anime in case ID was specified
	if animeID != "" {
		anime, _ := arn.GetAnime(animeID)

		if anime != nil {
			work(anime)
		}

		return
	}

	// Otherwise refresh all anime
	for anime := range arn.StreamAnime() {
		work(anime)
	}
}

// work refreshes the average color of the given anime.
func work(anime *arn.Anime) {
	base := path.Join(arn.Root, "/images/anime/medium/", anime.ID)

	if _, err := os.Stat(base + ".jpg"); err != nil {
		color.Red(err.Error())
		return
	}

	update(anime, base+".jpg")
}

// update expects a file to load as image for the anime and update the average color.
func update(anime *arn.Anime, filePath string) {
	fmt.Println(anime.ID, anime)

	// Load
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		color.Red(err.Error())
		return
	}

	// Decode
	img, _, err := image.Decode(bytes.NewReader(data))

	if err != nil {
		color.Red(err.Error())
		return
	}

	anime.Image.AverageColor = arn.GetAverageColor(img)
	anime.Save()
}
