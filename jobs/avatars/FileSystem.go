package main

import (
	"bytes"
	"image"
	"io/ioutil"

	"github.com/animenotifier/arn"
)

var fileSystemLog = avatarLog.NewChannel("SSD")

// FileSystem loads avatar from the local filesystem.
type FileSystem struct {
	Directory string
}

// GetAvatar returns the local image for the user.
func (source *FileSystem) GetAvatar(user *arn.User) *Avatar {
	fullPath := arn.FindFileWithExtension(user.ID, source.Directory, arn.OriginalImageExtensions)

	if fullPath == "" {
		fileSystemLog.Error(user.Nick, "Not found on file system")
		return nil
	}

	data, err := ioutil.ReadFile(fullPath)

	if err != nil {
		fileSystemLog.Error(user.Nick, err)
		return nil
	}

	// Decode
	img, format, decodeErr := image.Decode(bytes.NewReader(data))

	if decodeErr != nil {
		fileSystemLog.Error(user.Nick, decodeErr)
		return nil
	}

	return &Avatar{
		User:   user,
		Image:  img,
		Data:   data,
		Format: format,
	}
}
