package main

func main() {}

// import (
// 	"bytes"
// 	"fmt"
// 	"image"
// 	"image/color"
// 	_ "image/jpeg"
// 	"image/png"
// 	"io/ioutil"
// 	"os"
// )

// func main() {
// 	data, err := ioutil.ReadFile("input.jpg")

// 	if err != nil {
// 		panic(err)
// 	}

// 	img, _, err := image.Decode(bytes.NewReader(data))

// 	if err != nil {
// 		panic(err)
// 	}

// 	improved := ImproveQuality(img)
// 	f, err := os.Create("output.png")

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer f.Close()
// 	png.Encode(f, improved)
// }

// // Pixel ...
// type Pixel struct {
// 	X int
// 	Y int
// }

// // Area ...
// type Area struct {
// 	Pixels []Pixel
// 	totalR uint64
// 	totalG uint64
// 	totalB uint64
// 	totalA uint64
// }

// // Add ...
// func (area *Area) Add(x, y int, r, g, b, a uint32) {
// 	area.Pixels = append(area.Pixels, Pixel{
// 		X: x,
// 		Y: y,
// 	})

// 	area.totalR += uint64(r)
// 	area.totalG += uint64(g)
// 	area.totalB += uint64(b)
// 	area.totalA += uint64(a)
// }

// // AverageColor ...
// func (area *Area) AverageColor() color.Color {
// 	if len(area.Pixels) == 0 {
// 		return color.Transparent
// 	}

// 	return color.RGBA64{
// 		R: uint16(area.totalR / uint64(len(area.Pixels))),
// 		G: uint16(area.totalG / uint64(len(area.Pixels))),
// 		B: uint16(area.totalB / uint64(len(area.Pixels))),
// 		A: uint16(area.totalA / uint64(len(area.Pixels))),
// 	}
// }

// const (
// 	tolerance = uint32(3000)
// )

// func diffAbs(a uint32, b uint32) uint32 {
// 	if a > b {
// 		return a - b
// 	}

// 	return b - a
// }

// // ImproveQuality returns the average color of an image in HSL format.
// func ImproveQuality(img image.Image) *image.NRGBA {
// 	width := img.Bounds().Dx()
// 	height := img.Bounds().Dy()
// 	clone := image.NewNRGBA(image.Rect(0, 0, width, height))
// 	areas := []*Area{}
// 	areaIndexMap := make([]int, width*height)

// 	for x := 0; x < width; x++ {
// 		for y := 0; y < height; y++ {
// 			color := img.At(x, y)
// 			r, g, b, a := color.RGBA()
// 			areaIndex := -1

// 			// Find similar area
// 			for i := 0; i < len(areas); i++ {
// 				area := areas[i]
// 				avgR, avgG, avgB, _ := area.AverageColor().RGBA()

// 				// Is the color similar?
// 				if diffAbs(r, avgR) <= tolerance && diffAbs(g, avgG) <= tolerance && diffAbs(b, avgB) <= tolerance {
// 					areaIndex = i
// 					break
// 				}
// 			}

// 			// Insert new area
// 			if areaIndex == -1 {
// 				areaIndex = len(areas)
// 				areas = append(areas, &Area{})
// 			}

// 			areaIndexMap[y*width+x] = areaIndex
// 			areas[areaIndex].Add(x, y, r, g, b, a)
// 		}
// 	}

// 	fmt.Println(len(areas), "areas")

// 	// Reduce noise
// 	noiseCount := 0

// 	for areaIndex, area := range areas {
// 		noisePixelIndices := []int{}
// 		areaSurroundedBy := map[int]int{}

// 		for i := 0; i < len(area.Pixels); i++ {
// 			// If pixel is surrounded by 4 different areas, remove it
// 			pixel := area.Pixels[i]
// 			x := pixel.X
// 			y := pixel.Y
// 			left := areaIndex
// 			right := areaIndex
// 			top := areaIndex
// 			bottom := areaIndex

// 			if x > 0 {
// 				left = areaIndexMap[y*width+(x-1)]
// 			}

// 			if x < width-1 {
// 				right = areaIndexMap[y*width+(x+1)]
// 			}

// 			if y > 0 {
// 				top = areaIndexMap[(y-1)*width+x]
// 			}

