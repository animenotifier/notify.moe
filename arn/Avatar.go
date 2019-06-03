package arn

import (
	"image"
	"os"
	"path"
)

// OriginalImageExtensions includes all the formats that an avatar source could have sent to us.
var OriginalImageExtensions = []string{
	".jpg",
	".png",
	".gif",
}

// LoadImage loads an image from the given path.
func LoadImage(path string) (img image.Image, format string, err error) {
	f, openErr := os.Open(path)

	if openErr != nil {
		return nil, "", openErr
	}

	img, format, decodeErr := image.Decode(f)

	if decodeErr != nil {
		return nil, "", decodeErr
	}

	return img, format, nil
}

// FindFileWithExtension tries to test different file extensions.
func FindFileWithExtension(baseName string, dir string, extensions []string) string {
	for _, ext := range extensions {
		if _, err := os.Stat(path.Join(dir, baseName+ext)); !os.IsNotExist(err) {
			return dir + baseName + ext
		}
	}

	return ""
}
