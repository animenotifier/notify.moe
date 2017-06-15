package main

import (
	"bytes"
	"fmt"
	"image"
	"net/http"
	"time"

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

	// Retry after 5 seconds if service unavailable
	if response.StatusCode == http.StatusServiceUnavailable {
		time.Sleep(5 * time.Second)
		response, data, networkErr = gorequest.New().Get(url).EndBytes()
	}

	// Network errors
	if networkErr != nil {
		avatarLog.Error("NET", user.Nick, url, networkErr)
		return nil
	}

	// Bad status codes
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
