package grouproutes

import (
	"github.com/aerogo/layout"
	"github.com/animenotifier/notify.moe/pages/group"
	"github.com/animenotifier/notify.moe/pages/groups"
)

// Register registers the page routes.
func Register(l *layout.Layout) {
	// Groups
	l.Page("/groups", groups.Get)
	l.Page("/group/:id", group.Feed)
	l.Page("/group/:id/info", group.Info)
	l.Page("/group/:id/members", group.Members)
	l.Page("/group/:id/edit", group.Edit)
}
