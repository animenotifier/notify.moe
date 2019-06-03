package arn

import (
	"bytes"
	"image"
	"path"
	"time"

	"github.com/akyoto/imageserver"
)

const (
	// AnimeImageLargeWidth is the minimum width in pixels of a large anime image.
	AnimeImageLargeWidth = 250

	// AnimeImageLargeHeight is the minimum height in pixels of a large anime image.
	AnimeImageLargeHeight = 350

	// AnimeImageMediumWidth is the minimum width in pixels of a medium anime image.
	AnimeImageMediumWidth = 142

	// AnimeImageMediumHeight is the minimum height in pixels of a medium anime image.
	AnimeImageMediumHeight = 200

	// AnimeImageSmallWidth is the minimum width in pixels of a small anime image.
	AnimeImageSmallWidth = 55

	// AnimeImageSmallHeight is the minimum height in pixels of a small anime image.
	AnimeImageSmallHeight = 78

	// AnimeImageWebPQuality is the WebP quality of anime images.
	AnimeImageWebPQuality = 70

	// AnimeImageJPEGQuality is the JPEG quality of anime images.
	AnimeImageJPEGQuality = 70

	// AnimeImageQualityBonusLowDPI ...
	AnimeImageQualityBonusLowDPI = 10

	// AnimeImageQualityBonusLarge ...
	AnimeImageQualityBonusLarge = 5

	// AnimeImageQualityBonusMedium ...
	AnimeImageQualityBonusMedium = 10

	// AnimeImageQualityBonusSmall ...
	AnimeImageQualityBonusSmall = 10
)

// Define the anime image outputs
var animeImageOutputs = []imageserver.Output{
	// Original at full size
	&imageserver.OriginalFile{
		Directory: path.Join(Root, "images/anime/original/"),
		Width:     0,
		Height:    0,
		Quality:   0,
	},

	// JPEG - Large
	&imageserver.JPEGFile{
		Directory: path.Join(Root, "images/anime/large/"),
		Width:     AnimeImageLargeWidth,
		Height:    AnimeImageLargeHeight,
		Quality:   AnimeImageJPEGQuality + AnimeImageQualityBonusLowDPI + AnimeImageQualityBonusLarge,
	},

	// JPEG - Medium
	&imageserver.JPEGFile{
		Directory: path.Join(Root, "images/anime/medium/"),
		Width:     AnimeImageMediumWidth,
		Height:    AnimeImageMediumHeight,
		Quality:   AnimeImageJPEGQuality + AnimeImageQualityBonusLowDPI + AnimeImageQualityBonusMedium,
	},

	// JPEG - Small
	&imageserver.JPEGFile{
		Directory: path.Join(Root, "images/anime/small/"),
		Width:     AnimeImageSmallWidth,
		Height:    AnimeImageSmallHeight,
		Quality:   AnimeImageJPEGQuality + AnimeImageQualityBonusLowDPI + AnimeImageQualityBonusSmall,
	},

	// WebP - Large
	&imageserver.WebPFile{
		Directory: path.Join(Root, "images/anime/large/"),
		Width:     AnimeImageLargeWidth,
		Height:    AnimeImageLargeHeight,
		Quality:   AnimeImageWebPQuality + AnimeImageQualityBonusLowDPI + AnimeImageQualityBonusLarge,
	},

	// WebP - Medium
	&imageserver.WebPFile{
		Directory: path.Join(Root, "images/anime/medium/"),
		Width:     AnimeImageMediumWidth,
		Height:    AnimeImageMediumHeight,
		Quality:   AnimeImageWebPQuality + AnimeImageQualityBonusLowDPI + AnimeImageQualityBonusMedium,
	},

	// WebP - Small
	&imageserver.WebPFile{
		Directory: path.Join(Root, "images/anime/small/"),
		Width:     AnimeImageSmallWidth,
		Height:    AnimeImageSmallHeight,
		Quality:   AnimeImageWebPQuality + AnimeImageQualityBonusLowDPI + AnimeImageQualityBonusSmall,
	},
}

// Define the high DPI anime image outputs
var animeImageOutputsHighDPI = []imageserver.Output{
	// JPEG - Large
	&imageserver.JPEGFile{
		Directory: path.Join(Root, "images/anime/large/"),
		Width:     AnimeImageLargeWidth * 2,
		Height:    AnimeImageLargeHeight * 2,
		Quality:   AnimeImageJPEGQuality + AnimeImageQualityBonusLarge,
	},

	// JPEG - Medium
	&imageserver.JPEGFile{
		Directory: path.Join(Root, "images/anime/medium/"),
		Width:     AnimeImageMediumWidth * 2,
		Height:    AnimeImageMediumHeight * 2,
		Quality:   AnimeImageJPEGQuality + AnimeImageQualityBonusMedium,
	},

	// JPEG - Small
	&imageserver.JPEGFile{
		Directory: path.Join(Root, "images/anime/small/"),
		Width:     AnimeImageSmallWidth * 2,
		Height:    AnimeImageSmallHeight * 2,
		Quality:   AnimeImageJPEGQuality + AnimeImageQualityBonusSmall,
	},

	// WebP - Large
	&imageserver.WebPFile{
		Directory: path.Join(Root, "images/anime/large/"),
		Width:     AnimeImageLargeWidth * 2,
		Height:    AnimeImageLargeHeight * 2,
		Quality:   AnimeImageWebPQuality + AnimeImageQualityBonusLarge,
	},

	// WebP - Medium
	&imageserver.WebPFile{
		Directory: path.Join(Root, "images/anime/medium/"),
		Width:     AnimeImageMediumWidth * 2,
		Height:    AnimeImageMediumHeight * 2,
		Quality:   AnimeImageWebPQuality + AnimeImageQualityBonusMedium,
	},

	// WebP - Small
	&imageserver.WebPFile{
		Directory: path.Join(Root, "images/anime/small/"),
		Width:     AnimeImageSmallWidth * 2,
		Height:    AnimeImageSmallHeight * 2,
		Quality:   AnimeImageWebPQuality + AnimeImageQualityBonusSmall,
	},
}

// AnimeImage ...
type AnimeImage struct {
	Extension    string   `json:"extension"`
	Width        int      `json:"width"`
	Height       int      `json:"height"`
	AverageColor HSLColor `json:"averageColor"`
	LastModified int64    `json:"lastModified"`
}

// SetImageBytes accepts a byte buffer that represents an image file and updates the anime image.
func (anime *Anime) SetImageBytes(data []byte) error {
	// Decode
	img, format, err := image.Decode(bytes.NewReader(data))

	if err != nil {
		return err
	}

	return anime.SetImage(&imageserver.MetaImage{
		Image:  img,
		Format: format,
		Data:   data,
	})
}

// SetImage sets the anime image to the given MetaImage.
func (anime *Anime) SetImage(metaImage *imageserver.MetaImage) error {
	var lastError error

	// Save the different image formats and sizes in low DPI
	for _, output := range animeImageOutputs {
		err := output.Save(metaImage, anime.ID)

		if err != nil {
			lastError = err
		}
	}

	// Save the different image formats and sizes in high DPI
	for _, output := range animeImageOutputsHighDPI {
		err := output.Save(metaImage, anime.ID+"@2")

		if err != nil {
			lastError = err
		}
	}

	anime.Image.Extension = metaImage.Extension()
	anime.Image.Width = metaImage.Image.Bounds().Dx()
	anime.Image.Height = metaImage.Image.Bounds().Dy()
	anime.Image.AverageColor = GetAverageColor(metaImage.Image)
	anime.Image.LastModified = time.Now().Unix()
	return lastError
}
