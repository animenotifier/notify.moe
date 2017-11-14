package main

import (
	"github.com/animenotifier/arn"
)

// Character ...
type Character struct {
	ID          string                    `json:"id"`
	Name        string                    `json:"name"`
	Image       string                    `json:"image"`
	Description string                    `json:"description"`
	Attributes  []*arn.CharacterAttribute `json:"attributes"`
}

func main() {
	defer arn.Node.Close()

	// Overwrite existing type
	arn.DB.RegisterTypes((*Character)(nil))

	// characters := arn.AllCharacters()
}
