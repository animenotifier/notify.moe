package arn

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/api"
)

// Force interface implementations
var (
	_ api.Creatable = (*AnimeCharacter)(nil)
)

// Create sets the data for new anime characters.
func (character *AnimeCharacter) Create(ctx aero.Context) error {
	character.Role = "supporting"
	return nil
}
