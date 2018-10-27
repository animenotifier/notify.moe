package character

import (
	"net/http"
	"strconv"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Ranking returns the ranking information for the character via the API.
func Ranking(ctx *aero.Context) string {
	// Check character ID
	id := ctx.Get("id")
	_, err := arn.GetCharacter(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Character not found", err)
	}

	// Sort characters
	characters := arn.FilterCharacters(func(character *arn.Character) bool {
		return !character.IsDraft
	})

	arn.SortCharactersByLikes(characters)

	// Allow CORS
	ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")

	// Return ranking
	for index, character := range characters {
		if character.ID == id {
			return strconv.Itoa(index + 1)
		}
	}

	// If the ID wasn't found for some reason,
	// return an empty string.
	return ""
}
