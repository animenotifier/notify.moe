package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Updating anime thumbnails")

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
			sync(anime)
		}

		return
	}

	// Otherwise refresh all anime
	for anime := range arn.StreamAnime() {
		sync(anime)
	}
}

// sync refreshes the image of the given anime.
func sync(anime *arn.Anime) {
	base := path.Join(arn.Root, "/images/anime/original/", anime.ID)

	if _, err := os.Stat(base + ".png"); err == nil {
		update(anime, base+".png")
		return
	}

	if _, err := os.Stat(base + ".jpg"); err == nil {
		update(anime, base+".jpg")
		return
	}

	if _, err := os.Stat(base + ".jpeg"); err == nil {
		update(anime, base+".jpg")
		return
	}

	if _, err := os.Stat(base + ".gif"); err == nil {
		update(anime, base+".gif")
		return
	}

	if _, err := os.Stat(base + ".webp"); err == nil {
		update(anime, base+".webp")
		return
	}
}

// update expects a file to load as image for the anime and updates it.
func update(anime *arn.Anime, filePath string) {
	fmt.Println(anime.ID, anime)

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		color.Red(err.Error())
		return
	}

	err = anime.SetImageBytes(data)

	if err != nil {
		color.Red(err.Error())
		return
	}

	anime.Save()
}
