package server

import (
	"flag"
	"net/http"
	"strings"

	"github.com/aerogo/aero"
	nanostore "github.com/aerogo/session-store-nano"
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/assets"
	"github.com/animenotifier/notify.moe/pages"
	"github.com/animenotifier/notify.moe/server/auth"
	"github.com/animenotifier/notify.moe/server/https"
	"github.com/animenotifier/notify.moe/server/middleware"
)

// New creates a new web server.
func New() *aero.Application {
	app := aero.New()

	// Sessions
	app.Sessions.Duration = 3600 * 24 * 30 * 6
	app.Sessions.Store = nanostore.New(arn.DB.Collection("Session"))
	app.Sessions.SameSite = http.SameSiteNoneMode

	// Content security policy
	app.ContentSecurityPolicy.Set("img-src", "https: data:")
	app.ContentSecurityPolicy.Set("connect-src", "https: wss: data:")
	app.ContentSecurityPolicy.Set("font-src", "https: data:")

	// Security
	https.Configure(app)

	// Assets
	assets.Configure(app)

	// Pages
	pages.Configure(app)

	// Rewrite
	app.Rewrite(pages.Rewrite)

	// Middleware
	if IsTest() {
		app.Use(middleware.OpenGraph)
	} else {
		app.Use(
			middleware.Recover,
			middleware.HTTPSRedirect,
			middleware.OpenGraph,
			middleware.Log,
			middleware.Session,
			middleware.UserInfo,
		)
	}

	// API
	arn.API.Install(app)

	// Development server configuration
	if arn.IsDevelopment() {
		assets.Domain = "beta.notify.moe"
		assets.Manifest.Name += " - Beta"
	}

	// Authentication
	auth.Install(app)

	// Close the database node on shutdown
	app.OnEnd(arn.Node.Close)

	// Don't push when an underscore URL has been requested
	app.AddPushCondition(func(ctx aero.Context) bool {
		return !strings.HasPrefix(ctx.Path(), "/_")
	})

	// Show errors in the console
	app.OnError(func(ctx aero.Context, err error) {
		color.Red(err.Error())
	})

	// Check that this is the server
	if !arn.Node.IsServer() && !IsTest() {
		panic("Another program is currently running as the database server")
	}

	// Prefetch all collections
	arn.DB.Prefetch()

	// Do not use HTTP/2 push on service worker requests
	app.AddPushCondition(func(ctx aero.Context) bool {
		return !strings.Contains(ctx.Request().Header("Referer"), "/service-worker")
	})

	return app
}

// IsTest returns true if the program is currently running in the "go test" tool.
func IsTest() bool {
	return flag.Lookup("test.v") != nil
}
