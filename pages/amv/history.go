package amv

import (
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils/history"
)

// History of the edits.
var History = history.Handler(renderHistory, "AMV")

func renderHistory(obj interface{}, entries []*arn.EditLogEntry, user *arn.User) string {
	return components.AMVTabs(obj.(*arn.AMV), user) + components.EditLog(entries, user)
}
