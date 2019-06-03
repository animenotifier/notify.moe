package arn

import (
	"bytes"
	"fmt"
	"image"
	"net/http"
	"path"
	"time"

	"github.com/aerogo/http/client"
	"github.com/akyoto/imageserver"
)

const (
	// CharacterImageLargeWidth is the minimum width in pixels of a large character image.
	// We subtract 6 pixels due to border removal which can remove up to 6 pixels.
	CharacterImageLargeWidth = 225 - 6

	// CharacterImageLargeHeight is the minimum height in pixels of a large character image.
	// We subtract 6 pixels due to border removal which can remove up to 6 pixels.
	CharacterImageLargeHeight = 350 - 6

	// CharacterImageMediumWidth is the minimum width in pixels of a medium character image.
	CharacterImageMediumWidth = 112

	// CharacterImageMediumHeight is the minimum height in pixels of a medium character image.
	CharacterImageMediumHeight = 112

	// CharacterImageSmallWidth is the minimum width in pixels of a small character image.
	CharacterImageSmallWidth = 56

	// CharacterImageSmallHeight is the minimum height in pixels of a small character image.
	CharacterImageSmallHeight = 56

	// CharacterImageWebPQuality is the WebP quality of character images.
	CharacterImageWebPQuality = 70

	// CharacterImageJPEGQuality is the JPEG quality of character images.
	CharacterImageJPEGQuality = 70

	// CharacterImageQualityBonusLowDPI ...
	CharacterImageQualityBonusLowDPI = 12

	// CharacterImageQualityBonusLarge ...
	CharacterImageQualityBonusLarge = 10

	// CharacterImageQualityBonusMedium ...
	CharacterImageQualityBonusMedium = 15

	// CharacterImageQualityBonusSmall ...
	CharacterImageQualityBonusSmall = 15
)

// Define the character image outputs
var characterImageOutputs = []imageserver.Output{
	// Original at full size
	&imageserver.OriginalFile{
		Directory: path.Join(Root, "images/characters/original/"),
		Width:     0,
		Height:    0,
		Quality:   0,
	},

	// JPEG - Large
	&imageserver.JPEGFile{
		Directory: path.Join(Root, "images/characters/large/"),
		Width:     CharacterImageLargeWidth,
		Height:    CharacterImageLargeHeight,
		Quality:   CharacterImageJPEGQuality + CharacterImageQualityBonusLowDPI + CharacterImageQualityBonusLarge,
	},

	// JPEG - Medium
	&imageserver.JPEGFile{
		Directory: path.Join(Root, "images/characters/medium/"),
		Width:     CharacterImageMediumWidth,
		Height:    CharacterImageMediumHeight,
		Quality:   CharacterImageJPEGQuality + CharacterImageQualityBonusLowDPI + CharacterImageQualityBonusMedium,
	},

	// JPEG - Small
	&imageserver.JPEGFile{
		Directory: path.Join(Root, "images/characters/small/"),
		Width:     CharacterImageSmallWidth,
		Height:    CharacterImageSmallHeight,
		Quality:   CharacterImageJPEGQuality + CharacterImageQualityBonusLowDPI + CharacterImageQualityBonusSmall,
	},

	// WebP - Large
	&imageserver.WebPFile{
		Directory: path.Join(Root, "images/characters/large/"),
		Width:     CharacterImageLargeWidth,
		Height:    CharacterImageLargeHeight,
		Quality:   CharacterImageWebPQuality + CharacterImageQualityBonusLowDPI + CharacterImageQualityBonusLarge,
	},

	// WebP - Medium
	&imageserver.WebPFile{
		Directory: path.Join(Root, "images/characters/medium/"),
		Width:     CharacterImageMediumWidth,
		Height:    CharacterImageMediumHeight,
		Quality:   CharacterImageWebPQuality + CharacterImageQualityBonusLowDPI + CharacterImageQualityBonusMedium,
	},

	// WebP - Small
	&imageserver.WebPFile{
		Directory: path.Join(Root, "images/characters/small/"),
		Width:     CharacterImageSmallWidth,
		Height:    CharacterImageSmallHeight,
		Quality:   CharacterImageWebPQuality + CharacterImageQualityBonusLowDPI + CharacterImageQualityBonusSmall,
	},
}

// Define the high DPI character image outputs
var characterImageOutputsHighDPI = []imageserver.Output{
	// NOTE: We don't save "large" images in double size because that's usually the maximum size anyway.

	// JPEG - Medium
	&imageserver.JPEGFile{
		Directory: path.Join(Root, "images/characters/medium/"),
		Width:     CharacterImageMediumWidth * 2,
		Height:    CharacterImageMediumHeight * 2,
		Quality:   CharacterImageJPEGQuality + CharacterImageQualityBonusMedium,
	},

	// JPEG - Small
	&imageserver.JPEGFile{
		Directory: path.Join(Root, "images/characters/small/"),
		Width:     CharacterImageSmallWidth * 2,
		Height:    CharacterImageSmallHeight * 2,
		Quality:   CharacterImageJPEGQuality + CharacterImageQualityBonusSmall,
	},

	// WebP - Medium
	&imageserver.WebPFile{
		Directory: path.Join(Root, "images/characters/medium/"),
		Width:     CharacterImageMediumWidth * 2,
		Height:    CharacterImageMediumHeight * 2,
		Quality:   CharacterImageWebPQuality + CharacterImageQualityBonusMedium,
	},

	// WebP - Small
	&imageserver.WebPFile{
		Directory: path.Join(Root, "images/characters/small/"),
		Width:     CharacterImageSmallWidth * 2,
		Height:    CharacterImageSmallHeight * 2,
		Quality:   CharacterImageWebPQuality + CharacterImageQualityBonusSmall,
	},
}

// CharacterImage ...
type CharacterImage AnimeImage

// SetImageBytes accepts a byte buffer that represents an image file and updates the character image.
func (character *Character) SetImageBytes(data []byte) error {
	// Decode
	img, format, err := image.Decode(bytes.NewReader(data))

	if err != nil {
		return err
	}

	return character.SetImage(&imageserver.MetaImage{
		Image:  img,
		Format: format,
		Data:   data,
	})
}

// SetImage sets the character image to the given MetaImage.
func (character *Character) SetImage(metaImage *imageserver.MetaImage) error {
	var lastError error

	// Save the different image formats and sizes in low DPI
	for _, output := range characterImageOutputs {
		err := output.Save(metaImage, character.ID)

		if err != nil {
			lastError = err
		}
	}

	// Save the different image formats and sizes in high DPI
	for _, output := range characterImageOutputsHighDPI {
		err := output.Save(metaImage, character.ID+"@2")

		if err != nil {
			lastError = err
		}
	}

	character.Image.Extension = metaImage.Extension()
	character.Image.Width = metaImage.Image.Bounds().Dx()
	character.Image.Height = metaImage.Image.Bounds().Dy()
	character.Image.AverageColor = GetAverageColor(metaImage.Image)
	character.Image.LastModified = time.Now().Unix()
	return lastError
}

// DownloadImage ...
func (character *Character) DownloadImage(url string) error {
	response, err := client.Get(url).End()

	// Cancel the import if image could not be fetched
	if err != nil {
		return err
	}

	if response.StatusCode() != http.StatusOK {
		return fmt.Errorf("Image response status code: %d", response.StatusCode())
	}

	return character.SetImageBytes(response.Bytes())
}

// HasImage returns true if the character has an image.
func (character *Character) HasImage() bool {
	return character.Image.Extension != "" && character.Image.Width > 0
}
