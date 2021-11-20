package userlistroutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/users"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	// User lists
	page.Get(app, "/users", users.Active)
	page.Get(app, "/users/map", users.Map)
	page.Get(app, "/users/noavatar", users.ActiveNoAvatar)
	page.Get(app, "/users/staff", users.Staff)
	page.Get(app, "/users/pro", users.Pro)
	page.Get(app, "/users/editors", users.Editors)
	page.Get(app, "/users/country/:country", users.ByCountry)
}
