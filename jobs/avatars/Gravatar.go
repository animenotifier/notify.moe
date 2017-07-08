package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/animenotifier/arn"
	gravatar "github.com/ungerik/go-gravatar"
)

var gravatarLog = avatarLog.NewChannel("GRA")

// Gravatar - https://gravatar.com/
type Gravatar struct {
	Rating         string
	RequestLimiter *time.Ticker
}

// GetAvatar returns the Gravatar image for a user (if available).
func (source *Gravatar) GetAvatar(user *arn.User) *Avatar {
	// If the user has no Email registered we can't get a Gravatar.
	if user.Email == "" {
		gravatarLog.Error(user.Nick, "No Email")
		return nil
	}

	// Build URL
	gravatarURL := gravatar.Url(user.Email) + "?s=" + fmt.Sprint(arn.AvatarMaxSize) + "&d=404&r=" + source.Rating
	gravatarURL = strings.Replace(gravatarURL, "http://", "https://", 1)

	// Wait for request limiter to allow us to send a request
	<-source.RequestLimiter.C

	// Download
	return AvatarFromURL(gravatarURL, user)
}
