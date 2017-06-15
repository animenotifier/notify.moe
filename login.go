package main

import "github.com/aerogo/aero"

// EnableLogin ...
func EnableLogin(app *aero.Application) {
	// Google
	EnableGoogleLogin(app)

	// Logout
	app.Get("/logout", func(ctx *aero.Context) string {
		ctx.Session().Set("userId", nil)
		return ctx.Redirect("/")
	})

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
