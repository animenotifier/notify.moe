package soundtrack

import (
	"net/http"

	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// History of the edits.
func History(ctx *aero.Context) string {
	id := ctx.Get("id")
	user := utils.GetUser(ctx)
	track, err := arn.GetSoundTrack(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Track not found", err)
	}

	entries := arn.FilterEditLogEntries(func(entry *arn.EditLogEntry) bool {
		return entry.ObjectType == "SoundTrack" && entry.ObjectID == id
	})

	arn.SortEditLogEntriesLatestFirst(entries)

	return ctx.HTML(components.SoundTrackTabs(track, user) + components.EditLog(entries, user))
}
