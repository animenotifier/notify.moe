package character

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get character.
func Get(ctx *aero.Context) string {
	id := ctx.Get("id")
	character, err := arn.GetCharacter(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Character not found", err)
	}

	return ctx.HTML(components.CharacterDetails(character))
}
