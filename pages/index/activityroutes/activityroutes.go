package activityroutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/activity"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	page.Get(app, "/activity", activity.Global)
	page.Get(app, "/activity/from/:index", activity.Global)
	page.Get(app, "/activity/followed", activity.Followed)
	page.Get(app, "/activity/followed/from/:index", activity.Followed)
}
