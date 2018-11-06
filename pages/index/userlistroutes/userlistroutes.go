package userlistroutes

import (
	"github.com/aerogo/layout"
	"github.com/animenotifier/notify.moe/pages/users"
)

// Register registers the page routes.
func Register(l *layout.Layout) {
	// User lists
	l.Page("/users", users.Active)
	l.Page("/users/map", users.Map)
	l.Page("/users/noavatar", users.ActiveNoAvatar)
	l.Page("/users/games/osu", users.Osu)
	l.Page("/users/games/overwatch", users.Overwatch)
	l.Page("/users/games/ffxiv", users.FFXIV)
	l.Page("/users/staff", users.Staff)
	l.Page("/users/pro", users.Pro)
	l.Page("/users/editors", users.Editors)
	l.Page("/users/country/:country", users.ByCountry)
}
