package main

import (
	"strings"

	"github.com/aerogo/aero"
	nanostore "github.com/aerogo/session-store-nano"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/auth"
	"github.com/animenotifier/notify.moe/middleware"
	"github.com/animenotifier/notify.moe/pages"
	"github.com/animenotifier/notify.moe/utils/routetests"
)

var app = aero.New()

func main() {
	// Configure and start
	configure(app).Run()
}

func configure(app *aero.Application) *aero.Application {
	// Sessions
	app.Sessions.Duration = 3600 * 24 * 30 * 6
	app.Sessions.Store = nanostore.New(arn.DB.Collection("Session"))

	// Content security policy
	app.ContentSecurityPolicy.Set("img-src", "https: data:")

	// Security
	configureHTTPS(app)

	// Assets
	configureAssets(app)

	// Pages
	pages.Configure(app)

	// Rewrite
	app.Rewrite(rewrite)

	// Middleware
	app.Use(
		middleware.Log(),
		middleware.Session(),
		middleware.UserInfo(),
	)

	// API
	arn.API.Install(app)

	// Development server configuration
	if arn.IsDevelopment() {
		app.Config.Domain = "beta.notify.moe"
		app.Config.Title += " - Beta"
		app.Config.Manifest.Name = app.Config.Title

		// Test connectivity
		app.OnStart(testConnectivity)
	}

	// Authentication
	auth.Install(app)

	// Close the database node on shutdown
	app.OnEnd(arn.Node.Close)

	// Check that this is the server
	if !arn.Node.IsServer() && !arn.IsTest() {
		panic("Another program is currently running as the database server")
	}

	// Prefetch all collections
	arn.DB.Prefetch()

	// Do not use HTTP/2 push on service worker requests
	app.AddPushCondition(func(ctx *aero.Context) bool {
		return !strings.Contains(ctx.Request().Header().Get("Referer"), "/service-worker")
	})

	// Specify test routes
	for route, examples := range routetests.All() {
		app.Test(route, examples)
	}

	return app
}
