package editlog

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils/infinitescroll"
)

// render edit log.
func render(ctx aero.Context, compact bool) error {
	user := arn.GetUserFromContext(ctx)
	index, _ := ctx.GetInt("index")
	nick := ctx.Get("nick")

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return ctx.Error(http.StatusUnauthorized, "Not authorized")
	}

	viewUser, err := arn.GetUserByNick(nick)

	if nick != "" && err != nil {
		return ctx.Error(http.StatusNotFound, "User not found", err)
	}

	allEntries := arn.FilterEditLogEntries(func(entry *arn.EditLogEntry) bool {
		if viewUser != nil {
			return entry.UserID == viewUser.ID
		}

		return true
	})

	// Sort by creation date
	arn.SortEditLogEntriesLatestFirst(allEntries)

	// Show only one entry for a single object in compact mode
	if compact {
		filteredEntries := make([]*arn.EditLogEntry, 0, len(allEntries))
		lastObjectID := ""

		for _, entry := range allEntries {
			if lastObjectID == entry.ObjectID {
				continue
			}

			filteredEntries = append(filteredEntries, entry)
			lastObjectID = entry.ObjectID
		}

		allEntries = filteredEntries
	}

	// Slice the part that we need
	entries := allEntries[index:]
	maxLength := entriesFirstLoad

	if index > 0 {
		maxLength = entriesPerScroll
	}

	if len(entries) > maxLength {
		entries = entries[:maxLength]
	}

	// Next index
	nextIndex := infinitescroll.NextIndex(ctx, len(allEntries), maxLength, index)

	// In case we're scrolling, send log entries only (without the page frame)
	if index > 0 {
		return ctx.HTML(components.EditLogScrollable(entries, user))
	}

	// Otherwise, send the full page
	return ctx.HTML(components.EditLogPage(entries, nextIndex, viewUser, user))
}
