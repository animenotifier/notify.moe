package pages

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/layout"
	"github.com/animenotifier/notify.moe/layout"
	"github.com/animenotifier/notify.moe/pages/index/amvroutes"
	"github.com/animenotifier/notify.moe/pages/index/animeroutes"
	"github.com/animenotifier/notify.moe/pages/index/apiroutes"
	"github.com/animenotifier/notify.moe/pages/index/characterroutes"
	"github.com/animenotifier/notify.moe/pages/index/companyroutes"
	"github.com/animenotifier/notify.moe/pages/index/coreroutes"
	"github.com/animenotifier/notify.moe/pages/index/exploreroutes"
	"github.com/animenotifier/notify.moe/pages/index/forumroutes"
	"github.com/animenotifier/notify.moe/pages/index/grouproutes"
	"github.com/animenotifier/notify.moe/pages/index/importroutes"
	"github.com/animenotifier/notify.moe/pages/index/quoteroutes"
	"github.com/animenotifier/notify.moe/pages/index/searchroutes"
	"github.com/animenotifier/notify.moe/pages/index/settingsroutes"
	"github.com/animenotifier/notify.moe/pages/index/shoproutes"
	"github.com/animenotifier/notify.moe/pages/index/soundtrackroutes"
	"github.com/animenotifier/notify.moe/pages/index/staffroutes"
	"github.com/animenotifier/notify.moe/pages/index/userlistroutes"
	"github.com/animenotifier/notify.moe/pages/index/userroutes"
)

// Configure registers the page routes in the application.
func Configure(app *aero.Application) {
	l := layout.New(app)

	// Set render function for the layout
	l.Render = fullpage.Render

	// Register the routes
	coreroutes.Register(l)
	userroutes.Register(l)
	characterroutes.Register(l)
	exploreroutes.Register(l)
	amvroutes.Register(l)
	forumroutes.Register(l)
	animeroutes.Register(l)
	userlistroutes.Register(l)
	quoteroutes.Register(l)
	companyroutes.Register(l)
	soundtrackroutes.Register(l)
	grouproutes.Register(l)
	searchroutes.Register(l)
	importroutes.Register(l)
	shoproutes.Register(l, app)
	settingsroutes.Register(l)
	staffroutes.Register(l)
	apiroutes.Register(l, app)

	// Mixed
	// l.Page("/database", database.Get)
	// app.Get("/api/select/:data-type/where/:field/is/:field-value", database.Select)
}
