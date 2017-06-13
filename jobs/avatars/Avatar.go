package main

import "image"

// Avatar represents a single image and the name of the format.
type Avatar struct {
	Image  image.Image
	Data   []byte
	Format string
}
