package filtercharacters

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// NoImage ...
func NoImage(ctx aero.Context) error {
	return characterList(
		ctx,
		"Characters without an image",
		func(character *arn.Character) bool {
			return !character.HasImage()
		},
		nil,
	)
}
