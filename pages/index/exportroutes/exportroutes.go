package exportroutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/export"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	app.Get("/user/:nick/animelist/export/csv", export.CSV)
	app.Get("/user/:nick/animelist/export/txt", export.TXT)
	app.Get("/user/:nick/animelist/export/json", export.JSON)
}
