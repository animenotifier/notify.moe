package characters

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Best characters.
func Best(ctx aero.Context) error {
	characters := fetchAll()

	arn.SortCharactersByLikes(characters)

	return render(ctx, characters)
}
