package main

import (
	"image"

	"github.com/animenotifier/arn"
)

// Avatar represents a single image and the name of the format.
type Avatar struct {
	User   *arn.User
	Image  image.Image
	Data   []byte
	Format string
}
