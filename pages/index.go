package pages

import (
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/index/activityroutes"
	"github.com/animenotifier/notify.moe/pages/index/amvroutes"
	"github.com/animenotifier/notify.moe/pages/index/animeroutes"
	"github.com/animenotifier/notify.moe/pages/index/apiroutes"
	"github.com/animenotifier/notify.moe/pages/index/characterroutes"
	"github.com/animenotifier/notify.moe/pages/index/companyroutes"
	"github.com/animenotifier/notify.moe/pages/index/coreroutes"
	"github.com/animenotifier/notify.moe/pages/index/exploreroutes"
	"github.com/animenotifier/notify.moe/pages/index/exportroutes"
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
	// Register the routes
	coreroutes.Register(app)
	userroutes.Register(app)
	characterroutes.Register(app)
	exploreroutes.Register(app)
	activityroutes.Register(app)
	amvroutes.Register(app)
	forumroutes.Register(app)
	animeroutes.Register(app)
	userlistroutes.Register(app)
	quoteroutes.Register(app)
	companyroutes.Register(app)
	soundtrackroutes.Register(app)
	grouproutes.Register(app)
	searchroutes.Register(app)
	importroutes.Register(app)
	exportroutes.Register(app)
	shoproutes.Register(app)
	settingsroutes.Register(app)
	staffroutes.Register(app)
	apiroutes.Register(app)
}

// Rewrite will rewrite the path before routing happens.
func Rewrite(ctx aero.RewriteContext) {
	requestURI := ctx.Path()

	// User profiles
	if strings.HasPrefix(requestURI, "/+") {
		newURI := "/user/"
		userName := requestURI[2:]
		ctx.SetPath(newURI + userName)
		return
	}

	if strings.HasPrefix(requestURI, "/_/+") {
		newURI := "/_/user/"
		userName := requestURI[4:]
		ctx.SetPath(newURI + userName)
		return
	}

	// Analytics
	if requestURI == "/dark-flame-master" {
		ctx.SetPath("/api/new/analytics")
		return
	}
}
