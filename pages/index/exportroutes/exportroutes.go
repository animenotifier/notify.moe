package exportroutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/export"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	app.Get("/user/:nick/animelist/export/text", export.Text)
	app.Get("/user/:nick/animelist/export/json", export.JSON)
}
