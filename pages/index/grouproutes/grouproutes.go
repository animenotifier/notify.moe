package grouproutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/group"
	"github.com/animenotifier/notify.moe/pages/groups"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	// Groups
	page.Get(app, "/groups", groups.Latest)
	page.Get(app, "/groups/from/:index", groups.Latest)
	page.Get(app, "/groups/popular", groups.Popular)
	page.Get(app, "/groups/popular/from/:index", groups.Popular)
	page.Get(app, "/groups/joined", groups.Joined)
	page.Get(app, "/groups/joined/from/:index", groups.Joined)
	page.Get(app, "/group/:id", group.Feed)
	page.Get(app, "/group/:id/info", group.Info)
	page.Get(app, "/group/:id/members", group.Members)
	page.Get(app, "/group/:id/edit", group.Edit)
	page.Get(app, "/group/:id/edit/image", group.EditImage)
	page.Get(app, "/group/:id/history", group.History)
}
