package arn

import (
	"bytes"
	"image"
	"path"
	"time"

	"github.com/akyoto/imageserver"
)

const (
	// GroupImageSmallWidth is the minimum width in pixels of a small group image.
	GroupImageSmallWidth = 70

	// GroupImageSmallHeight is the minimum height in pixels of a small group image.
	GroupImageSmallHeight = 70

	// GroupImageLargeWidth is the minimum width in pixels of a large group image.
	GroupImageLargeWidth = 280

	// GroupImageLargeHeight is the minimum height in pixels of a large group image.
	GroupImageLargeHeight = 280

	// GroupImageWebPQuality is the WebP quality of group images.
	GroupImageWebPQuality = 70

	// GroupImageJPEGQuality is the JPEG quality of group images.
	GroupImageJPEGQuality = 70

	// GroupImageQualityBonusLowDPI ...
	GroupImageQualityBonusLowDPI = 12

	// GroupImageQualityBonusLarge ...
	GroupImageQualityBonusLarge = 10

	// GroupImageQualityBonusSmall ...
	GroupImageQualityBonusSmall = 15
)

// Define the group image outputs
var groupImageOutputs = []imageserver.Output{
	// Original at full size
	&imageserver.OriginalFile{
		Directory: path.Join(Root, "images/groups/original/"),
		Width:     0,
		Height:    0,
		Quality:   0,
	},

	// JPEG - Small
	&imageserver.JPEGFile{
		Directory: path.Join(Root, "images/groups/small/"),
		Width:     GroupImageSmallWidth,
		Height:    GroupImageSmallHeight,
		Quality:   GroupImageJPEGQuality + GroupImageQualityBonusLowDPI + GroupImageQualityBonusSmall,
	},

	// JPEG - Large
	&imageserver.JPEGFile{
		Directory: path.Join(Root, "images/groups/large/"),
		Width:     GroupImageLargeWidth,
		Height:    GroupImageLargeHeight,
		Quality:   GroupImageJPEGQuality + GroupImageQualityBonusLowDPI + GroupImageQualityBonusLarge,
	},

	// WebP - Small
	&imageserver.WebPFile{
		Directory: path.Join(Root, "images/groups/small/"),
		Width:     GroupImageSmallWidth,
		Height:    GroupImageSmallHeight,
		Quality:   GroupImageWebPQuality + GroupImageQualityBonusLowDPI + GroupImageQualityBonusSmall,
	},

	// WebP - Large
	&imageserver.WebPFile{
		Directory: path.Join(Root, "images/groups/large/"),
		Width:     GroupImageLargeWidth,
		Height:    GroupImageLargeHeight,
		Quality:   GroupImageWebPQuality + GroupImageQualityBonusLowDPI + GroupImageQualityBonusLarge,
	},
}

// Define the high DPI group image outputs
var groupImageOutputsHighDPI = []imageserver.Output{
	// JPEG - Small
	&imageserver.JPEGFile{
		Directory: path.Join(Root, "images/groups/small/"),
		Width:     GroupImageSmallWidth * 2,
		Height:    GroupImageSmallHeight * 2,
		Quality:   GroupImageJPEGQuality + GroupImageQualityBonusSmall,
	},

	// JPEG - Large
	&imageserver.JPEGFile{
		Directory: path.Join(Root, "images/groups/large/"),
		Width:     GroupImageLargeWidth * 2,
		Height:    GroupImageLargeHeight * 2,
		Quality:   GroupImageJPEGQuality + GroupImageQualityBonusLarge,
	},

	// WebP - Small
	&imageserver.WebPFile{
		Directory: path.Join(Root, "images/groups/small/"),
		Width:     GroupImageSmallWidth * 2,
		Height:    GroupImageSmallHeight * 2,
		Quality:   GroupImageWebPQuality + GroupImageQualityBonusSmall,
	},

	// WebP - Large
	&imageserver.WebPFile{
		Directory: path.Join(Root, "images/groups/large/"),
		Width:     GroupImageLargeWidth * 2,
		Height:    GroupImageLargeHeight * 2,
		Quality:   GroupImageWebPQuality + GroupImageQualityBonusLarge,
	},
}

// GroupImage ...
type GroupImage AnimeImage

// SetImageBytes accepts a byte buffer that represents an image file and updates the group image.
func (group *Group) SetImageBytes(data []byte) error {
	// Decode
	img, format, err := image.Decode(bytes.NewReader(data))

	if err != nil {
		return err
	}

	return group.SetImage(&imageserver.MetaImage{
		Image:  img,
		Format: format,
		Data:   data,
	})
}

// SetImage sets the group image to the given MetaImage.
func (group *Group) SetImage(metaImage *imageserver.MetaImage) error {
	var lastError error

	// Save the different image formats and sizes in low DPI
	for _, output := range groupImageOutputs {
		err := output.Save(metaImage, group.ID)

		if err != nil {
			lastError = err
		}
	}

	// Save the different image formats and sizes in high DPI
	for _, output := range groupImageOutputsHighDPI {
		err := output.Save(metaImage, group.ID+"@2")

		if err != nil {
			lastError = err
		}
	}

	group.Image.Extension = metaImage.Extension()
	group.Image.Width = metaImage.Image.Bounds().Dx()
	group.Image.Height = metaImage.Image.Bounds().Dy()
	group.Image.AverageColor = GetAverageColor(metaImage.Image)
	group.Image.LastModified = time.Now().Unix()
	return lastError
}

// HasImage returns true if the group has an image.
func (group *Group) HasImage() bool {
	return group.Image.Extension != "" && group.Image.Width > 0
}
