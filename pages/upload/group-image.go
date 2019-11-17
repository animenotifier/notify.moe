package upload

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// GroupImage handles the group image upload.
func GroupImage(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	groupID := ctx.Get("id")

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return ctx.Error(http.StatusUnauthorized, "Not authorized")
	}

	group, err := arn.GetGroup(groupID)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Group not found", err)
	}

	// Retrieve file from post body
	data, err := ctx.Request().Body().Bytes()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Reading request body failed", err)
	}

	// Set group image file
	err = group.SetImageBytes(data)

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Invalid image format", err)
	}

	// Save image information
	group.Save()

	// Write log entry
	logEntry := arn.NewEditLogEntry(user.ID, "edit", "Group", group.ID, "Image", "", "")
	logEntry.Save()

	return nil
}
