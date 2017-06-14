package main

import (
	"bytes"
	"errors"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"

	"github.com/nfnt/resize"
)

// AvatarOriginalFileOutput ...
type AvatarOriginalFileOutput struct {
	Directory string
	Size      int
}

// SaveAvatar writes the original avatar to the file system.
func (output *AvatarOriginalFileOutput) SaveAvatar(avatar *Avatar) error {
	// Determine file extension
	extension := ""

	switch avatar.Format {
	case "jpg", "jpeg":
		extension = ".jpg"
	case "png":
		extension = ".png"
	case "gif":
		extension = ".gif"
	default:
		return errors.New("Unknown format: " + avatar.Format)
	}

	// Resize if needed
	data := avatar.Data
	img := avatar.Image

	if img.Bounds().Dx() > output.Size {
		img = resize.Resize(uint(output.Size), 0, img, resize.Lanczos3)
		buffer := new(bytes.Buffer)

		var err error
		switch extension {
		case ".jpg":
			err = jpeg.Encode(buffer, img, nil)
		case ".png":
			err = png.Encode(buffer, img)
		case ".gif":
			err = gif.Encode(buffer, img, nil)
		}

		if err != nil {
			return err
		}

		data = buffer.Bytes()
	}

	// Write to file
	fileName := output.Directory + avatar.User.ID + extension
	return ioutil.WriteFile(fileName, data, 0644)
}
