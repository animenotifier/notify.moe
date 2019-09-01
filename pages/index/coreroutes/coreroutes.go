package coreroutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/embed"
	"github.com/animenotifier/notify.moe/pages/home"
	"github.com/animenotifier/notify.moe/pages/login"
	"github.com/animenotifier/notify.moe/pages/terms"
	"github.com/animenotifier/notify.moe/pages/welcome"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	page.Get(app, "/", home.Get)
	page.Get(app, "/login", login.Get)
	page.Get(app, "/welcome", welcome.Get)
	page.Get(app, "/terms", terms.Get)

	// Browser extension
	page.Get(app, "/extension/embed", embed.Get)
}
