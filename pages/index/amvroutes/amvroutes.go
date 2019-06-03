package amvroutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/amv"
	"github.com/animenotifier/notify.moe/pages/amvs"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	// AMV
	page.Get(app, "/amv/:id", amv.Get)
	page.Get(app, "/amv/:id/edit", amv.Edit)
	page.Get(app, "/amv/:id/history", amv.History)

	// AMVs
	page.Get(app, "/amvs", amvs.Latest)
	page.Get(app, "/amvs/from/:index", amvs.Latest)
	page.Get(app, "/amvs/best", amvs.Best)
	page.Get(app, "/amvs/best/from/:index", amvs.Best)
}
