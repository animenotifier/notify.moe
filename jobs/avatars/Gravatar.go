package main

import (
	"bytes"
	"fmt"
	"image"

	"github.com/animenotifier/arn"
	"github.com/parnurzeal/gorequest"
	gravatar "github.com/ungerik/go-gravatar"
)

// Gravatar - https://gravatar.com/
type Gravatar struct{}

// GetAvatar returns the Gravatar image for a user (if available).
func (source *Gravatar) GetAvatar(user *arn.User) *Avatar {
	// If the user has no Email registered we can't get a Gravatar.
	if user.Email == "" {
		return nil
	}

	// Build URL
	gravatarURL := gravatar.Url(user.Email) + "?s=" + fmt.Sprint(arn.AvatarMaxSize) + "&d=404&r=pg"

	// Download
	response, data, networkErr := gorequest.New().Get(gravatarURL).EndBytes()

	if networkErr != nil || response.StatusCode != 200 {
		return nil
	}

	// Decode
	img, format, decodeErr := image.Decode(bytes.NewReader(data))

	if decodeErr != nil {
		return nil
	}

	return &Avatar{
		User:   user,
		Image:  img,
		Data:   data,
		Format: format,
	}
}
