package main

import (
	"bytes"
	"fmt"
	"image"
	"net/http"
	"strings"
	"time"

	"github.com/animenotifier/arn"
	"github.com/parnurzeal/gorequest"
)

var netLog = avatarLog.NewChannel("NET")

// Avatar represents a single image and the name of the format.
type Avatar struct {
	User   *arn.User
	Image  image.Image
	Data   []byte
	Format string
}

// Extension ...
func (avatar *Avatar) Extension() string {
	switch avatar.Format {
	case "jpg", "jpeg":
		return ".jpg"
	case "png":
		return ".png"
	case "gif":
		return ".gif"
	default:
		return ""
	}
}

// String returns a text representation of the format, width and height.
func (avatar *Avatar) String() string {
	return fmt.Sprint(avatar.Format, " | ", avatar.Image.Bounds().Dx(), "x", avatar.Image.Bounds().Dy())
}

// AvatarFromURL downloads and decodes the image from an URL and creates an Avatar.
func AvatarFromURL(url string, user *arn.User) *Avatar {
	// Download
	response, data, networkErrs := gorequest.New().Get(url).EndBytes()

	// Network errors
	if len(networkErrs) > 0 {
		netLog.Error(user.Nick, url, networkErrs[0])
		return nil
	}

	// Retry HTTP only version after 5 seconds if service unavailable
	if response == nil || response.StatusCode == http.StatusServiceUnavailable {
		time.Sleep(5 * time.Second)
		response, data, networkErrs = gorequest.New().Get(strings.Replace(url, "https://", "http://", 1)).EndBytes()
	}

	// Network errors on 2nd try
	if len(networkErrs) > 0 {
		netLog.Error(user.Nick, url, networkErrs[0])
		return nil
	}

	// Bad status codes
	if response.StatusCode != http.StatusOK {
		netLog.Error(user.Nick, url, response.StatusCode)
		return nil
	}

	// Decode
	img, format, decodeErr := image.Decode(bytes.NewReader(data))

	if decodeErr != nil {
		netLog.Error(user.Nick, url, decodeErr)
		return nil
	}

	return &Avatar{
		User:   user,
		Image:  img,
		Data:   data,
		Format: format,
	}
}
