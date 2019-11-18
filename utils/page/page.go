package page

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/server/middleware"
)

// Get registers a layout rendered route and a contents-only route.
func Get(app *aero.Application, route string, handler aero.Handler) {
	app.Get(route, middleware.Layout(handler))
	app.Get("/_"+route, handler)
}
