package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/chai2010/webp"
)

func convertFileToWebP(in string, out string, quality float32) error {
	f, openErr := os.Open(in)

	if openErr != nil {
		return openErr
	}

	img, format, decodeErr := image.Decode(f)

	if decodeErr != nil {
		return decodeErr
	}

	fmt.Println(format, img.Bounds().Dx(), img.Bounds().Dy())

	fileOut, writeErr := os.Create(out)

	if writeErr != nil {
		return writeErr
	}

	encodeErr := webp.Encode(fileOut, img, &webp.Options{
		Quality: quality,
	})

	return encodeErr
}
