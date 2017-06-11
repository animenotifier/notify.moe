package main

import "github.com/aerogo/aero"

func init() {
	// Authentication
	EnableGoogleLogin(app)

	// Session middleware
	app.Use(func(ctx *aero.Context, next func()) {
		// Handle the request first
		next()

		// Update session if it has been modified
		if ctx.HasSession() && ctx.Session().Modified() {
			app.Sessions.Store.Set(ctx.Session().ID(), ctx.Session())
		}
	})
}
