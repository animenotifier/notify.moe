package activityroutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/activity"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	page.Get(app, "/activity", activity.Posts)
	page.Get(app, "/activity/from/:index", activity.Posts)
	page.Get(app, "/activity/watch", activity.Watch)
	page.Get(app, "/activity/watch/from/:index", activity.Watch)
}
