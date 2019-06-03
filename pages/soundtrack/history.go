package soundtrack

import (
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils/history"

	"github.com/animenotifier/notify.moe/arn"
)

// History of the edits.
var History = history.Handler(renderHistory, "SoundTrack")

func renderHistory(obj interface{}, entries []*arn.EditLogEntry, user *arn.User) string {
	return components.SoundTrackTabs(obj.(*arn.SoundTrack), user) + components.EditLog(entries, user)
}