// 			if y < height-1 {
// 				bottom = areaIndexMap[(y+1)*width+x]
// 			}

// 			differentNeighbors := 0

// 			if left != areaIndex {
// 				differentNeighbors++
// 			}

// 			if right != areaIndex {
// 				differentNeighbors++
// 			}

// 			if top != areaIndex {
// 				differentNeighbors++
// 			}

// 			if bottom != areaIndex {
// 				differentNeighbors++
// 			}

// 			// Determine surrounding area
// 			areaIndexScore := map[int]int{}

// 			areaIndexScore[left]++
// 			areaIndexScore[right]++
// 			areaIndexScore[top]++
// 			areaIndexScore[bottom]++

// 			areaSurroundedBy[left]++
// 			areaSurroundedBy[right]++
// 			areaSurroundedBy[top]++
// 			areaSurroundedBy[bottom]++

// 			newAreaIndex := -1
// 			bestScore := 0

// 			for checkIndex, score := range areaIndexScore {
// 				if score > bestScore {
// 					bestScore = score
// 					newAreaIndex = checkIndex
// 				}
// 			}

// 			if differentNeighbors >= 3 && bestScore >= 3 {
// 				noiseCount++
// 				noisePixelIndices = append(noisePixelIndices, i)

// 				// Add to surrounding area
// 				r, g, b, a := img.At(x, y).RGBA()
// 				areas[newAreaIndex].Add(x, y, r, g, b, a)

// 				area.totalR -= uint64(r)
// 				area.totalG -= uint64(g)
// 				area.totalB -= uint64(b)
// 				area.totalA -= uint64(a)
// 			}
// 		}

// 		// Remove noise pixels
// 		offset := 0

// 		for _, removal := range noisePixelIndices {
// 			index := removal - offset
// 			area.Pixels = append(area.Pixels[:index], area.Pixels[index+1:]...)
// 			offset++
// 		}

// 		// // Determine surrounding area
// 		// surroundingAreaIndex := -1
// 		// bestScore := 0

// 		// for checkIndex, score := range areaSurroundedBy {
// 		// 	if score > bestScore && checkIndex != areaIndex {
// 		// 		bestScore = score
// 		// 		surroundingAreaIndex = checkIndex
// 		// 	}
// 		// }

// 		// surroundingArea := areas[surroundingAreaIndex]

// 		// if areaIndex != surroundingAreaIndex && len(surroundingArea.Pixels) > len(area.Pixels)*2 {
// 		// 	// const surroundTolerance = 5000

// 		// 	// r1, g1, b1, a1 := area.AverageColor().RGBA()
// 		// 	// r2, g2, b2, a2 := surroundingArea.AverageColor().RGBA()

// 		// 	// if diffAbs(r1, r2) < surroundTolerance && diffAbs(g1, g2) < surroundTolerance && diffAbs(b1, b2) < surroundTolerance && diffAbs(a1, a2) < surroundTolerance {
// 		// 	// 	// fmt.Println(areaIndex, "surrounded by", surroundingAreaIndex, "|", len(area.Pixels), len(surroundingArea.Pixels))

// 		// 	// 	// Add pixels to surrounding area
// 		// 	// 	for _, pixel := range area.Pixels {
// 		// 	// 		r, g, b, a := img.At(pixel.X, pixel.Y).RGBA()
// 		// 	// 		surroundingArea.Add(pixel.X, pixel.Y, r, g, b, a)
// 		// 	// 	}

// 		// 	// 	// Remove this area
// 		// 	// 	area.Pixels = nil
// 		// 	// 	area.totalR = 0
// 		// 	// 	area.totalG = 0
// 		// 	// 	area.totalB = 0
// 		// 	// 	area.totalA = 0
// 		// 	// }
// 		// }
// 	}

// 	fmt.Println(noiseCount, "noise pixels")

// 	pixelCount := 0

// 	for _, area := range areas {
// 		pixelCount += len(area.Pixels)
// 	}

// 	fmt.Println(pixelCount, "pixels", width*height)

// 	// Build image from areas
// 	for _, area := range areas {
// 		avgColor := area.AverageColor()

// 		for _, pixel := range area.Pixels {
// 			clone.Set(pixel.X, pixel.Y, avgColor)
// 		}
// 	}

// 	return clone
// }
