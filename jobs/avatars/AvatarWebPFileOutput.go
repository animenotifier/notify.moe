package main

import (
	"github.com/animenotifier/arn"
	"github.com/nfnt/resize"
)

// AvatarWebPFileOutput ...
type AvatarWebPFileOutput struct {
	Directory string
	Size      int
	Quality   float32
}

// SaveAvatar writes the avatar in WebP format to the file system.
func (output *AvatarWebPFileOutput) SaveAvatar(avatar *Avatar) error {
	img := avatar.Image

	// Resize if needed
	if img.Bounds().Dx() != output.Size {
		// Use Lanczos interpolation for downscales
		interpolation := resize.Lanczos3

		// Use Mitchell interpolation for upscales
		if output.Size > img.Bounds().Dx() {
			interpolation = resize.MitchellNetravali
		}

		img = resize.Resize(arn.AvatarSmallSize, 0, img, interpolation)
	}

	// Write to file
	fileName := output.Directory + avatar.User.ID + ".webp"
	return arn.SaveWebP(img, fileName, output.Quality)
}
