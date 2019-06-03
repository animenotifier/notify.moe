package arn

import (
	"fmt"
	"image"
	"math"
)

// HSLColor ...
type HSLColor struct {
	Hue        float64 `json:"hue"`
	Saturation float64 `json:"saturation"`
	Lightness  float64 `json:"lightness"`
}

// String returns a representation like hsl(0, 0%, 0%).
func (color HSLColor) String() string {
	return fmt.Sprintf("hsl(%.1f, %.1f%%, %.1f%%)", color.Hue*360, color.Saturation*100, color.Lightness*100)
}

// StringWithAlpha returns a representation like hsla(0, 0%, 0%, 0.5).
func (color HSLColor) StringWithAlpha(alpha float64) string {
	return fmt.Sprintf("hsla(%.1f, %.1f%%, %.1f%%, %.2f)", color.Hue*360, color.Saturation*100, color.Lightness*100, alpha)
}

// GetAverageColor returns the average color of an image in HSL format.
func GetAverageColor(img image.Image) HSLColor {
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

	h, s, l := RGBToHSL(averageR, averageG, averageB)
	return HSLColor{h, s, l}
}

// RGBToHSL converts RGB to HSL (RGB input and HSL output are floats in the 0..1 range).
// Original source: https://github.com/gerow/go-color
func RGBToHSL(r, g, b float64) (h, s, l float64) {
	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)

	// Luminosity is the average of the max and min rgb color intensities.
	l = (max + min) / 2

	// Saturation
	delta := max - min

	if delta == 0 {
		// It's gray
		return 0, 0, l
	}

	// It's not gray
	if l < 0.5 {
		s = delta / (max + min)
	} else {
		s = delta / (2 - max - min)
	}

	// Hue
	r2 := (((max - r) / 6) + (delta / 2)) / delta
	g2 := (((max - g) / 6) + (delta / 2)) / delta
	b2 := (((max - b) / 6) + (delta / 2)) / delta

	switch {
	case r == max:
		h = b2 - g2
	case g == max:
		h = (1.0 / 3.0) + r2 - b2
	case b == max:
		h = (2.0 / 3.0) + g2 - r2
	}

	// fix wraparounds
	switch {
	case h < 0:
		h++
	case h > 1:
		h--
	}

	return h, s, l
}
