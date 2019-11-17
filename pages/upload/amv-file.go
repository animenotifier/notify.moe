package upload

import (
	"net/http"

	"github.com/animenotifier/notify.moe/arn"

	"github.com/aerogo/aero"
)

// AMVFile handles the video upload for AMV files.
func AMVFile(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	amvID := ctx.Get("id")

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	amv, err := arn.GetAMV(amvID)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "AMV not found", err)
	}

	// Retrieve file from post body
	reader := ctx.Request().Body().Reader()
	defer reader.Close()

	// Set amv video file
	err = amv.SetVideoReader(reader)

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Invalid video format", err)
	}

	// Save video information
	amv.Save()

	// Write log entry
	logEntry := arn.NewEditLogEntry(user.ID, "edit", "AMV", amv.ID, "File", "", "")
	logEntry.Save()

	return nil
}
