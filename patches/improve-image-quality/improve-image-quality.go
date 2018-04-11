package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("input.jpg")

	if err != nil {
		panic(err)
	}

	img, format, err := image.Decode(bytes.NewReader(data))

	if err != nil {
		panic(err)
	}

	fmt.Println(img.Bounds().Dx(), img.Bounds().Dy(), format)
	improved := ImproveQuality(img)
	f, err := os.Create("output.png")

	if err != nil {
		panic(err)
	}

	defer f.Close()
	png.Encode(f, improved)
}

const max = float64(65535)

// Pixel ...
type Pixel struct {
	X int
	Y int
}

// Area ...
type Area struct {
	color.Color
	Pixels []Pixel
	totalR uint64
	totalG uint64
	totalB uint64
	totalA uint64
}

// Add ...
func (area *Area) Add(x, y int, r, g, b, a uint32) {
	area.Pixels = append(area.Pixels, Pixel{
		X: x,
		Y: y,
	})

	area.totalR += uint64(r)
	area.totalG += uint64(g)
	area.totalB += uint64(b)
	area.totalA += uint64(a)
}

// AverageColor ...
func (area *Area) AverageColor() color.Color {
	return color.RGBA64{
		R: uint16(area.totalR / uint64(len(area.Pixels))),
		G: uint16(area.totalG / uint64(len(area.Pixels))),
		B: uint16(area.totalB / uint64(len(area.Pixels))),
		A: uint16(area.totalA / uint64(len(area.Pixels))),
	}
}

const (
	tolerance           = uint32(3000)
	hueTolerance        = 0.1
	lightnessTolerance  = 0.1
	saturationTolerance = 0.1
)

func diffAbs(a uint32, b uint32) uint32 {
	if a > b {
		return a - b
	}

	return b - a
}

// ImproveQuality returns the average color of an image in HSL format.
func ImproveQuality(img image.Image) *image.NRGBA {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	clone := image.NewNRGBA(image.Rect(0, 0, width, height))
	hueAreas := []Area{}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			color := img.At(x, y)
			r, g, b, a := color.RGBA()
			areaIndex := -1

			// Find similar area
			for i := 0; i < len(hueAreas); i++ {
				area := hueAreas[i]
				avgR, avgG, avgB, _ := area.AverageColor().RGBA()

				// Is the color similar?
				if diffAbs(r, avgR) <= tolerance && diffAbs(g, avgG) <= tolerance && diffAbs(b, avgB) <= tolerance {
					areaIndex = i
					break
				}
			}

			// Insert new area
			if areaIndex == -1 {
				areaIndex = len(hueAreas)
				hueAreas = append(hueAreas, Area{})
			}

			hueAreas[areaIndex].Add(x, y, r, g, b, a)
		}
	}

	fmt.Println(len(hueAreas), "areas")

	// Build image from areas
	for _, area := range hueAreas {
		avgColor := area.AverageColor()

		for _, pixel := range area.Pixels {
			clone.Set(pixel.X, pixel.Y, avgColor)
		}
	}

	return clone
}
