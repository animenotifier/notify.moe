package upload

import (
	"net/http"

	"github.com/animenotifier/notify.moe/arn"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/utils"
)

// AMVFile handles the video upload for AMV files.
func AMVFile(ctx aero.Context) error {
	user := utils.GetUser(ctx)
	amvID := ctx.Get("id")

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	amv, err := arn.GetAMV(amvID)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "AMV not found", err)
	}

	// Retrieve file from post body
	data, err := ctx.Request().Body().Bytes()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Reading request body failed", err)
	}

	// Set amv image file
	err = amv.SetVideoBytes(data)

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Invalid video format", err)
	}

	// Save image information
	amv.Save()

	// Write log entry
	logEntry := arn.NewEditLogEntry(user.ID, "edit", "AMV", amv.ID, "File", "", "")
	logEntry.Save()

	return nil
}
