package editlog

import (
	"net/http"

	"github.com/animenotifier/arn"

	"github.com/animenotifier/notify.moe/components"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/utils"
)

const maxEntries = 120

// Get edit log.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return ctx.Error(http.StatusUnauthorized, "Not authorized", nil)
	}

	entries := arn.AllEditLogEntries()

	// Sort by creation date
	arn.SortEditLogEntriesLatestFirst(entries)

	// Limit results
	if len(entries) > maxEntries {
		entries = entries[:maxEntries]
	}

	return ctx.HTML(components.EditLogPage(entries, user))
}
