package filtercharacters

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxCharacterEntries = 70

// characterList renders the characters with the given filter for editors.
func characterList(ctx aero.Context, title string, filter func(*arn.Character) bool, searchLink func(*arn.Character) string) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil || (user.Role != "admin" && user.Role != "editor") {
		return ctx.Error(http.StatusUnauthorized, "Not authorized")
	}

	characters, count := filterCharacters(filter)

	return ctx.HTML(components.CharacterEditorListFull(
		title,
		characters,
		count,
		searchLink,
		user,
	))
}

// filterCharacters filters anime by the given filter function and
// additionally applies year and types filters if specified.
func filterCharacters(filter func(*arn.Character) bool) ([]*arn.Character, int) {
	// Filter
	characters := arn.FilterCharacters(func(character *arn.Character) bool {
		if character.IsDraft {
			return false
		}

		return filter(character)
	})

	// Sort
	arn.SortCharactersByLikes(characters)

	// Limit
	count := len(characters)

	if count > maxCharacterEntries {
		characters = characters[:maxCharacterEntries]
	}

	return characters, count
}
