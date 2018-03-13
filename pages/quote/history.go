package quote

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// History of the edits.
func History(ctx *aero.Context) string {
	id := ctx.Get("id")
	user := utils.GetUser(ctx)
	quote, err := arn.GetQuote(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Quote not found", err)
	}

	entries := arn.FilterEditLogEntries(func(entry *arn.EditLogEntry) bool {
		return entry.ObjectType == "Quote" && entry.ObjectID == id
	})

	arn.SortEditLogEntriesLatestFirst(entries)

	return ctx.HTML(components.QuoteTabs(quote, user) + components.EditLog(entries, user))
}
