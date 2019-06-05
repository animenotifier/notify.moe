package main

import (
	"bytes"
	"fmt"
	"image"
	"os"
	"path"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"

	icolor "image/color"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
)

func main() {
	defer color.Green("Finished.")
	defer arn.Node.Close()

	characters := arn.FilterCharacters(func(character *arn.Character) bool {
		return character.HasImage()
	})

	for index, character := range characters {
		fmt.Printf("[%d / %d] %s\n", index+1, len(characters), character)
		process(character)
	}

	// char, _ := arn.GetCharacter("EI3HwrmiRm")
	// process(char)
}

func process(character *arn.Character) {
	file, err := os.Open(path.Join(arn.Root, "images", "characters", "original", character.ID+character.Image.Extension))

	if err != nil {
		color.Red(err.Error())
		return
	}

	defer file.Close()

	img, format, err := image.Decode(file)

	if err != nil {
		color.Red(err.Error())
		return
	}

	newImg := removeBorders(img)

	fmt.Println(img.Bounds().Dx(), "->", newImg.Bounds().Dx(), "width", format)

	buffer := bytes.NewBuffer(nil)
	err = png.Encode(buffer, newImg)

	if err != nil {
		color.Red(err.Error())
		return
	}

	err = character.SetImageBytes(buffer.Bytes())

	if err != nil {
		color.Red(err.Error())
		return
	}

	character.Save()
}

func diffAbs(a uint32, b uint32) uint32 {
	if a > b {
		return a - b
	}

	return b - a
}

func removeBorders(img image.Image) image.Image {
	const maxBorderLength = 3

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	borderLength := 0

	for ; borderLength <= maxBorderLength; borderLength++ {
		edgeColors := []icolor.Color{}

		// Top edge
		for x := borderLength; x < width-borderLength*2; x++ {
			c := img.At(x, borderLength)
			edgeColors = append(edgeColors, c)
		}

		// Bottom edge
		for x := borderLength; x < width-borderLength*2; x++ {
			c := img.At(x, height-borderLength-1)
			edgeColors = append(edgeColors, c)
		}

		// Left edge
		for y := borderLength; y < height-borderLength*2; y++ {
			c := img.At(borderLength, y)
			edgeColors = append(edgeColors, c)
		}

		// Right edge
		for y := borderLength; y < height-borderLength*2; y++ {
			c := img.At(width-borderLength-1, y)
			edgeColors = append(edgeColors, c)
		}

		// Check if all edge colors are similar.
		// Find average color first.
		totalR := uint64(0)
		totalG := uint64(0)
		totalB := uint64(0)

		for _, c := range edgeColors {
			r, g, b, _ := c.RGBA()
			totalR += uint64(r)
			totalG += uint64(g)
			totalB += uint64(b)
		}

		averageR := uint32(totalR / uint64(len(edgeColors)))
		averageG := uint32(totalG / uint64(len(edgeColors)))
		averageB := uint32(totalB / uint64(len(edgeColors)))

		const tolerance = 12000

		// Check if the colors are close to the average color
		notSimilarCount := 0

		for _, c := range edgeColors {
			r, g, b, _ := c.RGBA()

			if diffAbs(r, averageR) > tolerance || diffAbs(g, averageG) > tolerance || diffAbs(b, averageB) > tolerance {
				notSimilarCount++
			}
		}

		if notSimilarCount >= 5 {
			break
		}
	}

	newWidth := width - borderLength*2
	newHeight := height - borderLength*2
	newImg := image.NewNRGBA(image.Rect(0, 0, newWidth, newHeight))

	for x := 0; x < newWidth; x++ {
		for y := 0; y < newHeight; y++ {
			newImg.Set(x, y, img.At(borderLength+x, borderLength+y))
		}
	}

	return newImg
}
