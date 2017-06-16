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
}
