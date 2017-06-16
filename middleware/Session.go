package middleware

import "github.com/aerogo/aero"

// Session middleware saves an existing session if it has been modified.
func Session() aero.Middleware {
	return func(ctx *aero.Context, next func()) {
		// Handle the request first
		next()

		// Update session if it has been modified
		if ctx.HasSession() && ctx.Session().Modified() {
			ctx.App.Sessions.Store.Set(ctx.Session().ID(), ctx.Session())
		}
	}
}
