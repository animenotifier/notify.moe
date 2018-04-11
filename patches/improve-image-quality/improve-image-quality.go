package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"math"
	"os"

	"github.com/animenotifier/arn"
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
	X     int
	Y     int
	Color arn.HSLColor
}

// Area ...
type Area struct {
	AverageColor color.Color
	Pixels       []Pixel
}

// Add ...
func (area *Area) Add(x, y int, hsl arn.HSLColor) {
	area.Pixels = append(area.Pixels, Pixel{
		X:     x,
		Y:     y,
		Color: hsl,
	})
}

const (
	hueTolerance        = 0.1
	lightnessTolerance  = 0.1
	saturationTolerance = 0.1
)

// ImproveQuality returns the average color of an image in HSL format.
func ImproveQuality(img image.Image) *image.NRGBA {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	clone := image.NewNRGBA(image.Rect(0, 0, width, height))
	hueAreas := []Area{}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			color := img.At(x, y)
			rUint, gUint, bUint, _ := color.RGBA()
			r := float64(rUint) / max
			g := float64(gUint) / max
			b := float64(bUint) / max
			h, s, l := arn.RGBToHSL(r, g, b)
			areaIndex := -1

			// Find similar area
			for i := 0; i < len(hueAreas); i++ {
				area := hueAreas[i]

				// Is the pixel close to any pixel in the area we're checking?
				for _, pixel := range area.Pixels {
					xDist := x - pixel.X
					yDist := y - pixel.Y

					if xDist < 0 {
						xDist = -xDist
					}

					if yDist < 0 {
						yDist = -yDist
					}

					if xDist <= 1 && yDist <= 1 {
						// Is the color similar?
						if math.Abs(h-pixel.Color.Hue) <= hueTolerance && math.Abs(s-pixel.Color.Saturation) <= saturationTolerance && math.Abs(l-pixel.Color.Lightness) <= lightnessTolerance {
							areaIndex = i
							break
						}
					}
				}

				if areaIndex != -1 {
					break
				}
			}

			// Insert new area
			if areaIndex == -1 {
				areaIndex = len(hueAreas)
				hueAreas = append(hueAreas, Area{})
			}

			hueAreas[areaIndex].Add(x, y, arn.HSLColor{
				Hue:        h,
				Saturation: s,
				Lightness:  l,
			})
		}
	}

	fmt.Println(len(hueAreas), "areas")

	// Build image from areas
	for _, area := range hueAreas {
		totalR := uint64(0)
		totalG := uint64(0)
		totalB := uint64(0)

		// Calculate area average color
		for _, pixel := range area.Pixels {
			col := img.At(pixel.X, pixel.Y)
			r, g, b, _ := col.RGBA()
			totalR += uint64(r)
			totalG += uint64(g)
			totalB += uint64(b)
		}

		averageR := float64(totalR/uint64(len(area.Pixels))) / max
		averageG := float64(totalG/uint64(len(area.Pixels))) / max
		averageB := float64(totalB/uint64(len(area.Pixels))) / max

		area.AverageColor = color.RGBA{
			R: uint8(averageR * 255),
			G: uint8(averageG * 255),
			B: uint8(averageB * 255),
			A: 255,
		}

		for _, pixel := range area.Pixels {
			clone.Set(pixel.X, pixel.Y, area.AverageColor)
		}
	}

	return clone
}
