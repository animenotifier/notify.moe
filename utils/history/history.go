package history

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Handler returns a function that renders the history of any object.
func Handler(render func(interface{}, []*arn.EditLogEntry, *arn.User) string, typeNames ...string) func(ctx aero.Context) error {
	return func(ctx aero.Context) error {
		id := ctx.Get("id")
		user := arn.GetUserFromContext(ctx)
		obj, err := arn.DB.Get(typeNames[0], id)

		if err != nil {
			return ctx.Error(http.StatusNotFound, typeNames[0]+" not found", err)
		}

		entries := arn.FilterEditLogEntries(func(entry *arn.EditLogEntry) bool {
			return entry.ObjectID == id && arn.Contains(typeNames, entry.ObjectType)
		})

		arn.SortEditLogEntriesLatestFirst(entries)

		return ctx.HTML(render(obj, entries, user))
	}
}
