package page

import "github.com/aerogo/aero"

// Get registers a layout rendered route and a contents-only route.
func Get(app *aero.Application, route string, handler aero.Handler) {
	app.Get(route, handler)
	app.Get("/_"+route, handler)
}
