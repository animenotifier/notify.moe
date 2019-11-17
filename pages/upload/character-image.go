package upload

import (
	"net/http"

	"github.com/animenotifier/notify.moe/arn"

	"github.com/aerogo/aero"
)

// CharacterImage handles the character image upload.
func CharacterImage(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	characterID := ctx.Get("id")

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return ctx.Error(http.StatusUnauthorized, "Not authorized")
	}

	character, err := arn.GetCharacter(characterID)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Character not found", err)
	}

	// Retrieve file from post body
	data, err := ctx.Request().Body().Bytes()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Reading request body failed", err)
	}

	// Set character image file
	err = character.SetImageBytes(data)

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Invalid image format", err)
	}

	// Save image information
	character.Save()

	// Write log entry
	logEntry := arn.NewEditLogEntry(user.ID, "edit", "Character", character.ID, "Image", "", "")
	logEntry.Save()

	return nil
}
