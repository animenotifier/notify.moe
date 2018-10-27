package characters

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Best characters.
func Best(ctx *aero.Context) string {
	characters := fetchAll()

	arn.SortCharactersByLikes(characters)

	return render(ctx, characters)
}
