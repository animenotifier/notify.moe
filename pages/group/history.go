package group

import (
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils/history"
)

// History of the edits.
var History = history.Handler(renderHistory, "Group")

func renderHistory(obj interface{}, entries []*arn.EditLogEntry, user *arn.User) string {
	group := obj.(*arn.Group)
	var member *arn.GroupMember

	if user != nil {
		member = group.FindMember(user.ID)
	}

	return components.GroupHeader(group, member, user) + components.EditLog(entries, user)
}
