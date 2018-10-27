package character

import (
	"net/http"
	"strconv"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Ranking returns the ranking information for the character via the API.
func Ranking(ctx *aero.Context) string {
	id := ctx.Get("id")
	_, err := arn.GetCharacter(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Character not found", err)
	}

	characters := arn.FilterCharacters(func(character *arn.Character) bool {
		return !character.IsDraft
	})

	arn.SortCharactersByLikes(characters)

	for index, character := range characters {
		if character.ID == id {
			return strconv.Itoa(index + 1)
		}
	}

	return ""
}
