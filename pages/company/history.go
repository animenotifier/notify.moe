package company

import (
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils/history"
)

// History of the edits.
var History = history.Handler(renderHistory, "Company")

func renderHistory(obj interface{}, entries []*arn.EditLogEntry, user *arn.User) string {
	return components.CompanyTabs(obj.(*arn.Company), user) + components.EditLog(entries, user)
}
