package main

import (
	"strings"

	"github.com/aerogo/aero"
	nanostore "github.com/aerogo/session-store-nano"
	"github.com/akyoto/color"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/assets"
	"github.com/animenotifier/notify.moe/auth"
	"github.com/animenotifier/notify.moe/graphql"
	"github.com/animenotifier/notify.moe/middleware"
	"github.com/animenotifier/notify.moe/pages"
	"github.com/animenotifier/notify.moe/utils/routetests"
)

func main() {
	// Configure and start
	app := aero.New()
	configure(app).Run()
}

func configure(app *aero.Application) *aero.Application {
	// Sessions
	app.Sessions.Duration = 3600 * 24 * 30 * 6
	app.Sessions.Store = nanostore.New(arn.DB.Collection("Session"))

	// Content security policy
	app.ContentSecurityPolicy.Set("img-src", "https: data:")
	app.ContentSecurityPolicy.Set("connect-src", "https: wss: data:")
	app.ContentSecurityPolicy.Set("font-src", "https: data:")

	// Security
	configureHTTPS(app)

	// Assets
	assets.Configure(app)

	// Pages
	pages.Configure(app)

	// Rewrite
	app.Rewrite(rewrite)

	// Middleware
	app.Use(
		middleware.Recover,
		middleware.OpenGraph,
		middleware.Log,
		middleware.Session,
		middleware.UserInfo,
	)

	// API
	arn.API.Install(app)

	// Development server configuration
	if arn.IsDevelopment() {
		assets.Domain = "beta.notify.moe"
		assets.Manifest.Name += " - Beta"
	}

	// Authentication
	auth.Install(app)

	// GraphQL
	graphql.Install(app)

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
	if !arn.Node.IsServer() && !arn.IsTest() {
		panic("Another program is currently running as the database server")
	}

	// Prefetch all collections
	arn.DB.Prefetch()

	// Do not use HTTP/2 push on service worker requests
	app.AddPushCondition(func(ctx aero.Context) bool {
		return !strings.Contains(ctx.Request().Header("Referer"), "/service-worker")
	})

	// Specify test routes
	for route, examples := range routetests.All() {
		app.Test(route, examples...)
	}

	return app
}
