package character

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils/editform"
)

// Edit character.
func Edit(ctx aero.Context) error {
	id := ctx.Get("id")
	character, err := arn.GetCharacter(id)
	user := arn.GetUserFromContext(ctx)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Character not found", err)
	}

	return ctx.HTML(components.CharacterTabs(character, user) + editform.Render(character, "Edit character", user))
}

// EditImages renders the form to edit the character images.
func EditImages(ctx aero.Context) error {
	id := ctx.Get("id")
	character, err := arn.GetCharacter(id)
	user := arn.GetUserFromContext(ctx)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Character not found", err)
	}

	return ctx.HTML(components.EditCharacterImages(character, user))
}
