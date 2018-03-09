package editlog

import (
	"net/http"
	"sort"

	"github.com/animenotifier/arn"

	"github.com/animenotifier/notify.moe/components"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/utils"
)

const maxEntries = 40

// Get edit log.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return ctx.Error(http.StatusUnauthorized, "Not authorized", nil)
	}

	entries := arn.AllEditLogEntries()

	// Sort by creation date
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Created > entries[j].Created
	})

	if len(entries) > maxEntries {
		entries = entries[:maxEntries]
	}

	return ctx.HTML(components.EditLog(entries, user))
}
