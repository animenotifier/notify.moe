package main

import (
	"fmt"
	"time"

	"github.com/animenotifier/arn"
	gravatar "github.com/ungerik/go-gravatar"
)

// Gravatar - https://gravatar.com/
type Gravatar struct {
	Rating         string
	RequestLimiter *time.Ticker
}

// GetAvatar returns the Gravatar image for a user (if available).
func (source *Gravatar) GetAvatar(user *arn.User) *Avatar {
	// If the user has no Email registered we can't get a Gravatar.
	if user.Email == "" {
		return nil
	}

	// Build URL
	gravatarURL := gravatar.Url(user.Email) + "?s=" + fmt.Sprint(arn.AvatarMaxSize) + "&d=404&r=" + source.Rating

	// Wait for request limiter to allow us to send a request
	<-source.RequestLimiter.C

	// Download
	return AvatarFromURL(gravatarURL, user)
}
