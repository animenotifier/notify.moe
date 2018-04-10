package character

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
	character, err := arn.GetCharacter(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Character not found", err)
	}

	entries := arn.FilterEditLogEntries(func(entry *arn.EditLogEntry) bool {
		return entry.ObjectType == "Character" && entry.ObjectID == id
	})

	arn.SortEditLogEntriesLatestFirst(entries)

	return ctx.HTML(components.CharacterTabs(character, user) + components.EditLog(entries, user))
}
