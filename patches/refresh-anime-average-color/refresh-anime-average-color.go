package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"path"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Updating anime image average colors")

	defer arn.Node.Close()
	defer color.Green("Finished.")

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
	base := path.Join(arn.Root, "/images/anime/small/", anime.ID)

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

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	totalR := uint64(0)
	totalG := uint64(0)
	totalB := uint64(0)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			totalR += uint64(r)
			totalG += uint64(g)
			totalB += uint64(b)
		}
	}

	pixels := uint64(width * height)

	const max = float64(65535)
	averageR := float64(totalR/pixels) / max
	averageG := float64(totalG/pixels) / max
	averageB := float64(totalB/pixels) / max

	h, s, l := arn.RGBToHSL(averageR, averageG, averageB)

	anime.Image.AverageColor.Hue = h
	anime.Image.AverageColor.Saturation = s
	anime.Image.AverageColor.Lightness = l

	anime.Save()
}
