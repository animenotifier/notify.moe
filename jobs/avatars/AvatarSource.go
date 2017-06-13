package main

import (
	"github.com/animenotifier/arn"
)

// AvatarSource describes a source where we can find avatar images for a user.
type AvatarSource interface {
	GetAvatar(*arn.User) *Avatar
}
