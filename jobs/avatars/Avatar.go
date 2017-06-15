package main

import (
	"bytes"
	"fmt"
	"image"
	"net/http"

	"github.com/animenotifier/arn"
	"github.com/parnurzeal/gorequest"
)

// Avatar represents a single image and the name of the format.
type Avatar struct {
	User   *arn.User
	Image  image.Image
	Data   []byte
	Format string
}

// String returns a text representation of the format, width and height.
func (avatar *Avatar) String() string {
	return fmt.Sprint(avatar.Format, " | ", avatar.Image.Bounds().Dx(), "x", avatar.Image.Bounds().Dy())
}

// AvatarFromURL downloads and decodes the image from an URL and creates an Avatar.
func AvatarFromURL(url string, user *arn.User) *Avatar {
	// Download
	response, data, networkErr := gorequest.New().Get(url).EndBytes()

	if networkErr != nil {
		avatarLog.Error("NET", user.Nick, url, networkErr)
		return nil
	}

	if response.StatusCode != http.StatusOK {
		avatarLog.Error("NET", user.Nick, url, response.StatusCode)
		return nil
	}

	// Decode
	img, format, decodeErr := image.Decode(bytes.NewReader(data))

	if decodeErr != nil {
		avatarLog.Error("IMG", user.Nick, url, decodeErr)
		return nil
	}

	return &Avatar{
		User:   user,
		Image:  img,
		Data:   data,
		Format: format,
	}
}
