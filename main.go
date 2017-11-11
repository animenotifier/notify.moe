package main

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/session-store-nano"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/auth"
	"github.com/animenotifier/notify.moe/components/css"
	"github.com/animenotifier/notify.moe/middleware"
	"github.com/animenotifier/notify.moe/pages"
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

	// CSS
	app.SetStyle(css.Bundle())

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
		middleware.Firewall(),
		middleware.Log(),
		middleware.Session(),
		middleware.UserInfo(),
	)

	// API
	arn.API.Install(app)

	// Domain
	if arn.IsDevelopment() {
		app.Config.Domain = "beta.notify.moe"
	}

	// Authentication
	auth.Install(app)

	// Close the database node on shutdown
	app.OnShutdown(arn.Node.Close)

	// Prefetch all collections
	arn.DB.Prefetch()

	// Do not use HTTP/2 push on service worker requests
	app.AddPushCondition(func(ctx *aero.Context) bool {
		return ctx.Request().Header().Get("X-Source") != "service-worker"
	})

	// Specify test routes
	for route, examples := range routeTests {
		app.Test(route, examples)
	}

	return app
}
