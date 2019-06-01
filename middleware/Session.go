package middleware

import "github.com/aerogo/aero"

// Session middleware saves an existing session if it has been modified.
func Session(next aero.Handler) aero.Handler {
	return func(ctx aero.Context) error {
		// Handle the request first
		err := next(ctx)

		// Update session if it has been modified
		if ctx.HasSession() && ctx.Session().Modified() {
			_ = ctx.App().Sessions.Store.Set(ctx.Session().ID(), ctx.Session())
		}

		return err
	}
}
