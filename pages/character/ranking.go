package character

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Ranking returns the ranking information for the character via the API.
func Ranking(ctx aero.Context) error {
	// Check character ID
	id := ctx.Get("id")
	_, err := arn.GetCharacter(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Character not found", err)
	}

	// Create response object
	response := struct {
		Rank       int     `json:"rank"`
		Percentile float64 `json:"percentile"`
	}{}

	// Sort characters
	characters := arn.FilterCharacters(func(character *arn.Character) bool {
		return !character.IsDraft
	})

	if len(characters) == 0 {
		return ctx.JSON(response)
	}

	arn.SortCharactersByLikes(characters)

	// Allow CORS
	ctx.Response().SetHeader("Access-Control-Allow-Origin", "*")

	// Return ranking
	for index, character := range characters {
		if character.ID == id {
			response.Rank = index + 1
			response.Percentile = float64(response.Rank) / float64(len(characters))
			return ctx.JSON(response)
		}
	}

	// If the ID wasn't found for some reason,
	// return an empty string.
	return ctx.JSON(response)
}
